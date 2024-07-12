package deployworker

import (
	"context"
	"fmt"
	"log/slog"
	"synchronizationService/internal/entity"
)

type algorithmStatusRepo interface {
	GetAllAlgorithmStatuses(ctx context.Context) ([]*entity.AlgorithmStatus, error)
}

type deployer interface {
	CreatePod(name string) error
	DeletePod(name string) error
	GetPodList() ([]string, error)
}

// Worker содержит в себе логику для запуска и остановки подов-алгоритмов клиентов согласно статусу алгоритмов
type Worker struct {
	algorithmStatus algorithmStatusRepo
	deployer        deployer
}

func NewWorker(deployer deployer, algorithmStatus algorithmStatusRepo) *Worker {
	return &Worker{deployer: deployer, algorithmStatus: algorithmStatus}
}

// Work входная точка Worker
func (w *Worker) Work(ctx context.Context) {
	//- Запросить лист подов и сложить в мапку A map[string]struct{}
	podList, err := w.deployer.GetPodList()
	if err != nil {
		slog.ErrorContext(
			ctx,
			"deploy worker",
			slog.String("error", fmt.Errorf("get pod list: %w", err).Error()),
		)
		return
	}

	mapPodList := make(map[string]struct{}, len(podList))
	for _, pod := range podList {
		mapPodList[pod] = struct{}{}
	}

	//- Запросить все алгоритм статусы и сложить в мапку B map[string]bool,
	//где ключ - %s_%s, client_id, alg_name (на каждый статуc по 3 значения соответственно)

	statuses, err := w.algorithmStatus.GetAllAlgorithmStatuses(ctx)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"deploy worker",
			slog.String("error", fmt.Errorf("get algorithm statuses: %w", err).Error()),
		)
		return
	}

	statusesNames := make(map[string]bool, len(statuses)*3)
	for _, status := range statuses {
		statusesNames[fmt.Sprintf("%d_%s", status.ClientID, "VWAP")] = status.VWAP
		statusesNames[fmt.Sprintf("%d_%s", status.ClientID, "TWAP")] = status.TWAP
		statusesNames[fmt.Sprintf("%d_%s", status.ClientID, "HFT")] = status.HFT
	}

	//- Пройтись по мапке B
	//   - Если ключа из B нет в А и значение true, то создать под
	//   - Если ключ из B есть в А и значение false, то удаляем под

	for algorithm, enabled := range statusesNames {
		_, podExists := mapPodList[algorithm]
		if !podExists && enabled {
			if err := w.deployer.CreatePod(algorithm); err != nil {
				slog.ErrorContext(
					ctx,
					"deploy worker",
					slog.String("error", fmt.Errorf("map iterating: %w", err).Error()),
				)
				return
			}
		} else if podExists && !enabled {
			if err := w.deployer.DeletePod(algorithm); err != nil {
				slog.ErrorContext(
					ctx,
					"deploy worker",
					slog.String("error", fmt.Errorf("map iterating: %w", err).Error()),
				)
				return
			}
		}
	}
	return
}

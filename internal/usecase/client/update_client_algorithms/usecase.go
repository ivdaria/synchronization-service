package updateclientalgorithms

import (
	"context"
	"fmt"
	"synchronizationService/internal/entity"
)

//go:generate mockgen -source=usecase.go -destination=./mocks/mock.go -typed

type algorithmStatusesRepo interface {
	UpdateAlgorithmStatus(ctx context.Context, algorithm *entity.AlgorithmStatus) error
}

// UseCase содержит в себе бизнес логику обновления информации о статусе алгоритмов клиента.
type UseCase struct {
	algorithmStatusRepo algorithmStatusesRepo
}

func NewUseCase(algorithmStatusRepo algorithmStatusesRepo) *UseCase {
	return &UseCase{algorithmStatusRepo: algorithmStatusRepo}
}

// UpdateAlgorithmStatus Входная точка UseCase
func (uc *UseCase) UpdateAlgorithmStatus(ctx context.Context, algStat *entity.AlgorithmStatus) error {
	err := uc.algorithmStatusRepo.UpdateAlgorithmStatus(ctx, algStat)
	if err != nil {
		return fmt.Errorf("update algorithm status: %w", err)
	}
	return nil
}

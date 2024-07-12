package createclient

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"synchronizationService/internal/entity"
)

//go:generate mockgen -source=usecase.go -destination=./mocks/mock.go -typed

type txMock interface {
	pgx.Tx
}

type txManager interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

type algorithmStatusesRepo interface {
	CreateAlgorithmWithTx(ctx context.Context, tx pgx.Tx, algorithm *entity.AlgorithmStatus) error
}

type clientsRepo interface {
	AddClientWithTx(ctx context.Context, tx pgx.Tx, client *entity.Client) (int64, error)
}

// UseCase содержит в себе бизнес логику создания клиента.
// Создает клиента и его статусы
type UseCase struct {
	algorithmStatuses algorithmStatusesRepo
	clients           clientsRepo
	txManager         txManager
}

func NewUseCase(algorithmStatuses algorithmStatusesRepo, clients clientsRepo, txManager txManager) *UseCase {
	return &UseCase{algorithmStatuses: algorithmStatuses, clients: clients, txManager: txManager}
}

// CreateClient входная точка UseCase
func (uc *UseCase) CreateClient(ctx context.Context, client *entity.Client) (int64, error) {
	tx, err := uc.txManager.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	id, err := uc.clients.AddClientWithTx(ctx, tx, client)
	if err != nil {
		return 0, fmt.Errorf("add client by repo: %w", err)
	}

	if err := uc.algorithmStatuses.CreateAlgorithmWithTx(ctx, tx, &entity.AlgorithmStatus{
		ClientID: id,
	}); err != nil {
		return 0, fmt.Errorf("add algorithm status by repo: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("commit tx: %w", err)
	}

	return id, nil
}

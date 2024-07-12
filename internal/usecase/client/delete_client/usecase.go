package deleteclient

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -source=usecase.go -destination=./mocks/mock.go -typed

type txMock interface {
	pgx.Tx
}

type txManager interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

type algorithmStatusesRepo interface {
	DeleteAlgorithmStatus(ctx context.Context, tx pgx.Tx, id int64) error
}

type clientsRepo interface {
	DeleteClient(ctx context.Context, tx pgx.Tx, id int64) error
}

// UseCase содержит в себе бизнес логику удаления клиента.
// Удаляет клиента и его статусы
type UseCase struct {
	algorithmStatuses algorithmStatusesRepo
	clients           clientsRepo
	txManager         txManager
}

func NewUseCase(algorithmStatuses algorithmStatusesRepo, clients clientsRepo, txManager txManager) *UseCase {
	return &UseCase{algorithmStatuses: algorithmStatuses, clients: clients, txManager: txManager}
}

// DeleteClient входная точка UseCase
func (uc *UseCase) DeleteClient(ctx context.Context, id int64) error {
	tx, err := uc.txManager.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	err = uc.algorithmStatuses.DeleteAlgorithmStatus(ctx, tx, id)
	if err != nil {
		return fmt.Errorf("delete algorithm status: %w", err)
	}

	err = uc.clients.DeleteClient(ctx, tx, id)
	if err != nil {
		return fmt.Errorf("delete client: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

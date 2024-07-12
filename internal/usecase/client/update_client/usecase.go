package updateclient

import (
	"context"
	"fmt"
	"synchronizationService/internal/entity"
)

//go:generate mockgen -source=usecase.go -destination=./mocks/mock.go -typed

type clientsRepo interface {
	UpdateClient(ctx context.Context, client *entity.Client) error
}

// UseCase содержит в себе бизнес логику обновления информации о клиенте.
type UseCase struct {
	clients clientsRepo
}

func NewUseCase(clients clientsRepo) *UseCase {
	return &UseCase{clients: clients}
}

// UpdateClient входная точка UseCase
func (uc *UseCase) UpdateClient(ctx context.Context, client *entity.Client) error {
	err := uc.clients.UpdateClient(ctx, client)
	if err != nil {
		return fmt.Errorf("update client: %w", err)
	}
	return nil
}

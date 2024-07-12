package updateclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"synchronizationService/internal/entity"
	mock_update_client "synchronizationService/internal/usecase/client/update_client/mocks"
	"testing"
	"time"
)

func TestUseCase_UpdateClient(t *testing.T) {
	type args struct {
		client *entity.Client
	}
	testClient := &entity.Client{
		ID:          12,
		ClientName:  "1",
		Version:     2,
		Image:       "3",
		CPU:         "4",
		Memory:      "5",
		Priority:    6,
		NeedRestart: true,
		SpawnedAt:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	tests := []struct {
		name  string
		args  args
		setup func(
			clientsMock *mock_update_client.MockclientsRepo,
		)
		wantErr error
	}{
		{
			name: "success",
			args: args{
				client: testClient,
			},
			setup: func(
				clientsMock *mock_update_client.MockclientsRepo,
			) {
				clientsMock.EXPECT().UpdateClient(gomock.Any(), testClient).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "error update client",
			args: args{
				client: testClient,
			},
			setup: func(
				clientsMock *mock_update_client.MockclientsRepo,
			) {
				clientsMock.EXPECT().UpdateClient(gomock.Any(), testClient).Return(io.EOF)
			},
			wantErr: io.EOF,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			clientsMock := mock_update_client.NewMockclientsRepo(ctrl)
			tt.setup(clientsMock)

			uc := NewUseCase(clientsMock)

			err := uc.UpdateClient(context.Background(), tt.args.client)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

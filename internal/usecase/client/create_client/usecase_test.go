package createclient

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"synchronizationService/internal/entity"
	mock_createclient "synchronizationService/internal/usecase/client/create_client/mocks"
	"testing"
	"time"
)

func TestUseCase_CreateClient(t *testing.T) {
	type args struct {
		client *entity.Client
	}

	testClient := &entity.Client{
		ID:          0,
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
			ctrl *gomock.Controller,
			algorithmStatusesMock *mock_createclient.MockalgorithmStatusesRepo,
			clientsMock *mock_createclient.MockclientsRepo,
			txManagerMock *mock_createclient.MocktxManager,
		)
		want    int64
		wantErr error
	}{
		{
			name: "success",
			args: args{
				client: testClient,
			},
			setup: func(
				ctrl *gomock.Controller,
				algorithmStatusesMock *mock_createclient.MockalgorithmStatusesRepo,
				clientsMock *mock_createclient.MockclientsRepo,
				txManagerMock *mock_createclient.MocktxManager,
			) {

				mockTx := mock_createclient.NewMocktxMock(ctrl)

				txManagerMock.EXPECT().Begin(gomock.Any()).Return(mockTx, nil)

				clientsMock.EXPECT().AddClientWithTx(gomock.Any(), mockTx, testClient).
					Return(11, nil)

				algorithmStatusesMock.EXPECT().CreateAlgorithmWithTx(gomock.Any(), mockTx, &entity.AlgorithmStatus{
					ClientID: 11,
				}).Return(nil)

				mockTx.EXPECT().Commit(gomock.Any()).Return(nil)

				mockTx.EXPECT().Rollback(gomock.Any()).Return(pgx.ErrTxClosed)
			},
			want:    11,
			wantErr: nil,
		},
		{
			name: "err commit",
			args: args{
				client: testClient,
			},
			setup: func(
				ctrl *gomock.Controller,
				algorithmStatusesMock *mock_createclient.MockalgorithmStatusesRepo,
				clientsMock *mock_createclient.MockclientsRepo,
				txManagerMock *mock_createclient.MocktxManager,
			) {
				mockTx := mock_createclient.NewMocktxMock(ctrl)
				txManagerMock.EXPECT().Begin(gomock.Any()).Return(mockTx, nil)
				clientsMock.EXPECT().AddClientWithTx(gomock.Any(), mockTx, testClient).
					Return(11, nil)

				algorithmStatusesMock.EXPECT().CreateAlgorithmWithTx(gomock.Any(), mockTx, &entity.AlgorithmStatus{
					ClientID: 11,
				}).Return(nil)

				mockTx.EXPECT().Commit(gomock.Any()).Return(io.EOF)
				mockTx.EXPECT().Rollback(gomock.Any()).Return(pgx.ErrTxClosed)
			},
			want:    0,
			wantErr: io.EOF,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			algorithmStatusesMock := mock_createclient.NewMockalgorithmStatusesRepo(ctrl)
			clientsMock := mock_createclient.NewMockclientsRepo(ctrl)
			txManagerMock := mock_createclient.NewMocktxManager(ctrl)

			tt.setup(ctrl, algorithmStatusesMock, clientsMock, txManagerMock)

			uc := NewUseCase(algorithmStatusesMock, clientsMock, txManagerMock)

			got, err := uc.CreateClient(context.Background(), tt.args.client)

			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

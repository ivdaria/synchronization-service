package deleteclient

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	mock_deleteclient "synchronizationService/internal/usecase/client/delete_client/mocks"
	"testing"
)

func TestUseCase_DeleteClient(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name  string
		args  args
		setup func(
			ctrl *gomock.Controller,
			algorithmStatusesMock *mock_deleteclient.MockalgorithmStatusesRepo,
			clientsMock *mock_deleteclient.MockclientsRepo,
			txManagerMock *mock_deleteclient.MocktxManager,
		)
		wantErr error
	}{
		{
			name: "success",
			args: args{
				id: 12,
			},
			setup: func(
				ctrl *gomock.Controller,
				algorithmStatusesMock *mock_deleteclient.MockalgorithmStatusesRepo,
				clientsMock *mock_deleteclient.MockclientsRepo,
				txManagerMock *mock_deleteclient.MocktxManager,
			) {
				mockTx := mock_deleteclient.NewMocktxMock(ctrl)
				txManagerMock.EXPECT().Begin(gomock.Any()).Return(mockTx, nil)
				algorithmStatusesMock.EXPECT().DeleteAlgorithmStatus(gomock.Any(), mockTx, int64(12)).Return(nil)
				clientsMock.EXPECT().DeleteClient(gomock.Any(), mockTx, int64(12)).
					Return(nil)
				mockTx.EXPECT().Commit(gomock.Any()).Return(nil)
				mockTx.EXPECT().Rollback(gomock.Any()).Return(pgx.ErrTxClosed)
			},
			wantErr: nil,
		},
		{
			name: "error on commit",
			args: args{
				id: 12,
			},
			setup: func(
				ctrl *gomock.Controller,
				algorithmStatusesMock *mock_deleteclient.MockalgorithmStatusesRepo,
				clientsMock *mock_deleteclient.MockclientsRepo,
				txManagerMock *mock_deleteclient.MocktxManager,
			) {
				mockTx := mock_deleteclient.NewMocktxMock(ctrl)
				txManagerMock.EXPECT().Begin(gomock.Any()).Return(mockTx, nil)
				algorithmStatusesMock.EXPECT().DeleteAlgorithmStatus(gomock.Any(), mockTx, int64(12)).Return(nil)
				clientsMock.EXPECT().DeleteClient(gomock.Any(), mockTx, int64(12)).
					Return(nil)
				mockTx.EXPECT().Commit(gomock.Any()).Return(io.EOF)
				mockTx.EXPECT().Rollback(gomock.Any()).Return(pgx.ErrTxClosed)
			},
			wantErr: io.EOF,
		},
		{
			name: "error on delete client",
			args: args{
				id: 12,
			},
			setup: func(
				ctrl *gomock.Controller,
				algorithmStatusesMock *mock_deleteclient.MockalgorithmStatusesRepo,
				clientsMock *mock_deleteclient.MockclientsRepo,
				txManagerMock *mock_deleteclient.MocktxManager,
			) {
				mockTx := mock_deleteclient.NewMocktxMock(ctrl)
				txManagerMock.EXPECT().Begin(gomock.Any()).Return(mockTx, nil)
				algorithmStatusesMock.EXPECT().DeleteAlgorithmStatus(gomock.Any(), mockTx, int64(12)).Return(nil)
				clientsMock.EXPECT().DeleteClient(gomock.Any(), mockTx, int64(12)).
					Return(io.EOF)
				mockTx.EXPECT().Rollback(gomock.Any()).Return(pgx.ErrTxClosed)
			},
			wantErr: io.EOF,
		},
		{
			name: "error on delete algStatus",
			args: args{
				id: 12,
			},
			setup: func(
				ctrl *gomock.Controller,
				algorithmStatusesMock *mock_deleteclient.MockalgorithmStatusesRepo,
				clientsMock *mock_deleteclient.MockclientsRepo,
				txManagerMock *mock_deleteclient.MocktxManager,
			) {
				mockTx := mock_deleteclient.NewMocktxMock(ctrl)
				txManagerMock.EXPECT().Begin(gomock.Any()).Return(mockTx, nil)
				algorithmStatusesMock.EXPECT().DeleteAlgorithmStatus(gomock.Any(), mockTx, int64(12)).Return(io.EOF)
				mockTx.EXPECT().Rollback(gomock.Any()).Return(pgx.ErrTxClosed)
			},
			wantErr: io.EOF,
		},
		{
			name: "error on begin",
			args: args{
				id: 12,
			},
			setup: func(
				ctrl *gomock.Controller,
				algorithmStatusesMock *mock_deleteclient.MockalgorithmStatusesRepo,
				clientsMock *mock_deleteclient.MockclientsRepo,
				txManagerMock *mock_deleteclient.MocktxManager,
			) {
				txManagerMock.EXPECT().Begin(gomock.Any()).Return(nil, io.EOF)
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
			algorithmStatusesMock := mock_deleteclient.NewMockalgorithmStatusesRepo(ctrl)
			clientsMock := mock_deleteclient.NewMockclientsRepo(ctrl)
			txManagerMock := mock_deleteclient.NewMocktxManager(ctrl)
			tt.setup(ctrl, algorithmStatusesMock, clientsMock, txManagerMock)

			uc := NewUseCase(algorithmStatusesMock, clientsMock, txManagerMock)

			err := uc.DeleteClient(context.Background(), tt.args.id)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

package updateclientalgorithms

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"synchronizationService/internal/entity"
	mock_updateclientalgorithms "synchronizationService/internal/usecase/client/update_client_algorithms/mocks"
	"testing"
)

func TestUseCase_UpdateAlgorithmStatus(t *testing.T) {
	type args struct {
		algStat *entity.AlgorithmStatus
	}
	testAlgStat := &entity.AlgorithmStatus{
		ID:       12,
		ClientID: 2,
		VWAP:     true,
		TWAP:     false,
		HFT:      true,
	}

	tests := []struct {
		name  string
		setup func(
			clientsMock *mock_updateclientalgorithms.MockalgorithmStatusesRepo,
		)
		args    args
		wantErr error
	}{
		{
			name: "success",
			setup: func(algMock *mock_updateclientalgorithms.MockalgorithmStatusesRepo) {
				algMock.EXPECT().UpdateAlgorithmStatus(gomock.Any(), testAlgStat).Return(nil)
			},
			args: args{
				algStat: testAlgStat,
			},
			wantErr: nil,
		},
		{
			name: "error update algorithm status",
			setup: func(algMock *mock_updateclientalgorithms.MockalgorithmStatusesRepo) {
				algMock.EXPECT().UpdateAlgorithmStatus(gomock.Any(), testAlgStat).Return(io.EOF)
			},
			args: args{
				algStat: testAlgStat,
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
			clientsMock := mock_updateclientalgorithms.NewMockalgorithmStatusesRepo(ctrl)
			tt.setup(clientsMock)

			uc := NewUseCase(clientsMock)
			err := uc.UpdateAlgorithmStatus(context.Background(), tt.args.algStat)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

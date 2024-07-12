package convert

import (
	"reflect"
	"synchronizationService/internal/entity"
	"synchronizationService/pkg/gateway/model"
	"testing"
)

func TestStatusFromUpdateStatusRequestBody(t *testing.T) {
	type args struct {
		id  int64
		mdl *model.UpdateStatusRequestBody
	}
	tests := []struct {
		name string
		args args
		want *entity.AlgorithmStatus
	}{
		{
			name: "all false",
			args: args{
				id:  1,
				mdl: &model.UpdateStatusRequestBody{},
			},
			want: &entity.AlgorithmStatus{
				ClientID: 1,
			},
		},
		{
			name: "VWAP true",
			args: args{
				id: 1,
				mdl: &model.UpdateStatusRequestBody{
					VWAP: true,
				},
			},
			want: &entity.AlgorithmStatus{
				ClientID: 1,
				VWAP:     true,
			},
		},
		{
			name: "TWAP true",
			args: args{
				id: 1,
				mdl: &model.UpdateStatusRequestBody{
					TWAP: true,
				},
			},
			want: &entity.AlgorithmStatus{
				ClientID: 1,
				TWAP:     true,
			},
		},
		{
			name: "HFT true",
			args: args{
				id: 1,
				mdl: &model.UpdateStatusRequestBody{
					HFT: true,
				},
			},
			want: &entity.AlgorithmStatus{
				ClientID: 1,
				HFT:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusFromUpdateStatusRequestBody(tt.args.id, tt.args.mdl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusFromUpdateStatusRequestBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

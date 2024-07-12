package algorithmstatus

import (
	"reflect"
	"synchronizationService/internal/entity"
	"testing"
)

func Test_model_toAlgorithmStatus(t *testing.T) {
	tests := []struct {
		name string
		want *entity.AlgorithmStatus
	}{
		{
			name: "to alg stat",
			want: &entity.AlgorithmStatus{
				ID:       1,
				ClientID: 2,
				VWAP:     false,
				TWAP:     false,
				HFT:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &model{
				ID:       1,
				ClientID: 2,
				VWAP:     false,
				TWAP:     false,
				HFT:      false,
			}
			if got := m.toAlgorithmStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toAlgorithmStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_modelFromAlgorithmStatus(t *testing.T) {
	type args struct {
		item *entity.AlgorithmStatus
	}
	tests := []struct {
		name string
		args args
		want *model
	}{
		{
			name: "to model",
			args: args{
				item: &entity.AlgorithmStatus{
					ID:       1,
					ClientID: 2,
					VWAP:     false,
					TWAP:     false,
					HFT:      false,
				},
			},
			want: &model{
				ID:       1,
				ClientID: 2,
				VWAP:     false,
				TWAP:     false,
				HFT:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modelFromAlgorithmStatus(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modelFromAlgorithmStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

package convert

import (
	"reflect"
	"synchronizationService/internal/entity"
	"synchronizationService/pkg/gateway/model"
	"testing"
	"time"
)

func TestClientFromAddClientRequestBody(t *testing.T) {
	type args struct {
		mdl *model.AddClientRequestBody
	}

	testTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		args args
		want *entity.Client
	}{
		{
			name: "first",
			args: args{
				mdl: &model.AddClientRequestBody{
					ClientName:  "potato",
					Version:     12,
					Image:       "cat",
					CPU:         "beaver",
					Memory:      "yet another beaver",
					Priority:    11.5,
					NeedRestart: false,
					SpawnedAt:   testTime,
				},
			},
			want: &entity.Client{
				ClientName:  "potato",
				Version:     12,
				Image:       "cat",
				CPU:         "beaver",
				Memory:      "yet another beaver",
				Priority:    11.5,
				NeedRestart: false,
				SpawnedAt:   testTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClientFromAddClientRequestBody(tt.args.mdl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientFromAddClientRequestBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

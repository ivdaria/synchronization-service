package client

import (
	"github.com/stretchr/testify/assert"
	"synchronizationService/internal/entity"
	"testing"
	"time"
)

func Test_model_toClient(t *testing.T) {
	testTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		want *entity.Client
	}{
		{
			name: "entity client",
			want: &entity.Client{
				ID:          1,
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

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &model{
				ID:          1,
				ClientName:  "potato",
				Version:     12,
				Image:       "cat",
				CPU:         "beaver",
				Memory:      "yet another beaver",
				Priority:    11.5,
				NeedRestart: false,
				SpawnedAt:   testTime,
			}
			got := m.toClient()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_modelFromClient(t *testing.T) {
	testTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	defTime := time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	testClient := &entity.Client{
		ID:          1,
		ClientName:  "potato",
		Version:     12,
		Image:       "cat",
		CPU:         "beaver",
		Memory:      "yet another beaver",
		Priority:    11.5,
		NeedRestart: false,
		SpawnedAt:   testTime,
		CreatedAt:   defTime,
		UpdatedAt:   defTime,
	}
	type args struct {
		client *entity.Client
	}
	tests := []struct {
		name string
		args args
		want *model
	}{
		{
			name: "model from client",
			args: args{
				client: testClient,
			},
			want: &model{
				ID:          1,
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

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := modelFromClient(tt.args.client)
			assert.Equal(t, tt.want, got)
		})
	}
}

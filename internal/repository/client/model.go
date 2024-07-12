package client

import (
	"github.com/jackc/pgx/v5"
	"synchronizationService/internal/entity"
	"time"
)

type model struct {
	ID          int64
	ClientName  string
	Version     int
	Image       string
	CPU         string
	Memory      string
	Priority    float64
	NeedRestart bool
	SpawnedAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *model) ScanRow(rows pgx.Rows) error {
	return rows.Scan(&m.ID, &m.ClientName, &m.Version, &m.Image, &m.CPU, &m.Memory,
		&m.Priority, &m.NeedRestart, &m.SpawnedAt, &m.CreatedAt, &m.UpdatedAt)
}

func (m *model) toClient() *entity.Client {
	return &entity.Client{
		ID:          m.ID,
		ClientName:  m.ClientName,
		Version:     m.Version,
		Image:       m.Image,
		CPU:         m.CPU,
		Memory:      m.Memory,
		Priority:    m.Priority,
		NeedRestart: m.NeedRestart,
		SpawnedAt:   m.SpawnedAt,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func modelFromClient(client *entity.Client) *model {
	return &model{
		ID:          client.ID,
		ClientName:  client.ClientName,
		Version:     client.Version,
		Image:       client.Image,
		CPU:         client.CPU,
		Memory:      client.Memory,
		Priority:    client.Priority,
		NeedRestart: client.NeedRestart,
		SpawnedAt:   client.SpawnedAt,
	}
}

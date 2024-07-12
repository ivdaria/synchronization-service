package client

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"synchronizationService/internal/entity"
	er "synchronizationService/internal/errors"
)

// Repo содержит в себе логику хранения клиента
type Repo struct {
	DB *pgx.Conn
}

func NewRepo(DB *pgx.Conn) *Repo {
	return &Repo{DB: DB}
}

// AddClientWithTx добавляет нового клиента в БД в транзакции
func (r *Repo) AddClientWithTx(ctx context.Context, tx pgx.Tx, client *entity.Client) (int64, error) {
	const query = `INSERT INTO clients (client_name, version, image, cpu, memory, priority, need_restart, spawned_at) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`

	var id int64
	mdl := modelFromClient(client)
	err := tx.QueryRow(ctx, query, mdl.ClientName, mdl.Version, mdl.Image, mdl.CPU, mdl.Memory,
		mdl.Priority, mdl.NeedRestart, mdl.SpawnedAt).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("AddClientWithTx scan row:  %w", err)
	}

	return id, nil
}

// UpdateClient обновляет информацию о клиенте в БД
func (r *Repo) UpdateClient(ctx context.Context, client *entity.Client) error {
	const query = `UPDATE clients SET client_name = $2, version = $3, image = $4, cpu = $5, memory = $6, 
                   priority = $7, need_restart = $8, spawned_at = $9, updated_at = now() WHERE id = $1`

	mdl := modelFromClient(client)
	commandTag, err := r.DB.Exec(ctx, query, mdl.ID, mdl.ClientName, mdl.Version, mdl.Image, mdl.CPU, mdl.Memory, mdl.Priority,
		mdl.NeedRestart, mdl.SpawnedAt)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return er.ErrNoRowsAffected
	}
	return nil
}

// DeleteClient удаляет информацию о клиенте из БД в транзакции
func (r *Repo) DeleteClient(ctx context.Context, tx pgx.Tx, id int64) error {
	const query = `DELETE FROM clients WHERE id = $1`

	commandTag, err := tx.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query to delete client by id: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return er.ErrNoRowsAffected
	}
	return nil
}

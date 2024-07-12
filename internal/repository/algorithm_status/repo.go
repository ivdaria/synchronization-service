package algorithmstatus

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"synchronizationService/internal/entity"
	er "synchronizationService/internal/errors"
)

// Repo содержит в себе логику хранения статусов алгоритмов
type Repo struct {
	DB *pgx.Conn
}

func NewRepo(db *pgx.Conn) *Repo {
	return &Repo{DB: db}
}

// CreateAlgorithmWithTx создает статусы алгоритмов в БД в транзакции
func (r *Repo) CreateAlgorithmWithTx(ctx context.Context, tx pgx.Tx, algorithm *entity.AlgorithmStatus) error {
	const query = `INSERT INTO algorithmstatus(client_id, VWAP, TWAP, HFT) VALUES($1, $2, $3, $4) RETURNING id`
	var id int64
	mdl := modelFromAlgorithmStatus(algorithm)
	err := tx.QueryRow(ctx, query, mdl.ClientID, mdl.VWAP, mdl.TWAP, mdl.HFT).Scan(&id)
	if err != nil {
		return fmt.Errorf("create algorithm: %w", err)
	}

	return nil
}

// UpdateAlgorithmStatus обновляет статусы алгоритмов в БД
func (r *Repo) UpdateAlgorithmStatus(ctx context.Context, algorithm *entity.AlgorithmStatus) error {
	const query = `UPDATE algorithmstatus SET vwap = $1, twap = $2, hft = $3 WHERE client_id = $4`

	mdl := modelFromAlgorithmStatus(algorithm)

	commandTag, err := r.DB.Exec(ctx, query, mdl.VWAP, mdl.TWAP, mdl.HFT, mdl.ClientID)

	if err != nil {
		return fmt.Errorf("update algorithm status: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return er.ErrNoRowsAffected
	}

	return nil
}

// DeleteAlgorithmStatus удаляет статусы алгоритмов в БД в транзакции
func (r *Repo) DeleteAlgorithmStatus(ctx context.Context, tx pgx.Tx, id int64) error {
	const query = `DELETE FROM algorithmstatus WHERE algorithmstatus.client_id = $1`
	commandTag, err := tx.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete algorithm status: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return er.ErrNoRowsAffected
	}

	return nil
}

// GetAllAlgorithmStatuses получает статусы алгоритмов всех клиентов из БД
func (r *Repo) GetAllAlgorithmStatuses(ctx context.Context) ([]*entity.AlgorithmStatus, error) {
	const query = `SELECT * FROM algorithmstatus`

	var mdls models

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("select all statuses: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var mdl model
		if err := rows.Scan(&mdl); err != nil {
			return nil, fmt.Errorf("list of statuses scan row:  %w", err)
		}
		mdls = append(mdls, &mdl)
	}

	return mdls.toAlgorithmStatuses(), nil
}

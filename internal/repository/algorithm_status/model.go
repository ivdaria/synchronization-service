package algorithmstatus

import (
	"github.com/jackc/pgx/v5"
	"synchronizationService/internal/entity"
)

type model struct {
	ID       int64
	ClientID int64
	VWAP     bool
	TWAP     bool
	HFT      bool
}

type models []*model

func (m *model) ScanRow(rows pgx.Rows) error {
	return rows.Scan(&m.ID, &m.ClientID, &m.VWAP, &m.TWAP, &m.HFT)
}

func (m *model) toAlgorithmStatus() *entity.AlgorithmStatus {
	return &entity.AlgorithmStatus{
		ID:       m.ID,
		ClientID: m.ClientID,
		VWAP:     m.VWAP,
		TWAP:     m.TWAP,
		HFT:      m.HFT,
	}
}

func (mdls models) toAlgorithmStatuses() []*entity.AlgorithmStatus {
	if len(mdls) == 0 {
		return nil
	}

	result := make([]*entity.AlgorithmStatus, 0, len(mdls))
	for _, m := range mdls {
		result = append(result, m.toAlgorithmStatus())
	}

	return result
}

func modelFromAlgorithmStatus(item *entity.AlgorithmStatus) *model {
	return &model{
		ID:       item.ID,
		ClientID: item.ClientID,
		VWAP:     item.VWAP,
		TWAP:     item.TWAP,
		HFT:      item.HFT,
	}
}

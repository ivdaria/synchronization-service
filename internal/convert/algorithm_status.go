package convert

import (
	"synchronizationService/internal/entity"
	"synchronizationService/pkg/gateway/model"
)

func StatusFromUpdateStatusRequestBody(clientID int64, mdl *model.UpdateStatusRequestBody) *entity.AlgorithmStatus {
	return &entity.AlgorithmStatus{
		ClientID: clientID,
		VWAP:     mdl.VWAP,
		TWAP:     mdl.TWAP,
		HFT:      mdl.HFT,
	}
}

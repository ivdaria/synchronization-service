package convert

import (
	"synchronizationService/internal/entity"
	"synchronizationService/pkg/gateway/model"
)

func ClientFromAddClientRequestBody(mdl *model.AddClientRequestBody) *entity.Client {
	return &entity.Client{
		ID:          0,
		ClientName:  mdl.ClientName,
		Version:     mdl.Version,
		Image:       mdl.Image,
		CPU:         mdl.CPU,
		Memory:      mdl.Memory,
		Priority:    mdl.Priority,
		NeedRestart: mdl.NeedRestart,
		SpawnedAt:   mdl.SpawnedAt,
	}
}

func ClientFromUpdateClientRequestBody(id int64, mdl *model.UpdateClientRequestBody) *entity.Client {
	return &entity.Client{
		ID:          id,
		ClientName:  mdl.ClientName,
		Version:     mdl.Version,
		Image:       mdl.Image,
		CPU:         mdl.CPU,
		Memory:      mdl.Memory,
		Priority:    mdl.Priority,
		NeedRestart: mdl.NeedRestart,
		SpawnedAt:   mdl.SpawnedAt,
	}
}

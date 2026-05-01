package entity_mappers

import (
	"auptex.com/botnova/internals/domain/models"
	gormentities "auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToScriptDomain(e gormentities.Script) models.Script {
	return models.Script{
		Id:        e.Id,
		ProjectID: e.ProjectId,
		Name:      e.Name,
		Type:      e.Type,
		Content:   e.Content,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToScriptEntity(m models.Script) gormentities.Script {
	return gormentities.Script{
		Id:        m.Id,
		ProjectId: m.ProjectID,
		Name:      m.Name,
		Type:      m.Type,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

package entity_mappers

import (
	"auptex.com/botnova/internals/domain/models"
	gormentities "auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToDomain(e gormentities.Project) models.Project {
	return models.Project{
		Id:          e.Id,
		UserId:      e.UserID,
		Name:        e.Name,
		Description: e.Description,
		Options:     e.Options,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToEntity(m models.Project) gormentities.Project {
	return gormentities.Project{
		Id:          m.Id,
		UserID:      m.UserId,
		Name:        m.Name,
		Description: m.Description,
		Options:     m.Options,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

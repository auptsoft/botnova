package entity_mappers

import (
	"auptex.com/botnova/internals/domain/models"
	gormentities "auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToDomain(e gormentities.Project) models.Project {
	return models.Project{
		Id:          e.Id,
		UserID:      e.UserID,
		Name:        e.Name,
		Description: e.Description,
		Options:     e.Options,
	}
}

func ToEntity(m models.Project) gormentities.Project {
	return gormentities.Project{
		Id:          m.Id,
		UserID:      m.UserID,
		Name:        m.Name,
		Description: m.Description,
		Options:     m.Options,
	}
}

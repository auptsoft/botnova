package entity_mappers

import (
	"auptex.com/botnova/internals/domain/models"
	gormentities "auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToDomainRobot(e gormentities.Robot) models.Robot {
	return models.Robot{
		Id:        e.Id,
		Name:      e.Name,
		ModelID:   e.ModelID,
		UserID:    e.UserID,
		Type:      e.Type,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToEntityRobot(m models.Robot) gormentities.Robot {
	return gormentities.Robot{
		Id:        m.Id,
		Name:      m.Name,
		ModelID:   m.ModelID,
		UserID:    m.UserID,
		Type:      m.Type,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

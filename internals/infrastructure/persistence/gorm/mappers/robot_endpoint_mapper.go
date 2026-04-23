package entity_mappers

import (
	"auptex.com/botnova/internals/domain/models"
	gormentities "auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToDomainRobotEndpoint(e gormentities.RobotEndpoint) models.RobotEndpoint {
	return models.RobotEndpoint{
		Id:        e.Id,
		RobotID:   e.RobotID,
		Type:      e.Type,
		Address:   e.Address,
		IsActive:  e.IsActive,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToEntityRobotEndpoint(m models.RobotEndpoint) gormentities.RobotEndpoint {
	return gormentities.RobotEndpoint{
		Id:        m.Id,
		RobotID:   m.RobotID,
		Address:   m.Address,
		Type:      m.Type,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

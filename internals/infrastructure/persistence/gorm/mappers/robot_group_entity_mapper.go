package entity_mappers

import (
	"auptex.com/botnova/internals/domain/models"
	gormentities "auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
)

func ToRobotGroupDomain(e gormentities.RobotGroup) models.RobotGroup {
	return models.RobotGroup{
		Id:             e.Id,
		UserId:         e.UserId,
		Name:           e.Name,
		Description:    e.Description,
		Mode:           e.Mode,
		PrimaryRobotId: e.PrimaryRobotId,
		Options:        e.Options,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}

func ToRobotGroupEntity(m models.RobotGroup) gormentities.RobotGroup {
	return gormentities.RobotGroup{
		Id:             m.Id,
		UserId:         m.UserId,
		Name:           m.Name,
		Description:    m.Description,
		Mode:           m.Mode,
		PrimaryRobotId: m.PrimaryRobotId,
		Options:        m.Options,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

func ToRobotGroupMemberDomain(e gormentities.RobotGroupMember) models.RobotGroupMember {
	return models.RobotGroupMember{
		Id:        e.Id,
		GroupId:   e.GroupId,
		RobotId:   e.RobotId,
		Role:      e.Role,
		Priority:  e.Priority,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToRobotGroupMemberEntity(m models.RobotGroupMember) gormentities.RobotGroupMember {
	return gormentities.RobotGroupMember{
		Id:        m.Id,
		GroupId:   m.GroupId,
		RobotId:   m.RobotId,
		Role:      m.Role,
		Priority:  m.Priority,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

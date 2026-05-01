package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type RobotGroupMemberRepository interface {
	Create(member *models.RobotGroupMember) error
	GetByGroupId(groupId string) ([]models.RobotGroupMember, error)
	GetByRobotId(robotId string) ([]models.RobotGroupMember, error)
	GetById(id string) (*models.RobotGroupMember, error)
	GetByGroupIdAndRole(groupId string, role string) ([]models.RobotGroupMember, error)
	Update(member *models.RobotGroupMember) error
	Delete(groupId string) error
}

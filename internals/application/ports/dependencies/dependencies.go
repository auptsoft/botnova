package dependencies

import (
	"auptex.com/botnova/internals/application/ports"
	repositorydefinitions "auptex.com/botnova/internals/application/ports/repository_definitions"
	"auptex.com/botnova/internals/application/services"
)

type Dependencies struct {
	EventBus      ports.EventBus
	ServiceLogger ports.Logger
	WsTransport   ports.EventTransport
	AuthConfig    services.AuthConfig

	//Repositories
	ProjectRepository          repositorydefinitions.ProjectRepository
	UserRepository             repositorydefinitions.UserRepository
	RobotRepository            repositorydefinitions.RobotRepository
	RobotModelRepository       repositorydefinitions.RobotModelRepository
	ScriptRepository           repositorydefinitions.ScriptRepository
	RobotEndpointRepository    repositorydefinitions.RobotEndpointRepository
	RobotGroupRepository       repositorydefinitions.RobotGroupRepository
	RobotGroupMemberRepository repositorydefinitions.RobotGroupMemberRepository
	CalibrationRepository      repositorydefinitions.CalibrationRepository

	StateStore ports.StateStore

	//Services
	ProjectService   *services.ProjectService
	UserService      *services.UserService
	TransportService *services.TransportService
}

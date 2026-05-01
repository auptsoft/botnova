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
	ProjectRepository repositorydefinitions.ProjectRepository
	UserRepository    repositorydefinitions.UserRepository

	//Services
	ProjectService   *services.ProjectService
	UserService      *services.UserService
	TransportService *services.TransportService
}

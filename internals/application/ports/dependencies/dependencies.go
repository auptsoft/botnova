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

	//Repositories
	ProjectRepository repositorydefinitions.ProjectRepository

	//Services
	ProjectService   *services.ProjectService
	TransportService *services.TransportService
}

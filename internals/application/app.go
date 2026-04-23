package application

import (
	"auptex.com/botnova/internals/api"
	"auptex.com/botnova/internals/application/adapters"
	"auptex.com/botnova/internals/application/ports"
	"auptex.com/botnova/internals/application/ports/dependencies"
	"auptex.com/botnova/internals/application/services"
	"auptex.com/botnova/internals/infrastructure/transport/websocket"
)

type App struct {
	logger   ports.Logger
	eventBus ports.EventBus
	wsServer *websocket.Server
}

func StartAppServer(addr string, deps *dependencies.Dependencies) *App {
	app := &App{
		eventBus: deps.EventBus,
		logger:   deps.ServiceLogger,
	}

	//Set service dependencies
	deps.ProjectService = services.NewProjectService(deps.ProjectRepository, app.logger)
	deps.UserService = services.NewUserService(deps.UserRepository, app.logger)
	deps.TransportService = services.NewTransportService(deps.WsTransport, app.logger)

	app.wsServer = adapters.InitWebsocket(deps.EventBus, deps.ServiceLogger)

	router := api.SetupRouter(deps, app.wsServer)

	router.Run(addr)

	return app
}

package main

import (
	"fmt"

	"auptex.com/botnova/cmd/common"
	"auptex.com/botnova/internals/application"
	"auptex.com/botnova/internals/application/ports"
	"auptex.com/botnova/internals/application/ports/dependencies"
	"auptex.com/botnova/internals/infrastructure/bus"
	"auptex.com/botnova/internals/infrastructure/logger"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/repositories"
	"auptex.com/botnova/internals/infrastructure/transport/websocket"
)

func main() {

	fmt.Println("Initializing BotNova Server...")

	log, err := logger.NewZapLogger()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	cmdConfig, stateConfig, defaultConfig := common.GetEventBusConfigs()
	var eventBus ports.EventBus = bus.NewBus(log, cmdConfig, stateConfig, defaultConfig)

	dbConfig := gorm.Config{
		Driver: "sqlite",
		DSN:    "db.sqlite",
	}

	db := gorm.NewDB(dbConfig)

	dependencies := dependencies.Dependencies{
		EventBus:      eventBus,
		ServiceLogger: log,
		WsTransport:   websocket.NewWebsocketTransport(eventBus),

		//Repositories
		ProjectRepository:          repositories.NewProjectRepository(db),
		UserRepository:             repositories.NewUserRepository(db),
		RobotRepository:            repositories.NewRobotRepository(db),
		RobotModelRepository:       repositories.NewRobotModelRepository(db),
		ScriptRepository:           repositories.NewScriptRepository(db),
		RobotEndpointRepository:    repositories.NewRobotEndpointRepository(db),
		RobotGroupRepository:       repositories.NewRobotGroupRepository(db),
		RobotGroupMemberRepository: repositories.NewRobotGroupMemberRepository(db),
		CalibrationRepository:      repositories.NewCalibrationRepository(db),
	}

	log.Info("Starting botnova server...")
	application.StartAppServer(":5050", &dependencies)
}

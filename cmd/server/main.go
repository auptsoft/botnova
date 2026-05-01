package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"auptex.com/botnova/cmd/common"
	"auptex.com/botnova/internals/application"
	"auptex.com/botnova/internals/application/ports"
	"auptex.com/botnova/internals/application/ports/dependencies"
	"auptex.com/botnova/internals/application/services"
	"auptex.com/botnova/internals/infrastructure/bus"
	"auptex.com/botnova/internals/infrastructure/logger"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/repositories"
	"auptex.com/botnova/internals/infrastructure/transport/websocket"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Initializing BotNova Server...")

	err := godotenv.Load()
	if err != nil {
		// We could panic here but I decided not to because it's not necessary to have environment variables in this project yet, we use SQLite and we have fallbacks.
		fmt.Sprintf("Failed to fetch environment variables: %v", err)
	}

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

	jwtTTLHours := 24 * 7
	if ttlFromEnv := os.Getenv("JWT_TTL_HOURS"); ttlFromEnv != "" {
		parsedTTL, err := strconv.Atoi(ttlFromEnv)
		if err == nil && parsedTTL > 0 {
			jwtTTLHours = parsedTTL
		}
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-me"
	}

	authConfig := &services.AuthConfig{
		JwtSecret: []byte(jwtSecret),
		JwtTTL:    time.Duration(jwtTTLHours) * time.Hour,
	}

	dependencies := dependencies.Dependencies{
		EventBus:      eventBus,
		ServiceLogger: log,
		WsTransport:   websocket.NewWebsocketTransport(eventBus),

		//Repositories
		ProjectRepository: repositories.NewProjectRepository(db),
		UserRepository:    repositories.NewUserRepository(db),
		AuthConfig:        *authConfig,
	}

	log.Info("Starting botnova server...")
	application.StartAppServer(":5050", &dependencies)
}

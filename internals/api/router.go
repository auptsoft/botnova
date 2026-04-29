package api

import (
	"time"

	"auptex.com/botnova/internals/api/handlers"
	"auptex.com/botnova/internals/api/middlewares"
	"auptex.com/botnova/internals/application/ports/dependencies"
	"auptex.com/botnova/internals/infrastructure/transport/websocket"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	_ "auptex.com/botnova/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(deps *dependencies.Dependencies, wsServer *websocket.Server) *gin.Engine {
	router := gin.Default()

	//Configure logger
	router.Use(ginzap.Ginzap(deps.ServiceLogger.GetZapLogger(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(deps.ServiceLogger.GetZapLogger(), true))

	//Add swagger docs endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Add scalar docs endpoint
	router.GET("/docs", handlers.ScalarDocsHandler)

	// Health check endpoint
	router.GET("/health", handlers.HealthHandler)

	userHandler := handlers.NewUserHandler(deps.UserService)

	//Api Routes
	apiRoutes := router.Group("/api")

	// Public user auth routes
	userPublicRoutes := apiRoutes.Group("/user")
	{
		userPublicRoutes.POST("/", userHandler.SignUp)
		userPublicRoutes.POST("/signup", userHandler.SignUp)
		userPublicRoutes.POST("/login", userHandler.Login)
	}

	// Protected API routes
	protectedRoutes := apiRoutes.Group("")
	protectedRoutes.Use(middlewares.AuthMiddleware())
	{
		protectedRoutes.GET("/ws", wsServer.HandleWebSocket)

		projectRoutes := protectedRoutes.Group("/project")
		{
			projectHandler := handlers.NewProjectHandler(deps.ProjectService)

			projectRoutes.GET("/", projectHandler.ListProjects)
			projectRoutes.POST("/", projectHandler.CreateProject)
			projectRoutes.GET("/:id", projectHandler.GetByID)
			projectRoutes.PUT("/", projectHandler.UpdateProject)
			projectRoutes.DELETE("/:id", projectHandler.Delete)
		}

		userProtectedRoutes := protectedRoutes.Group("/user")
		{
			userProtectedRoutes.GET("/me", userHandler.GetCurrentUser)
			userProtectedRoutes.PUT("/me", userHandler.UpdateCurrentUser)
			userProtectedRoutes.DELETE("/me", userHandler.DeleteCurrentUser)

			userProtectedRoutes.GET("/:id", userHandler.GetByID)
			userProtectedRoutes.PUT("/", userHandler.UpdateUser)
			userProtectedRoutes.DELETE("/:id", userHandler.Delete)
		}

		transportRoutes := protectedRoutes.Group("/transport")
		{
			transportHandler := handlers.NewTransportHandler(deps.TransportService)
			transportRoutes.POST("/websocket", transportHandler.SendToWebsocket)
		}
	}

	return router
}

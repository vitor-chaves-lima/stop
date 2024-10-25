package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vitor-chaves-lima/stop/cmd/web-api/controller"
	"github.com/vitor-chaves-lima/stop/cmd/web-api/middleware"
	"github.com/vitor-chaves-lima/stop/config"
	"github.com/vitor-chaves-lima/stop/internal/data/database"
)

var APIVersions = []string{"v1"}

type Dependencies struct {
	mongoDBManager *database.MongoDBManager
}

type APIControllers struct {
	docs *controller.Docs
	game *controller.Game
}

func setupDependencies(appConfig *config.Config) *Dependencies {
	mongoDbManager := database.NewMongoDBManager(context.Background(), appConfig.Database)

	return &Dependencies{
		mongoDBManager: mongoDbManager,
	}
}

func setupControllers() *APIControllers {
	openAPIDocuments, err := loadOpenAPIDocuments()
	if err != nil {
		panic(err)
	}

	docsController := controller.NewDocsController(openAPIDocuments)
	gameController := controller.NewGameController()

	apiControllers := &APIControllers{
		docs: docsController,
		game: gameController,
	}

	return apiControllers
}

func setupRoutes(apiRouter *gin.RouterGroup, apiControllers *APIControllers) {
	apiRouter.GET("/docs/:v/openapi.yaml", apiControllers.docs.HandleOpenAPIYAML)

	v1Router := apiRouter.Group("/v1")
	{
		v1Router.GET("/game/categories", apiControllers.game.HandleGetCategories)
		v1Router.POST("/game/session", apiControllers.game.HandleCreateSession)
	}
}

func main() {
	appConfig := config.LoadConfig()

	r := gin.Default()
	apiRouter := r.Group("/api")

	apiRouter.Use(middleware.ErrorHandler)
	apiRouter.Use(middleware.ResponseFormatter)

	dependencies := setupDependencies(appConfig)
	defer dependencies.mongoDBManager.Disconnect()

	apiControllers := setupControllers()
	setupRoutes(apiRouter, apiControllers)

	err := r.Run()
	if err != nil {
		return
	}
}

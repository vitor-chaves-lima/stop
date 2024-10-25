package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitor-chaves-lima/stop/cmd/web-api/controller"
	"github.com/vitor-chaves-lima/stop/cmd/web-api/middleware"
)

var APIVersions = []string{"v1"}

type APIControllers struct {
	docs *controller.Docs
	game *controller.Game
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
	r := gin.Default()
	apiRouter := r.Group("/api")

	apiRouter.Use(middleware.ErrorHandler)
	apiRouter.Use(middleware.ResponseFormatter)

	apiControllers := setupControllers()
	setupRoutes(apiRouter, apiControllers)

	err := r.Run()
	if err != nil {
		return
	}
}

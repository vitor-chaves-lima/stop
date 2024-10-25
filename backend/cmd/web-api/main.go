package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var APIVersions = []string{"v1"}

func main() {
	openAPIDocuments, err := loadOpenAPIDocuments()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/docs/:v/openapi.yaml", func(c *gin.Context) {
			c.String(http.StatusOK, *openAPIDocuments[c.Param("v")])
		})
	}

	r.Run()
}

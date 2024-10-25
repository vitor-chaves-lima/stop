package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Docs struct {
	openAPIDocuments map[string]*string
}

func (ctr *Docs) HandleOpenAPIYAML(c *gin.Context) {
	apiVersion := c.Param("v")

	openAPIDocument, ok := ctr.openAPIDocuments[apiVersion]
	if !ok {
		_ = c.AbortWithError(http.StatusNotFound, errors.New("openapi.yaml not found"))
		return
	}

	c.String(http.StatusOK, *openAPIDocument)
}

func NewDocsController(openApiDocuments map[string]*string) *Docs {
	return &Docs{
		openAPIDocuments: openApiDocuments,
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vitor-chaves-lima/stop/cmd/web-api/context"
)

type response struct {
	Data  any `json:"data,omitempty"`
	Meta  any `json:"meta,omitempty"`
	Links any `json:"links,omitempty"`
}

func ResponseFormatter(c *gin.Context) {
	c.Next()

	if !c.IsAborted() {
		statusCode, data, meta, links := context.GetResponse(c)
		c.JSON(statusCode, response{data, meta, links})
	}
}

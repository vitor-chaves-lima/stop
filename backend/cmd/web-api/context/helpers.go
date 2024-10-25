package context

import "github.com/gin-gonic/gin"

const (
	StatusCode = "StatusCode"
	Data       = "Data"
	Meta       = "Metadata"
	Links      = "Links"
)

func SetResponse(c *gin.Context, statusCode int, data any, meta any, links any) {
	c.Set(StatusCode, statusCode)
	c.Set(Data, data)
	c.Set(Meta, meta)
	c.Set(Links, links)
}

func GetResponse(c *gin.Context) (statusCode int, data interface{}, meta interface{}, links interface{}) {
	statusCode = c.GetInt(StatusCode)
	data, exists := c.Get(Data)
	if !exists {
		data = nil
	}

	meta, exists = c.Get(Meta)
	if !exists {
		meta = nil
	}

	links, exists = c.Get(Links)
	if !exists {
		links = nil
	}

	return statusCode, data, meta, links
}

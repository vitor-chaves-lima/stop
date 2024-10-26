package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/vitor-chaves-lima/stop/internal/logic"
)

type apiError struct {
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		formattedErrors := make([]apiError, len(c.Errors))

		for i, err := range c.Errors {
			var applicationError *logic.Error
			if errors.As(err, &applicationError) {
				formattedErrors[i] = apiError{
					Code:   applicationError.Code,
					Title:  "TODO",
					Detail: applicationError.Error(),
				}

				break
			}

			c.Status(c.Writer.Status())
			return
		}

		c.JSON(c.Writer.Status(), gin.H{"errors": formattedErrors})
	}
}

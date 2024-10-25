package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitor-chaves-lima/stop/cmd/web-api/context"
)

type Game struct {
}

func (ctr *Game) HandleGetCategories(c *gin.Context) {
	context.SetResponse(c, http.StatusOK, "Hello World", nil, nil)
}

func (ctr *Game) HandleCreateSession(c *gin.Context) {

}

func NewGameController() *Game {
	return &Game{}
}

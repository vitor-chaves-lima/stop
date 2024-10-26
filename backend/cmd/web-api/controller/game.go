package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitor-chaves-lima/stop/cmd/web-api/context"
	"github.com/vitor-chaves-lima/stop/internal/logic/services"
)

type Game struct {
	categoryService services.CategoryServiceInterface
}

func (ctr *Game) HandleGetCategories(c *gin.Context) {
	categories, paginationInfo, err := ctr.categoryService.ListCategories(c, struct {
		Page  int
		Limit int
	}{Page: 0, Limit: 10})

	if err != nil {
		//context.SetResponse(c, err.StatusCode, nil, nil, err)
		return
	}

	context.SetResponse(c, http.StatusOK, categories, paginationInfo, nil)
}

func (ctr *Game) HandleCreateSession(c *gin.Context) {

}

func NewGameController(categoryService *services.CategoryService) *Game {
	return &Game{
		categoryService: categoryService,
	}
}

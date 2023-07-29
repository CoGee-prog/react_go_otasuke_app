package controllers

import (
	"net/http"
	"react_go_otasuke_app/database"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*BaseController
}

// ユーザーコントローラーを返す
func NewUserController(db *database.GormDatabase) *UserController {
	return &UserController{
		BaseController: NewBaseController(db),
	}
}

func (uc *UserController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"OK",
		))
	}
}

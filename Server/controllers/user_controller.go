package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// ユーザーコントローラーを返す
func NewUserController() *UserController {
	return new(UserController)
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

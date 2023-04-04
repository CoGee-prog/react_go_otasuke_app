package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OpponentRecruitingController struct{}

func NewOpponentRecruitingController() *OpponentRecruitingController {
	return new(OpponentRecruitingController)
}

func (orc *OpponentRecruitingController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"OK",
		))
	}
}

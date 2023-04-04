package server

import (
	"react_go_otasuke_app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	router := gin.Default()
	v1 := router.Group("/")
	userController := controllers.NewUserController()
	opponentRecruitingController := controllers.NewOpponentRecruitingController()
	{
		v1.POST("/user", userController.Create())
	}
	{
		v1.POST("/opponent_recruiting", opponentRecruitingController.Create())
	}
	return router, nil
}

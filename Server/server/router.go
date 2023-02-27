package server

import (
	"react_go_otasuke_app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/")
	UserController := controllers.NewUserController()
	v1.GET("/", UserController.Create())
	return router
}
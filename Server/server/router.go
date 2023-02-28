package server

import (
	"react_go_otasuke_app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	router := gin.Default()
	v1 := router.Group("/")
	UserController := controllers.NewUserController()
	v1.GET("/", UserController.Create())
	return router, nil
}

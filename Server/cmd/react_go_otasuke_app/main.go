package main

import (
	"react_go_otasuke_app/Server/server"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.LoadHTMLGlob("web/templates/**/*.gtpl.html")
	r.Static("/web/assets/", "./web/assets/")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	r.Use((server.GinContextToContextMiddleware()))

	db.LoadEnv()
	db.Migrate()

	router.DefineRoutes(r, entclient.EntClient{})

	r.Run()
}
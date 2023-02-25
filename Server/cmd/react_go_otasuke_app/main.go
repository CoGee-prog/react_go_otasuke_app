package main

import (

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "success",
		})
	})
	// store := cookie.NewStore([]byte("secret"))
	// r.Use(sessions.Sessions("session", store))

	// db.LoadEnv()
	// db.Migrate()

	// router.DefineRoutes(r, entclient.EntClient{})

	r.Run()
}
package main

import (

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	r.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "success",
		})
	})

	r.Run()
}
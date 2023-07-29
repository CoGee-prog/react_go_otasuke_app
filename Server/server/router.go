package server

import (
	"react_go_otasuke_app/controllers"
	"react_go_otasuke_app/database"

	"github.com/gin-gonic/gin"
)

// ルーティング設定
func NewRouter() (*gin.Engine, error) {
	router := gin.Default()
	v1 := router.Group("/")
	// DIのためここでDBを取得する
	db := database.GetDB()
	gormDatabase := database.NewGormDatabase(db)
	// コントローラーを作成する
	userController := controllers.NewUserController(gormDatabase)
	opponentRecruitingController := controllers.NewOpponentRecruitingController(gormDatabase)
	{
		v1.POST("/user", userController.Create())
	}
	{
		v1.GET("/opponent_recruiting/:page", opponentRecruitingController.Index())
		v1.POST("/opponent_recruiting", opponentRecruitingController.Create())
		v1.DELETE("/opponent_recruiting/:id", opponentRecruitingController.Delete())
	}
	return router, nil
}

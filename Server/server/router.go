package server

import (
	"react_go_otasuke_app/controllers"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/middlewares"

	"github.com/gin-gonic/gin"
)

// ルーティング設定
func NewRouter() (*gin.Engine, error) {
	// DIのためここでDBを取得する
	db := database.GetDB()
	gormDatabase := database.NewGormDatabase(db)
	// コントローラーを作成する
	userController := controllers.NewUserController(gormDatabase)
	opponentRecruitingController := controllers.NewOpponentRecruitingController(gormDatabase)

	router := gin.Default()

	// 認証がいらないエンドポイント
	{
		router.GET("/opponent_recruitings/:page", opponentRecruitingController.Index())
	}

	// 認証が必要なエンドポイント
	authRequired := router.Group("/")

	middlewares.Init()
	authRequired.Use(middlewares.FirebaseAuthMiddleware())

	{
		authRequired.POST("/users", userController.Create())
		authRequired.POST("/opponent_recruitings", opponentRecruitingController.Create())
		authRequired.PATCH("/opponent_recruitings/:id", opponentRecruitingController.Update())
		authRequired.DELETE("/opponent_recruitings/:id", opponentRecruitingController.Delete())
	}
	return router, nil
}

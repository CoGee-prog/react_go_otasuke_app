package server

import (
	"react_go_otasuke_app/controllers"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/middlewares"
	"react_go_otasuke_app/services"

	"github.com/gin-gonic/gin"
)

// ルーティング設定
func NewRouter() (*gin.Engine, error) {
	// DIのためここでDBを取得する
	db := database.GetDB()
	router := gin.Default()
	// トランザクションを開始する
	router.Use(middlewares.Transaction(db))
	// DIのためここでサービスを作成する
	userTeamService := services.NewUserTeamService()
	userService := services.NewUserService()
	teamService := services.NewTeamService()
	opponentRecruitingService := services.NewOpponentRecruitingService(userTeamService)

	// コントローラーを作成する
	userController := controllers.NewUserController(userService)
	teamController := controllers.NewTeamController(userService, teamService)
	opponentRecruitingController := controllers.NewOpponentRecruitingController(opponentRecruitingService, userService)

	// CORSを設定
	setCors(router)

	// 認証がいらないエンドポイント
	{
		router.GET("/opponent_recruitings", opponentRecruitingController.Index())
		router.GET("/opponent_recruitings/:id", opponentRecruitingController.Get())
		router.POST("/login", userController.Login())
	}

	firebaseApp := userService.GetFireBaseApp()
	// 認証が必要なエンドポイント
	authRequired := router.Group("/")
	authRequired.Use(middlewares.AuthMiddleware(firebaseApp))

	{
		authRequired.POST("/logout", userController.Logout())
		authRequired.POST("/teams", teamController.Create())
		authRequired.POST("/opponent_recruitings", opponentRecruitingController.Create())
		authRequired.PATCH("/opponent_recruitings/:id", opponentRecruitingController.Update())
		authRequired.DELETE("/opponent_recruitings/:id", opponentRecruitingController.Delete())
	}

	return router, nil
}

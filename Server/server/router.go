package server

import (
	"react_go_otasuke_app/controllers"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/middlewares"
	"react_go_otasuke_app/repositories"
	"react_go_otasuke_app/services"

	"github.com/gin-gonic/gin"
)

// ルーティング設定
func NewRouter() (*gin.Engine, error) {
	// DIのためここでDBを取得
	db := database.GetDB()
	router := gin.Default()
	// トランザクションを開始
	router.Use(middlewares.Transaction(db))

	// リポジトリの作成
	userRepo := repositories.NewUserRepository()
	teamRepo := repositories.NewTeamRepository()
	userTeamRepo := repositories.NewUserTeamRepository()
	OpponentRecruitingRepo := repositories.NewOpponentRecruitingRepository()
	OpponentRecruitingCommentRepo := repositories.NewOpponentRecruitingCommentRepository()

	// サービスを作成
	userTeamService := services.NewUserTeamService(userTeamRepo)
	userService := services.NewUserService(userRepo, userTeamRepo)
	teamService := services.NewTeamService(teamRepo, userTeamRepo)
	opponentRecruitingService := services.NewOpponentRecruitingService(userTeamService, userRepo, OpponentRecruitingRepo, OpponentRecruitingCommentRepo)

	// コントローラーを作成
	userController := controllers.NewUserController(userService)
	teamController := controllers.NewTeamController(userService, teamService)
	opponentRecruitingController := controllers.NewOpponentRecruitingController(opponentRecruitingService, userService)
	opponentRecruitingCommentController := controllers.NewOpponentRecruitingCommentController(opponentRecruitingService, userService)

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
	authRequired.Use(middlewares.AuthMiddleware(firebaseApp, userService))

	{
		authRequired.POST("/logout", userController.Logout())
		authRequired.POST("/teams", teamController.Create())
		authRequired.POST("/opponent_recruitings", opponentRecruitingController.Create())
		authRequired.PATCH("/opponent_recruitings/:opponent_recruiting_id", opponentRecruitingController.Update())
		authRequired.PATCH("/opponent_recruitings/:opponent_recruiting_id/status", opponentRecruitingController.ChangeStatus())
		authRequired.DELETE("/opponent_recruitings/:opponent_recruiting_id", opponentRecruitingController.Delete())
		authRequired.GET("/opponent_recruitings/my_team", opponentRecruitingController.GetMyTeam())
		authRequired.POST("/opponent_recruitings/:opponent_recruiting_id/comments", opponentRecruitingCommentController.Create())
		authRequired.PATCH("/opponent_recruitings/:opponent_recruiting_id/comments/:comment_id", opponentRecruitingCommentController.Update())
		authRequired.DELETE("/opponent_recruitings/:opponent_recruiting_id/comments/:comment_id", opponentRecruitingCommentController.Delete())
	}

	return router, nil
}

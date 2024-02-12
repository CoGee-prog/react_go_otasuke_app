package middlewares

import (
	"net/http"
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

// Firebaseで認証を行うMiddleware関数
func AuthMiddleware(firebaseApp *firebase.App, db *database.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 開発環境の場合は認証をスキップする
		if config.Get().GetString("server.env") == "dev" {
			// ユーザーIDをセットする
			utils.SetUserID(c.GetHeader("x-user-id"))
			return
		}
		// クライアントから送信されたセッションCookieを取得
		cookie, err := c.Cookie("session")
		if err != nil {
			// セッションCookieが利用できない場合、認証エラー
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse(
				http.StatusUnauthorized,
				err.Error(),
				nil,
			))
			return
		}

		client, err := firebaseApp.Auth(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse(
				http.StatusUnauthorized,
				err.Error(),
				nil,
			))
			return
		}

		// セッションCookieの検証。ユーザーのFirebaseセッションが取り消されたかどうかもチェック
		decoded, err := client.VerifySessionCookieAndCheckRevoked(c, cookie)
		if err != nil {
			// セッションCookieが無効な場合、認証エラー
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse(
				http.StatusUnauthorized,
				err.Error(),
				nil,
			))
			return
		}
		// ユーザーサービスを取得する
		userService := services.NewUserService(db)
		userId := decoded.UID
		// ユーザーを取得する
		user, err := userService.GetUser(userId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}
		// ユーザーが見つからなければエラー
		if user == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// ユーザーIDをセットする
		utils.SetUserID(userId)

		// ユーザーが存在する場合はリクエストを続ける
		c.Next()
	}
}

package middlewares

import (
	"net/http"
	"react_go_otasuke_app/controllers"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/services"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)


// Firebaseで認証を行うMiddleware関数
func AuthMiddleware(firebaseApp *firebase.App, db *database.GormDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// クライアントから送信されたセッションCookieを取得
		cookie, err := c.Cookie("session")
		if err != nil {
			// セッションCookieが利用できない場合、ユーザーにログインを強制する
			c.Redirect(http.StatusFound, "/login")
			return
		}

		client, err := firebaseApp.Auth(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error getting Auth client"})
			return
		}

		// セッションCookieの検証。ユーザーのFirebaseセッションが取り消されたかどうかもチェック
		decoded, err := client.VerifySessionCookieAndCheckRevoked(c, cookie)
		if err != nil {
			// セッションCookieが無効な場合、ユーザーにログインを強制する
			c.Redirect(http.StatusFound, "/login")
			return
		}
		// ユーザーサービスを取得する
		userService := services.NewUserService(db)
		// ユーザーを取得する
		user, err :=  userService.GetUser(decoded.UID)
			if err != nil {
			c.JSON(http.StatusServiceUnavailable, controllers.NewResponse(
				http.StatusServiceUnavailable,
				err.Error(),
				"",
			))
			return
		}
		// ユーザーが見つからなければエラー
		if user == nil {
			c.JSON(http.StatusBadRequest, controllers.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				"",
			))
		}

		// ユーザーが存在する場合はリクエストを続ける
		c.Next()
	}
}

package middlewares

import (
	"net/http"
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Firebaseで認証を行うMiddleware関数
func AuthMiddleware(firebaseApp *firebase.App, userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// ローカル環境の場合は認証をスキップする
		if config.Get().GetString("server.env") == "local" {
			// ユーザーIDをセットする
			c.Set("userId", c.GetHeader("X-User-Id"))
			return
		}
		// クライアントから送信されたセッションCookieを取得
		cookie, err := c.Cookie("session")
		if err != nil {
			// セッションCookieが利用できない場合、認証エラー
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse(
				http.StatusUnauthorized,
				"認証に失敗しました",
				nil,
			))
			return
		}

		client, err := firebaseApp.Auth(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse(
				http.StatusUnauthorized,
				"認証に失敗しました",
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
				"認証に失敗しました",
				nil,
			))
			return
		}

		tx := c.MustGet("tx").(*gorm.DB)

		userId := decoded.UID
		// ユーザーを取得する
		user, err := userService.GetUserWithCurrentTeam(tx, userId)
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse(
				http.StatusUnauthorized,
				"ユーザーが見つかりません",
				nil,
			))
			return
		}

		// ユーザーIDをセットする
		c.Set("userId", userId)

		// ユーザーが存在する場合はリクエストを続ける
		c.Next()
	}
}

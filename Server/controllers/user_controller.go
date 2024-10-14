package controllers

import (
	"net/http"
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"
	"react_go_otasuke_app/views"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserService *services.UserService
}

// ユーザーコントローラーを返す
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

type loginResponse struct {
	User *views.UserView `json:"user"`
}

func (uc *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := c.MustGet("tx").(*gorm.DB)
		// 開発環境の場合はIDトークン検証をスキップしてユーザーを作成する
		if config.Get().GetString("server.env") == "local" {
			// ユーザーデータを検索
			devUser, err := uc.UserService.GetUser(tx, c.GetHeader("X-User-Id"))
			if err != nil {
				c.JSON(http.StatusServiceUnavailable, utils.NewResponse(
					http.StatusServiceUnavailable,
					err.Error(),
					nil,
				))
				return
			}
			// ユーザーデータがなければ作成
			if devUser == nil {
				devUser = &models.User{
					ID:   c.GetHeader("X-User-Id"),
					Name: "dev-user",
				}
				// ユーザーデータを作成
				if err := uc.UserService.CreateUser(tx, devUser); err != nil {
					c.JSON(http.StatusBadRequest, utils.NewResponse(
						http.StatusBadRequest,
						err.Error(),
						nil,
					))
					return
				}
			} else {
				// ユーザーデータがある場合はユーザーの現在のチームがあれば役割も設定する
				if devUser.CurrentTeam != nil {
					var userTeam models.UserTeam
					if err := tx.Where("user_id = ? AND team_id = ?", devUser.ID, devUser.CurrentTeamId).First(&userTeam).Error; err == nil {
						devUser.CurrentTeamRole = &userTeam.Role
					}
				}
			}

			c.JSON(http.StatusOK, utils.NewResponse(
				http.StatusOK,
				http.StatusText(http.StatusOK),
				&loginResponse{
					User: views.CreateUserView(devUser),
				},
			))
			return
		}

		// IDトークンを検証
		token, err := services.VerifyIDToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.NewResponse(
				http.StatusUnauthorized,
				err.Error(),
				nil,
			))
			return
		}

		// ユーザーデータを検索
		user, err := uc.UserService.GetUser(tx, token.UID)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, utils.NewResponse(
				http.StatusServiceUnavailable,
				err.Error(),
				nil,
			))
			return
		}

		// ユーザーデータがなければ作成
		if user == nil {
			name, _ := token.Claims["name"].(string)
			user = &models.User{
				ID:   token.UID,
				Name: name,
			}
			// ユーザーデータを作成
			if err := uc.UserService.CreateUser(tx, user); err != nil {
				c.JSON(http.StatusBadRequest, utils.NewResponse(
					http.StatusBadRequest,
					err.Error(),
					nil,
				))
				return
			}
		} else {
			// ユーザーデータがある場合はユーザーの現在のチームがあれば役割も設定する
			if user.CurrentTeam != nil {
				if userTeam, err := uc.UserService.GetUserTeam(tx, user.ID, *user.CurrentTeamId); err == nil {
					user.CurrentTeamRole = &userTeam.Role
				}
			}
		}

		// セッションを作成
		if err := services.CreateSessionCookie(c); err != nil {
			c.JSON(http.StatusInternalServerError, utils.NewResponse(
				http.StatusInternalServerError,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"ログインしました",
			&loginResponse{
				User: views.CreateUserView(user),
			},
		))
	}
}

func (uc *UserController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		uc.UserService.RevokeRefreshTokens(c)
		conf := config.Get()
		// クッキーを削除するレスポンスを設定
		c.SetCookie("session", "", -1, "/", conf.GetString("client.domain"), true, true)

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"ログアウトしました",
			[]string{},
		))
	}
}

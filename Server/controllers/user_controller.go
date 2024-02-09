package controllers

import (
	"net/http"
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"
	"react_go_otasuke_app/views"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	*BaseController
}

var userService *services.UserService

// ユーザーコントローラーを返す
func NewUserController(db *database.GormDatabase) (*UserController, *firebase.App) {
	userService = services.NewUserService(db)
	return &UserController{
		BaseController: NewBaseController(db),
	}, userService.GetFireBaseApp()
}

type loginResponse struct {
	User *views.UserView `json:"user"`
}

func (uc *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		user, err := userService.GetUser(token.UID)
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
				Id:   token.UID,
				Name: name,
			}
			// ユーザーデータを作成
			if err := userService.CreateUser(user); err != nil {
				c.JSON(http.StatusBadRequest, utils.NewResponse(
					http.StatusBadRequest,
					err.Error(),
					nil,
				))
				return
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
			http.StatusText(http.StatusOK),
			&loginResponse{
				User: views.CreateUserView(user),
			},
		))
	}
}

func (uc *UserController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		userService.RevokeRefreshTokens(c)
		conf := config.Get()
		// クッキーを削除するレスポンスを設定
		c.SetCookie("session", "", -1, "/", conf.GetString("client.domain"), true, true)

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			[]string{},
		))
	}
}

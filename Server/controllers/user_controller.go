package controllers

import (
	"net/http"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
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
	userService := services.NewUserService(db)
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
			c.JSON(http.StatusUnauthorized, NewResponse(
				http.StatusUnauthorized,
				err.Error(),
				"",
			))
			return
		}

		// ユーザーデータを検索
		user, err := userService.GetUser(token.UID)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, NewResponse(
				http.StatusServiceUnavailable,
				err.Error(),
				"",
			))
			return
		}

		newUser := &models.User{}
		// ユーザーデータがなければ作成
		if user == nil {
			name, _ := token.Claims["name"].(string)
			newUser = &models.User{
				Id:   token.UID,
				Name: name,
			}
			// ユーザーデータを作成
			if err := userService.CreateUser(newUser); err != nil {
				c.JSON(http.StatusBadRequest, NewResponse(
					http.StatusBadRequest,
					err.Error(),
					"",
				))
				return
			}
		}

		// セッションを作成
		if err := services.CreateSessionCookie(c); err != nil {
			c.JSON(http.StatusInternalServerError, NewResponse(
				http.StatusInternalServerError,
				err.Error(),
				"",
			))
			return
		}

		c.JSON(http.StatusOK, NewResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			&loginResponse{
				User: views.CreateUserView(newUser),
			},
		))
	}
}

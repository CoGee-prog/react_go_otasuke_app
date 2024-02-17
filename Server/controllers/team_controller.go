package controllers

import (
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	UserService *services.UserService
	TeamService *services.TeamService
}

// チームコントローラーを返す
func NewTeamController(userService *services.UserService, teamService *services.TeamService) *TeamController {
	return &TeamController{
		UserService: userService,
		TeamService: teamService,
	}
}

func (tc *TeamController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		team := &models.Team{}

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(team); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
			  "不正なリクエストです",
				nil,
			))
			return
		}

		// リクエストのバリデーションチェック
		if err := team.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// チームを作成する
		if err := tc.TeamService.CreateTeam(team); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"チーム作成に失敗しました",
				nil,
			))
			return
		}

		// ユーザーの現在のチームを作成したチームに変更する
		if err := tc.UserService.UpdateCurrentTeam(team.ID); err != nil{
				c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"チームを作成しました",
			nil,
		))
	}
}


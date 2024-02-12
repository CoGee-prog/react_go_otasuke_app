package controllers

import (
	"net/http"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	*BaseController
}

var teamService *services.TeamService

// チームコントローラーを返す
func NewTeamController(db *database.GormDatabase) (*TeamController) {
	teamService = services.NewTeamService(db)
	return &TeamController{
		BaseController: NewBaseController(db),
	}
}

func (tc *TeamController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		team := &models.Team{}

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(team); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
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
		if err := teamService.CreateTeam(team); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"チーム作成に失敗しました",
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


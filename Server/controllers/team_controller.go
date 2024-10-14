package controllers

import (
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

type TeamCreateResponse struct {
	CurrentTeamId   uint            `json:"current_team_id"`
	CurrentTeamName string          `json:"current_team_name"`
	CurrentTeamRole models.TeamRole `json:"current_team_role"`
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

		tx := c.MustGet("tx").(*gorm.DB)
		userId := c.MustGet("userId").(string)
		// チームを作成する
		if err := tc.TeamService.CreateTeam(tx, userId, team); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"チーム作成に失敗しました",
				nil,
			))
			return
		}
		// ユーザーの現在のチームを作成したチームに変更する
		if err := tc.UserService.UpdateCurrentTeam(tx, userId, team.ID); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// ユーザーチームを取得する
		userTeam, _ := tc.UserService.GetUserTeam(tx, userId, team.ID)

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"チームを作成しました",
			&TeamCreateResponse{
				CurrentTeamId:   team.ID,
				CurrentTeamName: team.Name,
				CurrentTeamRole: userTeam.Role,
			},
		))
	}
}

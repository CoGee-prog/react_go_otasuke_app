package controllers

import (
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeamController struct {
	UserService services.UserService
	TeamService services.TeamService
}

// チームコントローラーを返す
func NewTeamController(userService services.UserService, teamService services.TeamService) *TeamController {
	return &TeamController{
		UserService: userService,
		TeamService: teamService,
	}
}

type TeamGetResponse struct {
	Name         string              `json:"name"`
	PrefectureId models.PrefectureID `json:"prefecture_id"`
	LevelId      models.TeamLevelId  `json:"level_id"`
	HomePageUrl  *string             `json:"home_page_url"`
	Other        *string             `json:"other"`
}

func (tc *TeamController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		tx := c.MustGet("tx").(*gorm.DB)

		team, err := tc.TeamService.GetTeam(tx, uint(id))

		if err != nil || team == nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"チームが見つかりません",
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"",
			&TeamGetResponse{
				Name:         team.Name,
				PrefectureId: team.PrefectureId,
				LevelId:      team.LevelId,
				HomePageUrl:  team.HomePageUrl,
				Other:        team.Other,
			},
		))
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
		if err := tc.TeamService.CreateTeamWithAdmin(tx, userId, team); err != nil {
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

func (tc *TeamController) Update() gin.HandlerFunc {
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

		teamId, _ := strconv.Atoi(c.Param("team_id"))
		tx := c.MustGet("tx").(*gorm.DB)

		// リクエストのバリデーションチェック
		if err := team.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		userId := c.MustGet("userId").(string)

		// チームを更新
		if err := tc.TeamService.UpdateTeam(tx, userId, uint(teamId), team); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"",
			&TeamGetResponse{
				Name:         team.Name,
				PrefectureId: team.PrefectureId,
				LevelId:      team.LevelId,
				HomePageUrl:  team.HomePageUrl,
				Other:        team.Other,
			},
		))
	}
}
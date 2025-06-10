package controllers

import (
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"
	"react_go_otasuke_app/views"
	"time"

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

type TeamGetAndUpdateResponse struct {
	Team *views.TeamView `json:"team"`
}

func (tc *TeamController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("team_id")
		tx := c.MustGet("tx").(*gorm.DB)

		team, err := tc.TeamService.GetTeam(tx, id)

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
			&TeamGetAndUpdateResponse{
				Team: views.CreateTeamView(*team),
			},
		))
	}
}

type TeamCreateResponse struct {
	CurrentTeamId   string          `json:"current_team_id"`
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

		teamId := c.Param("team_id")
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
		if err := tc.TeamService.UpdateTeam(tx, userId, teamId, team); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"チームを更新しました",
			&TeamGetAndUpdateResponse{
				Team: views.CreateTeamView(*team),
			},
		))
	}
}

type CreateInviteTokenResponse struct {
	InviteToken string `json:"invite_token"`
	ExpiresAt   string `json:"expires_at"`
}

func (tc *TeamController) CreateInviteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("team_id")
		tx := c.MustGet("tx").(*gorm.DB)

		team, err := tc.TeamService.GetTeam(tx, id)

		if err != nil || team == nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"チームが見つかりません",
				nil,
			))
			return
		}

		userId := c.MustGet("userId").(string)
		// 招待トークンを作成する
		teamInvite, err := tc.TeamService.CreateInviteToken(tx, userId, team)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"招待トークンを発行しました",
			&CreateInviteTokenResponse{
				InviteToken: teamInvite.Token,
				ExpiresAt:   teamInvite.ExpiresAt.Format(time.RFC3339),
			},
		))
	}
}

type TeamInviteResponse struct {
	Team *views.TeamView `json:"team"`
}

func (tc *TeamController) GetByInviteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("invite_token")
		tx := c.MustGet("tx").(*gorm.DB)

		team, err := tc.TeamService.GetTeamByInviteToken(tx, token)
		if err != nil || team == nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			&TeamInviteResponse{
				Team: views.CreateTeamView(*team),
			},
		))
	}
}

func (tc *TeamController) JoinByInviteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("invite_token")
		tx := c.MustGet("tx").(*gorm.DB)
		userId := c.MustGet("userId").(string)

		team, err := tc.TeamService.JoinTeamByInviteToken(tx, userId, token)
		if err != nil || team == nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		user, _ := tc.UserService.GetUserWithCurrentTeam(tx, userId)
		if user != nil && user.CurrentTeamId == nil {
			tc.UserService.UpdateCurrentTeam(tx, userId, team.ID)
		}
		userTeam, _ := tc.UserService.GetUserTeam(tx, userId, team.ID)

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"チームに参加しました",
			&TeamCreateResponse{
				CurrentTeamId:   team.ID,
				CurrentTeamName: team.Name,
				CurrentTeamRole: userTeam.Role,
			},
		))
	}
}

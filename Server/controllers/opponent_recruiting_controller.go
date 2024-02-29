package controllers

import (
	"errors"
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"
	"react_go_otasuke_app/views"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OpponentRecruitingController struct {
	OpponentRecruitingService *services.OpponentRecruitingService
	UserService               *services.UserService
}

// 対戦相手募集のコントローラーを作成する
func NewOpponentRecruitingController(opponentRecruitingService *services.OpponentRecruitingService, userService *services.UserService) *OpponentRecruitingController {
	return &OpponentRecruitingController{
		OpponentRecruitingService: opponentRecruitingService,
		UserService:               userService,
	}
}

type OpponentRecruitingIndexResponse struct {
	OpponentRecruitings []*views.OpponentRecruitingView `json:"opponent_recruitings"`
	Page                *models.Page                    `json:"page"`
}

func (oc *OpponentRecruitingController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		// データを取得する
		opponentRecruitings, page := oc.OpponentRecruitingService.GetOpponentRecruitingList(c)

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			&OpponentRecruitingIndexResponse{
				OpponentRecruitings: views.CreateOpponentRecruitingView(opponentRecruitings),
				Page:                page,
			},
		))
	}
}

type OpponentRecruitingCreateRequest struct {
	Title        string              `json:"title"`
	HasGround    bool                `json:"has_ground"`
	GroundName   string              `json:"ground_name"`
	PrefectureId models.PrefectureId `json:"prefecture_id" binding:"required"`
	StartTime    time.Time           `json:"start_time" binding:"required"`
	EndTime      time.Time           `json:"end_time" binding:"required"`
	Detail       string              `json:"detail"`
}

func (oc *OpponentRecruitingController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request OpponentRecruitingCreateRequest

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"不正なリクエストです",
				nil,
			))
			return
		}

		db := c.MustGet("tx").(*gorm.DB)
		// ユーザーを取得する
		user, err := oc.UserService.GetUser(db, c.MustGet("userId").(string))
		if err != nil || user == nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				errors.New("ユーザーを取得できません").Error(),
				nil,
			))
			return
		}

		// 対戦相手募集の構造体を作成
		opponentRecruiting := &models.OpponentRecruiting{
			TeamId:       *user.CurrentTeamId,
			PrefectureId: request.PrefectureId,
			StartTime:    request.StartTime,
			EndTime:      request.EndTime,
			Detail:       request.Detail,
		}

		// リクエストのバリデーションチェック
		if err := opponentRecruiting.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}
		userId := c.MustGet("userId").(string)
		// データを作成する
		if err := oc.OpponentRecruitingService.CreateOpponentRecruiting(db, userId, opponentRecruiting); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"対戦相手募集を作成しました",
			nil,
		))
	}
}

type OpponentRecruitingUpdateRequest struct {
	PrefectureId models.PrefectureId `json:"prefecture_id"`
	DateTime     time.Time           `json:"date_time"`
	Detail       string              `json:"detail"`
}

func (oc *OpponentRecruitingController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &OpponentRecruitingUpdateRequest{}

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(request); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"不正なリクエストです",
				nil,
			))
			return
		}

		// 対戦相手募集の構造体を作成
		opponentRecruiting := &models.OpponentRecruiting{
			PrefectureId: request.PrefectureId,
			StartTime:    request.DateTime,
			Detail:       request.Detail,
		}
		// リクエストのバリデーションチェック
		if err := opponentRecruiting.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))
		db := c.MustGet("tx").(*gorm.DB)
		userId := c.MustGet("userId").(string)
		// データを更新する
		if err := oc.OpponentRecruitingService.UpdateOpponentRecruiting(db, userId, opponentRecruiting, uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"対戦相手募集を更新しました",
			nil,
		))
	}

}

func (oc *OpponentRecruitingController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		db := c.MustGet("tx").(*gorm.DB)
		userId := c.MustGet("userId").(string)
		// データを削除する
		if err := oc.OpponentRecruitingService.DeleteOpponentRecruiting(db, userId, uint(id)); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			"対戦相手募集を削除しました",
			nil,
		))
	}
}

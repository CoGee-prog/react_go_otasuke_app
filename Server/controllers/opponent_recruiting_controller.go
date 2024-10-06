package controllers

import (
	"errors"
	"net/http"
	"react_go_otasuke_app/database"
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
	OpponentRecruitings []*views.OpponentRecruitingIndexView `json:"opponent_recruitings"`
	Page                *database.Page                       `json:"page"`
}

func (oc *OpponentRecruitingController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		// データを取得する
		opponentRecruitings, page := oc.OpponentRecruitingService.GetOpponentRecruitingList(c, false)

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			&OpponentRecruitingIndexResponse{
				OpponentRecruitings: views.CreateOpponentRecruitingIndexView(opponentRecruitings),
				Page:                page,
			},
		))
	}
}

type OpponentRecruitingGetResponse struct {
	OpponentRecruiting *views.OpponentRecruitingGetView `json:"opponent_recruiting"`
}

func (oc *OpponentRecruitingController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		db := c.MustGet("tx").(*gorm.DB)
		// データを取得する
		opponentRecruiting, err := oc.OpponentRecruitingService.FindOpponentRecruitingWithComment(db, uint(id))
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
			"",
			&OpponentRecruitingGetResponse{
				OpponentRecruiting: views.CreateOpponentRecruitingGetView(opponentRecruiting),
			},
		))
	}
}

func (oc *OpponentRecruitingController) GetMyTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 自チームの対戦相手募集データを取得する
		opponentRecruitings, page := oc.OpponentRecruitingService.GetOpponentRecruitingList(c, true)

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			&OpponentRecruitingIndexResponse{
				OpponentRecruitings: views.CreateOpponentRecruitingIndexView(opponentRecruitings),
				Page:                page,
			},
		))
	}
}

type OpponentRecruitingCreateRequest struct {
	Title        string              `json:"title" binding:"required"`
	HasGround    bool                `json:"has_ground"`
	GroundName   *string             `json:"ground_name"`
	PrefectureId models.PrefectureID `json:"prefecture_id" binding:"required"`
	StartTime    time.Time           `json:"start_time" binding:"required"`
	EndTime      time.Time           `json:"end_time" binding:"required"`
	Detail       *string             `json:"detail"`
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
		user, err := oc.UserService.GetUser(c.MustGet("userId").(string))
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
			Title:        request.Title,
			HasGround:    request.HasGround,
			GroundName:   *request.GroundName,
			TeamID:       *user.CurrentTeamId,
			PrefectureID: request.PrefectureId,
			StartTime:    request.StartTime,
			EndTime:      request.EndTime,
			Detail:       *request.Detail,
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
		// データを作成する
		if err := oc.OpponentRecruitingService.CreateOpponentRecruiting(db, user.ID, opponentRecruiting); err != nil {
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
	Title        string              `json:"title" binding:"required"`
	HasGround    bool                `json:"has_ground"`
	GroundName   *string             `json:"ground_name"`
	PrefectureId models.PrefectureID `json:"prefecture_id" binding:"required"`
	StartTime    time.Time           `json:"start_time" binding:"required"`
	EndTime      time.Time           `json:"end_time" binding:"required"`
	Detail       *string             `json:"detail"`
}

func (oc *OpponentRecruitingController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request OpponentRecruitingUpdateRequest

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
		user, err := oc.UserService.GetUser(c.MustGet("userId").(string))
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
			Title:        request.Title,
			HasGround:    request.HasGround,
			GroundName:   *request.GroundName,
			TeamID:       *user.CurrentTeamId,
			PrefectureID: request.PrefectureId,
			StartTime:    request.StartTime,
			EndTime:      request.EndTime,
			Detail:       *request.Detail,
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

		opponentRecruitingId, _ := strconv.Atoi(c.Param("opponent_recruiting_id"))
		// データを更新する
		if err := oc.OpponentRecruitingService.UpdateOpponentRecruiting(db, user.ID, opponentRecruiting, uint(opponentRecruitingId)); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// データを取得する
		opponentRecruitingWithComment, err := oc.OpponentRecruitingService.FindOpponentRecruitingWithComment(db, uint(opponentRecruitingId))
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
			"対戦相手募集を更新しました",
			&OpponentRecruitingGetResponse{
				OpponentRecruiting: views.CreateOpponentRecruitingGetView(opponentRecruitingWithComment),
			},
		))
	}

}

func (oc *OpponentRecruitingController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("opponent_recruiting_id"))
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

type OpponentRecruitingChangeStatusRequest struct {
	IsActive *bool `json:"is_active" binding:"required"`
}

func (oc *OpponentRecruitingController) ChangeStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &OpponentRecruitingChangeStatusRequest{}

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(request); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"不正なリクエストです",
				nil,
			))
			return
		}

		// isActiveがない場合はエラー
		if request.IsActive == nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"不正なリクエストです",
				nil,
			))
			return
		}

		// 対戦相手募集の構造体を作成
		opponentRecruiting := &models.OpponentRecruiting{
			IsActive: *request.IsActive,
		}

		opponentRecruitingId, _ := strconv.Atoi(c.Param("opponent_recruiting_id"))

		db := c.MustGet("tx").(*gorm.DB)
		userId := c.MustGet("userId").(string)
		// データを更新する
		if err := oc.OpponentRecruitingService.UpdateStatusOpponentRecruiting(db, userId, opponentRecruiting, uint(opponentRecruitingId)); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		// レスポンスメッセージを設定
		var message string
		if *request.IsActive {
			message = "対戦相手募集を再開しました"
		} else {
			message = "対戦相手募集を終了しました"
		}

		// データを取得する
		opponentRecruiting, err := oc.OpponentRecruitingService.FindOpponentRecruitingWithComment(db, uint(opponentRecruitingId))
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
			message,
			&OpponentRecruitingGetResponse{
				OpponentRecruiting: views.CreateOpponentRecruitingGetView(opponentRecruiting),
			},
		))
	}

}

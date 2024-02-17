package controllers

import (
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"
	"react_go_otasuke_app/views"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OpponentRecruitingController struct {
	OpponentRecruitingService *services.OpponentRecruitingService
}

// 対戦相手募集のコントローラーを作成する
func NewOpponentRecruitingController(opponentRecruitingService *services.OpponentRecruitingService) *OpponentRecruitingController {
	return &OpponentRecruitingController{
		OpponentRecruitingService: opponentRecruitingService,
	}
}

type indexResponse struct {
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
			&indexResponse{
				OpponentRecruitings: views.CreateOpponentRecruitingView(opponentRecruitings),
				Page:                page,
			},
		))
	}
}

func (oc *OpponentRecruitingController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		opponentRecruiting := &models.OpponentRecruiting{}

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(opponentRecruiting); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
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
		if err := oc.OpponentRecruitingService.CreateOpponentRecruiting(opponentRecruiting); err != nil {
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

type UpdateRequest struct {
	AreaId   uint      `json:"area_id" gorm:"type:int; not null"`
	DateTime time.Time `json:"date_time"`
	Detail   *string   `json:"detail" gorm:"type:text"`
}

func (oc *OpponentRecruitingController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &UpdateRequest{}

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(request); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				err.Error(),
				nil,
			))
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))
		// 対戦相手募集の構造体を作成
		opponentRecruiting := &models.OpponentRecruiting{
			AreaId:   request.AreaId,
			DateTime: request.DateTime,
			Detail:   request.Detail,
		}

		// データを更新する
		err := oc.OpponentRecruitingService.UpdateOpponentRecruiting(opponentRecruiting, uint(id))
		// エラーが起きているかどうか
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
			nil,
		))
	}

}

func (oc *OpponentRecruitingController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		// データを削除する
		if err := oc.OpponentRecruitingService.DeleteOpponentRecruiting(uint(id)); err != nil {
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

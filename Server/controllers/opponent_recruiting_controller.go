package controllers

import (
	"net/http"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/utils"
	"react_go_otasuke_app/views"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OpponentRecruitingController struct {
	*BaseController
}

var opponentRecruitingService *services.OpponentRecruitingService

// 対戦相手募集のコントローラーを作成する
func NewOpponentRecruitingController(db *database.GormDatabase) *OpponentRecruitingController {
	opponentRecruitingService = services.NewOpponentRecruitingService(db)
	return &OpponentRecruitingController{
		BaseController: NewBaseController(db),
	}
}

type indexResponse struct {
	OpponentRecruitings []*views.OpponentRecruitingView `json:"opponent_recruitings"`
	Page                *models.Page                    `json:"page"`
}

type updateRequest struct {
	AreaId   int       `json:"area_id" gorm:"type:int; not null"`
	DateTime time.Time `json:"date_time"`
	Detail   *string   `json:"detail" gorm:"type:text"`
}

func (oc *OpponentRecruitingController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		// データを取得する
		opponentRecruitings, page := opponentRecruitingService.GetOpponentRecruitingList(c)

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
		if err := oc.db.DB.Create(opponentRecruiting).Error; err != nil {
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

func (oc *OpponentRecruitingController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &updateRequest{}

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
		// データを更新する
		result := oc.db.DB.Model(&models.OpponentRecruiting{}).Where("id = ?", id).Updates(request)
		// エラーが起きているかどうか
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				result.Error.Error(),
				nil,
			))
			return
		}

		// 更新したデータが0件の場合はエラー
		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"更新対象のデータがありません",
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			nil,
		))
	}

}

func (oc *OpponentRecruitingController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		// データを削除する
		result := oc.db.DB.Unscoped().Delete(&models.OpponentRecruiting{}, "id = ?", id)
		// エラーが起きているかどうか
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				result.Error.Error(),
				nil,
			))
			return
		}

		// 削除したデータが0件の場合はエラー
		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				http.StatusBadRequest,
				"削除対象のデータがありません",
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, utils.NewResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"OK",
		))
	}
}

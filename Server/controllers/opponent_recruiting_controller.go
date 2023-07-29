package controllers

import (
	"net/http"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/views"

	"github.com/gin-gonic/gin"
)

type OpponentRecruitingController struct {
	*BaseController
}

// 対戦相手募集のコントローラーを作成する
func NewOpponentRecruitingController(db *database.GormDatabase) *OpponentRecruitingController {
	return &OpponentRecruitingController{
		BaseController: NewBaseController(db),
	}
}

type indexResponse struct {
	OpponentRecruitings []*views.OpponentRecruitingView `json:"opponent_recruitings"`
	Page                *models.Page                    `json:"page"`
}

func (oc *OpponentRecruitingController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		opponentRecruitingService := services.NewOpponentRecruitingService(oc.db)
		// データを取得する
		opponentRecruitings, page := opponentRecruitingService.GetOpponentRecruitingList(c)

		c.JSON(http.StatusOK, newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			&indexResponse{
				OpponentRecruitings: views.IndexOpponentRecruitingView(opponentRecruitings),
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
			c.JSON(http.StatusBadRequest, newResponse(
				http.StatusBadRequest,
				err.Error(),
				"",
			))
			return
		}

		// リクエストのバリデーションチェック
		if err := opponentRecruiting.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, newResponse(
				http.StatusBadRequest,
				err.Error(),
				"",
			))
			return
		}

		// データを作成する
		if err := oc.db.DB.Create(opponentRecruiting).Error; err != nil {
			c.JSON(http.StatusBadRequest, newResponse(
				http.StatusBadRequest,
				err.Error(),
				"",
			))
			return
		}

		c.JSON(http.StatusOK, newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"OK",
		))
	}
}

func (oc *OpponentRecruitingController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

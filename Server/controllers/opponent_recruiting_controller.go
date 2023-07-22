package controllers

import (
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/services"
	"react_go_otasuke_app/views"

	"github.com/gin-gonic/gin"
)

type OpponentRecruitingController struct{}

// 対戦相手募集のコントローラーを返す
func NewOpponentRecruitingController() *OpponentRecruitingController {
	return new(OpponentRecruitingController)
}

type indexResponse struct {
	OpponentRecruitings []*views.OpponentRecruitingView `json:"opponent_recruitings"`
	Page                *models.Page                    `json:"page"`
}

func (oc *OpponentRecruitingController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		// データを取得する
		opponentRecruitings, page := services.GetOpponentRecruitingList(c)

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
		if err := opponentRecruiting.Create(); err != nil {
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

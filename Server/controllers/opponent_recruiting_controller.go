package controllers

import (
	"net/http"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/views"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OpponentRecruitingController struct{}

func NewOpponentRecruitingController() *OpponentRecruitingController {
	return new(OpponentRecruitingController)
}

type indexResponse struct {
	OpponentRecruitings []*views.OpponentRecruitingView `json:"opponent_recruitings"`
}

func (oc *OpponentRecruitingController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNumber, _ := strconv.Atoi(c.Param("page"))
		page := &models.Page{
			Number: pageNumber,
			Size:   10,
		}

		opponentRecruiting := &models.OpponentRecruiting{}

		// データを取得する
		opponentRecruitings := opponentRecruiting.GetByPagination(page)

		c.JSON(http.StatusOK, newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			&indexResponse{
				OpponentRecruitings: views.IndexOpponentRecruitingView(opponentRecruitings),
			},
		))
	}
}

func (oc *OpponentRecruitingController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		opponentRecruiting := &models.OpponentRecruiting{}

		// // リクエストパラメーターをバインドする
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

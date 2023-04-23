package controllers

import (
	"net/http"
	"react_go_otasuke_app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OpponentRecruitingController struct {
	OpponentRecruiting *models.OpponentRecruiting
}

func NewOpponentRecruitingController() *OpponentRecruitingController {
	return new(OpponentRecruitingController)
}

func (oc *OpponentRecruitingController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNumber, _ := strconv.Atoi(c.Param("page"))
		page := &models.Page{
			Number: pageNumber,
			Size:   10,
		}

		// データを取得する
		opponentRecruitings := oc.OpponentRecruiting.GetByPagination(page).Value

		c.JSON(http.StatusOK, newResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			opponentRecruitings,
		))

	}
}

func (oc *OpponentRecruitingController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		// リクエストパラメーターをバインドする
		if err := c.ShouldBindJSON(&oc.OpponentRecruiting); err != nil {
			c.JSON(http.StatusBadRequest, newResponse(
				http.StatusBadRequest,
				err.Error(),
				"",
			))
			return
		}

		// リクエストのバリデーションチェック
		if err := oc.OpponentRecruiting.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, newResponse(
				http.StatusBadRequest,
				err.Error(),
				"",
			))
			return
		}

		// データを作成する
		if err := oc.OpponentRecruiting.Create(); err != nil {
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

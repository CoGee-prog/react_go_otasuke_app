package services

import (
	"math"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 対戦相手募集の構造体の配列
var opponentRecruitings []*models.OpponentRecruiting

// リスト表示時の1ページあたりの要素数
var pageSize int = 10

// 対戦相手募集のリストとページ情報を返す
func GetOpponentRecruitingList(c *gin.Context) ([]*models.OpponentRecruiting, *models.Page) {
	pageNumber, _ := strconv.Atoi(c.Param("page"))
	db := database.GetDB()
	totalElements := int(db.Find(&opponentRecruitings).RowsAffected)
	if pageSize > totalElements {
		pageSize = totalElements
	}

	page := &models.Page{
		Number:        pageNumber,
		Size:          pageSize,
		TotalElements: totalElements,
		TotalPages:    int(math.Ceil(float64(totalElements) / float64(pageSize))),
	}

	sort := &models.Sort{
		IsDesc:  true,
		OrderBy: "created_at",
	}

	db.Scopes(page.Paginate()).Scopes(sort.Sort()).Find(&opponentRecruitings)
	return opponentRecruitings, page
}
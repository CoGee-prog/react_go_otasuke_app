package services

import (
	"math"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OpponentRecruitingService struct {
	db *database.GormDatabase
}

// 対戦相手募集のサービスを作成する
func NewOpponentRecruitingService(db *database.GormDatabase) *OpponentRecruitingService {
	return &OpponentRecruitingService{
		db: db,
	}
}

// 対戦相手募集の構造体の配列
var opponentRecruitings []*models.OpponentRecruiting

// リスト表示時の1ページあたりの要素数
var pageSize int = 10

// 対戦相手募集のリストとページ情報を返す
func (ors *OpponentRecruitingService) GetOpponentRecruitingList(c *gin.Context) ([]*models.OpponentRecruiting, *models.Page) {
	pageNumber, _ := strconv.Atoi(c.Param("page"))

	// 合計要素数
	totalElements := int(ors.db.DB.Find(&opponentRecruitings).RowsAffected)
	// ページサイズが合計要素数を超えている場合
	if pageSize > totalElements {
		// 合計要素数が0より大きければ合計要素数に合わせる
		if totalElements > 0 {
			pageSize = totalElements
		} else {
			// それより小さければ1にする
			pageSize = 1
		}
	}

	// 合計ページ数
	totalPages := int(math.Ceil(float64(totalElements) / float64(pageSize)))

	// 指定されたページ数が合計ページ数を超えていたら合計ページ数に合わせる
	if pageNumber > totalPages {
		pageNumber = totalPages
	}

	page := &models.Page{
		Number:        pageNumber,
		Size:          pageSize,
		TotalElements: totalElements,
		TotalPages:    totalPages,
	}

	sort := &models.Sort{
		IsDesc:  true,
		OrderBy: "created_at",
	}

	ors.db.DB.Scopes(page.Paginate()).Scopes(sort.Sort()).Find(&opponentRecruitings)
	return opponentRecruitings, page
}

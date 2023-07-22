package models

import (
	"errors"
	"math"
	"react_go_otasuke_app/database"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type OpponentRecruiting struct {
	gorm.Model
	TeamId   int       `json:"team_id" gorm:"type:int"`
	DateTime time.Time `json:"date_time"`
	AreaId   int       `json:"area" gorm:"type:int; not null"`
	Detail   *string   `json:"detail" gorm:"type:text"`
}

// 対戦相手募集の構造体の配列
var opponentRecruitings []*OpponentRecruiting

// リスト表示時の1ページあたりの要素数
var pageSize int = 10

// 対戦相手募集のバリデーション
func (oc *OpponentRecruiting) Validate() error {
	if oc.TeamId == 0 {
		return errors.New("チームが選択されていません")
	}
	if oc.AreaId == 0 {
		return errors.New("エリアが選択されていません")
	}
	if oc.DateTime.Before(time.Now()) {
		return errors.New("過去の日時は選択できません")
	}
	return nil
}

// 対戦相手募集のデータを作成する
func (or *OpponentRecruiting) Create() (err error) {
	db := database.GetDB()
	return db.Create(or).Error
}

// 対戦相手募集のリストとページ情報を返す
func (or *OpponentRecruiting) GetOpponentRecruitingList(c *gin.Context) ([]*OpponentRecruiting, *Page) {
	pageNumber, _ := strconv.Atoi(c.Param("page"))
	db := database.GetDB()
	totalElements := int(db.Find(&opponentRecruitings).RowsAffected)
	if pageSize > totalElements {
		pageSize = totalElements
	}

	page := &Page{
		Number:        pageNumber,
		Size:          pageSize,
		TotalElements: totalElements,
		TotalPages:    int(math.Ceil(float64(totalElements) / float64(pageSize))),
	}

	sort := &Sort{
		IsDesc:  true,
		OrderBy: "created_at",
	}

	db.Scopes(page.Paginate()).Scopes(sort.Sort()).Find(&opponentRecruitings)
	return opponentRecruitings, page
}

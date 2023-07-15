package models

import (
	"errors"
	"fmt"
	"react_go_otasuke_app/database"
	"time"

	"github.com/jinzhu/gorm"
)

type OpponentRecruiting struct {
	gorm.Model
	TeamId   int       `json:"team_id" gorm:"type:int"`
	DateTime time.Time `json:"date_time"`
	AreaId   int       `json:"area" gorm:"type:int; not null"`
	Detail   *string   `json:"detail" gorm:"type:text"`
}

var opponentRecruitings []OpponentRecruiting

func (oc *OpponentRecruiting) Validate() error {
	if oc.TeamId == 0 {
		return errors.New("チームが選択されていません")
	}
	if oc.AreaId == 0 {
		return errors.New("エリアが選択されていません")
	}
	fmt.Print(time.Now())
	if oc.DateTime.Before(time.Now()) {
		return errors.New("過去の日時は選択できません")
	}
	return nil
}

func (or *OpponentRecruiting) Create() (err error) {
	db := database.GetDB()
	return db.Create(or).Error
}

func (or *OpponentRecruiting) GetByPagination(page *Page) *gorm.DB {
	db := database.GetDB()
	sort := &Sort{
		IsDesc:  true,
		OrderBy: "created_at",
	}
	return db.Scopes(page.Paginate()).Scopes(sort.Sort()).Find(&opponentRecruitings)
}

package models

import (
	"errors"
	"react_go_otasuke_app/database"
	"time"

	"github.com/jinzhu/gorm"
)

type OpponentRecruiting struct {
	BaseModel
	TeamId int        `json:"team_id"`
	Date   *time.Time `json:"date" gorm:"type:date"`
	Time   *time.Time `json:"time" gorm:"type:time"`
	AreaId int        `json:"area" gorm:"not null"`
	Detail *string    `json:"detail" gorm:"type:text"`
}

var opponentRecruitings []OpponentRecruiting

func (opponentRecruiting *OpponentRecruiting) Validate() error {
	if opponentRecruiting.TeamId == 0 {
		return errors.New("チームが選択されていません")
	}
	if opponentRecruiting.AreaId == 0 {
		return errors.New("エリアが選択されていません")
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

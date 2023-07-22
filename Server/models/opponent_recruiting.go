package models

import (
	"errors"
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

// 対戦相手募集のバリデーション
func (or *OpponentRecruiting) Validate() error {
	if or.TeamId == 0 {
		return errors.New("チームが選択されていません")
	}
	if or.AreaId == 0 {
		return errors.New("エリアが選択されていません")
	}
	if or.DateTime.Before(time.Now()) {
		return errors.New("過去の日時は選択できません")
	}
	return nil
}

// 対戦相手募集のデータを作成する
func (or *OpponentRecruiting) Create() (err error) {
	db := database.GetDB()
	return db.Create(or).Error
}

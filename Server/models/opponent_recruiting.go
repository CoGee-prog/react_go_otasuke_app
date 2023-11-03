package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type OpponentRecruiting struct {
	gorm.Model
	TeamId   int       `json:"team_id" gorm:"type:int"`
	AreaId   int       `json:"area_id" gorm:"type:int; not null"`
	DateTime time.Time `json:"date_time"`
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

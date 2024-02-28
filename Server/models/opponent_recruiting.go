package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type OpponentRecruiting struct {
	gorm.Model
	TeamId       uint       `json:"team_id" gorm:"type:int;not null"`
	Team         Team       `gorm:"foreignKey:TeamId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PrefectureId Prefecture `json:"prefecture_id" gorm:"type:int; not null"`
	StartTime    time.Time  `json:"start_time" gorm:"not null"`
	EndTime      time.Time  `json:"end_time" gorm:"not null"`
	Detail       *string    `json:"detail" gorm:"type:text"`
}

// 対戦相手募集のバリデーション
func (or *OpponentRecruiting) Validate() error {
	if or.TeamId == 0 {
		return errors.New("チームが選択されていません")
	}
	if or.PrefectureId < Hokkaido || or.PrefectureId > Okinawa {
		return errors.New("不正な都道府県です")
	}
	if or.StartTime.Before(time.Now()) {
		return errors.New("過去の日時は選択できません")
	}
	if or.EndTime.Before(or.StartTime) {
		return errors.New("開始より過去の日時は選択できません")
	}
	return nil
}

package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type OpponentRecruiting struct {
	gorm.Model
	TeamId       uint         `json:"team_id" gorm:"type:int;not null"`
	Team         Team         `gorm:"foreignKey:TeamId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title        string       `json:"title" gorm:"type:text; not null"`
	HasGround    bool         `json:"has_ground" gorm:"not null;default:false"`
	GroundName   string       `json:"ground_name" gorm:"type:text"`
	PrefectureId PrefectureId `json:"prefecture_id" gorm:"type:int; not null"`
	StartTime    time.Time    `json:"start_time" gorm:"not null"`
	EndTime      time.Time    `json:"end_time" gorm:"not null"`
	Detail       string       `json:"detail" gorm:"type:text"`
}

// 対戦相手募集のバリデーション
func (or *OpponentRecruiting) Validate() error {
	if or.TeamId == 0 {
		return errors.New("チームが選択されていません")
	}
	if or.Title == "" { 
		return errors.New("タイトルは必須です")
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
	// 開始時間と終了時間が同じ日であるか確認
	if or.StartTime.Year() != or.EndTime.Year() || or.StartTime.Month() != or.EndTime.Month() || or.StartTime.Day() != or.EndTime.Day() {
		return errors.New("終了時間は開始日と同じ日にしてください")
	}
	// グラウンド有の場合、グラウンド名があるか確認
	if or.HasGround && or.GroundName == "" {
		return errors.New("グラウンド有の場合、グラウンド名は必須です")
	}
	// グラウンド無の場合、グラウンド名がないか確認
	if !or.HasGround && or.GroundName != "" {
		return errors.New("グラウンド無の場合、グラウンド名は不要です")
	}
	return nil
}

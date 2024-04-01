package models

import (
	"errors"
	"time"
	"unicode/utf8"

	"gorm.io/gorm"
)

type OpponentRecruiting struct {
	gorm.Model
	TeamID       uint                         `json:"team_id" gorm:"type:int; not null"`
	Team         Team                         `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title        string                       `json:"title" gorm:"type:varchar(255); not null"`
	HasGround    bool                         `json:"has_ground" gorm:"not null; default:false"`
	GroundName   string                       `json:"ground_name" gorm:"type:varchar(255)"`
	PrefectureID PrefectureID                 `json:"prefecture_id" gorm:"type:int; not null"`
	StartTime    time.Time                    `json:"start_time" gorm:"not null"`
	EndTime      time.Time                    `json:"end_time" gorm:"not null"`
	Detail       string                       `json:"detail" gorm:"type:text"`
	IsActive     bool                         `json:"is_active" gorm:"not null; default:true"`
	Comments     []*OpponentRecruitingComment `json:"comments" gorm:"foreignKey:OpponentRecruitingID"`
}

// 対戦相手募集のバリデーション
func (or *OpponentRecruiting) Validate() error {
	if or.TeamID == 0 {
		return errors.New("チームが選択されていません")
	}
	if or.Title == "" {
		return errors.New("タイトルは必須です")
	}
	if utf8.RuneCountInString(or.Title) > 50 {
		return errors.New("タイトルは50文字以下でなければなりません")
	}
	if or.PrefectureID < Hokkaido || or.PrefectureID > Okinawa {
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
	if utf8.RuneCountInString(or.GroundName) > 50 {
		return errors.New("グラウンド名は50文字以下でなければなりません")
	}
	if utf8.RuneCountInString(or.Detail) > 1000 {
		return errors.New("詳細は1000文字以下でなければなりません")
	}
	return nil
}

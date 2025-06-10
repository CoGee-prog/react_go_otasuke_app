package models

import (
	"errors"
	"time"
	"unicode/utf8"
)

type TeamLevelId int

const (
	// 地区レベル
	LocalLevel TeamLevelId = iota + 1
	// 地区上位レベル
	LocalEliteLevel
	// 都道府県レベル
	PrefecturalLevel
	// 都道府県上位レベル
	PrefecturalEliteLevel
	// 全国レベル
	NationalLevel
)

func (TL TeamLevelId) ToString() string {
	switch TL {
	case LocalLevel:
		return "地区レベル"
	case LocalEliteLevel:
		return "地区上位レベル"
	case PrefecturalLevel:
		return "都道府県レベル"
	case PrefecturalEliteLevel:
		return "都道府県上位レベル"
	case NationalLevel:
		return "全国レベル"
	default:
		return "未定義"
	}
}

type Team struct {
	ID                  string `json:"id" gorm:"type:varchar(64);primary_key"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time            `sql:"index"`
	Name                string                `json:"name" gorm:"type:text;not null"`
	PrefectureId        PrefectureID          `json:"prefecture_id" gorm:"type:int;not null"`
	LevelId             TeamLevelId           `json:"level_id" gorm:"type:int;not null"`
	HomePageUrl         *string               `json:"home_page_url" gorm:"type:text"`
	Other               *string               `json:"other" gorm:"type:text"`
	Users               []*User               `gorm:"many2many:user_teams"`
	OpponentRecruitings []*OpponentRecruiting `json:"opponent_recruitings" gorm:"foreignKey:TeamID;OnDelete:CASCADE;"`
}

// チームのバリデーション
func (t *Team) Validate() error {
	if t.Name == "" {
		return errors.New("チーム名が入力されていません")
	}
	if utf8.RuneCountInString(t.Name) > 32 {
		return errors.New("チーム名は32文字以下でなければなりません")
	}
	if t.PrefectureId < Hokkaido || t.PrefectureId > Okinawa {
		return errors.New("不正な活動拠点です")
	}
	if t.LevelId < LocalLevel || t.LevelId > NationalLevel {
		return errors.New("不正なチームレベルです")
	}
	if t.HomePageUrl != nil && utf8.RuneCountInString(*t.HomePageUrl) > 500 {
		return errors.New("ホームページリンクは500文字以下でなければなりません")
	}
	if t.Other != nil && utf8.RuneCountInString(*t.Other) > 500 {
		return errors.New("その他は500文字以下でなければなりません")
	}
	return nil
}

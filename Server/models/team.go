package models

import (
	"errors"

	"gorm.io/gorm"
)

type TeamLevel int

const (
	// 地区レベル
	LocalLevel TeamLevel = iota + 1
	// 地区上位レベル
	LocalEliteLevel
	// 都道府県レベル
	PrefecturalLevel
	// 都道府県上位レベル
	PrefecturalEliteLevel
	// 全国レベル
	NationalLevel
)

func (TL TeamLevel) ToString() string {
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
	gorm.Model
	Name        string    `json:"name" gorm:"type:text;not null"`
	LevelId     TeamLevel `json:"level_id" gorm:"type:integer;not null"`
	HomePageUrl *string   `json:"home_page_url" gorm:"type:text"`
	Other       *string   `json:"other" gorm:"type:text"`
	Users       []*User   `gorm:"many2many:user_teams"`
}

// チームのバリデーション
func (t *Team) Validate() error {
	if t.Name == "" {
		return errors.New("チーム名が入力されていません")
	}
	if t.LevelId < LocalLevel || t.LevelId > NationalLevel {
		return errors.New("不正なチームレベルです")
	}
	return nil
}

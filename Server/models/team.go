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
	// 全国レベル（ナショナル）
	NationalLevel
)

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

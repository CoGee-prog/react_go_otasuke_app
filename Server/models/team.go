package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:text"`
	LevelId     uint    `json:"level_id" gorm:"type:integer"`
	HomePageUrl string  `json:"home_page_url" gorm:"type:text"`
	Other       string  `json:"other" gorm:"type:text"`
	Users       []*User `gorm:"many2many:user_teams;"`
}

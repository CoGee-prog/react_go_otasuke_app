package views

import (
	"react_go_otasuke_app/models"
)

type TeamView struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" gorm:"type:text;not null"`
	Level       string  `json:"level" gorm:"type:integer;not null"`
	HomePageUrl *string `json:"home_page_url" gorm:"type:text"`
	Other       *string `json:"other" gorm:"type:text"`
}

// Teamの構造体から必要なキーのみ返す
func CreateTeamView(team models.Team) *TeamView {
	return &TeamView{
		ID:          team.ID,
		Name:        team.Name,
		Level:       team.LevelId.ToString(),
		HomePageUrl: team.HomePageUrl,
		Other:       team.Other,
	}
}

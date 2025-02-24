package views

import (
	"react_go_otasuke_app/models"
)

type TeamView struct {
	ID           string              `json:"id"`
	Name         string              `json:"name" gorm:"type:text;not null"`
	PrefectureID models.PrefectureID `json:"prefecture_id" gorm:"type:integer;not null"`
	LevelID      models.TeamLevelId  `json:"level_id" gorm:"type:integer;not null"`
	HomePageUrl  *string             `json:"home_page_url" gorm:"type:text"`
	Other        *string             `json:"other" gorm:"type:text"`
}

// Teamの構造体から必要なキーのみ返す
func CreateTeamView(team models.Team) *TeamView {
	return &TeamView{
		ID:           team.ID,
		Name:         team.Name,
		PrefectureID: team.PrefectureId,
		LevelID:      team.LevelId,
		HomePageUrl:  team.HomePageUrl,
		Other:        team.Other,
	}
}

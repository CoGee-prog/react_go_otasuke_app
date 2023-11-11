package models

import "time"

type User struct {
	Id           string `json:"id" gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	Name          string     `json:"name" gorm:"type:text"`
	CurrentTeamId *int       `json:"current_team_id" gorm:"type:int"`
}

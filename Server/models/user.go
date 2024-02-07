package models

import "time"

type User struct {
	Id            string `json:"id" gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	Name          string     `gorm:"type:text"`
	CurrentTeamId *uint      `gorm:"type:int"`
	CurrentTeam   *Team      `gorm:"foreignKey:CurrentTeamId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Teams         []*Team    `gorm:"many2many:user_teams;"`
}

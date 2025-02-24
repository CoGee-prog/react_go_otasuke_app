package models

import "time"

type User struct {
	ID              string `json:"id" gorm:"type:varchar(64);primary_key"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
	Name            string     `gorm:"type:text"`
	CurrentTeamId   *string    `gorm:"type:varchar(64)"`
	CurrentTeam     *Team      `gorm:"foreignKey:CurrentTeamId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CurrentTeamRole *TeamRole  `gorm:"-"`
	Teams           []*Team    `gorm:"many2many:user_teams;"`
}

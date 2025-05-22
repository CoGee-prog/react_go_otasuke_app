package models

import "time"

type TeamInvite struct {
	TeamID    string `json:"team_id" gorm:"not null"`
	Team      Team   `gorm:"foreignKey:TeamID;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Token     string     `gorm:"type:varchar(64);not null;unique"`
	ExpiresAt time.Time
}

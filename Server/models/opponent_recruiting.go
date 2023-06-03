package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OpponentRecruiting struct {
	gorm.Model
	TeamId         int       `json:"team_id"`
	TeamName       string    `json:"team_name" gorm:"not null"`
	Date           time.Time `json:"date" gorm:"type:date"`
	Time           time.Time `json:"time" gorm:"type:time"`
	AreaId         int       `json:"area" gorm:"not null"`
	Detail         string    `json:"detail" gorm:"type:text"`
}

// func (or *OpponentRecruiting) Create() (err error) {
	
// }

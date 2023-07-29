package database

import (
	"github.com/jinzhu/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func NewGormDatabase(db *gorm.DB) *GormDatabase {
	return &GormDatabase{
		DB: db,
	}
}

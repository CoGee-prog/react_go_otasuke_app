package database

import (
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func NewGormDatabase(db *gorm.DB) *GormDatabase {
	return &GormDatabase{
		DB: db,
	}
}

package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Sort struct {
	IsDesc  bool
	OrderBy string
}

func (sort *Sort) Sort() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		direction := "asc"
		if sort.IsDesc {
			direction = "desc"
		}

		order := strings.Join([]string{sort.OrderBy, direction}, " ")

		return db.Order(order)
	}
}

type Page struct {
	Number int
	Size   int
}

func (page *Page) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.Number <= 0 {
			page.Number = 1
		}

		switch {
		case page.Size > 100:
			page.Size = 100
		case page.Size <= 0:
			page.Size = 10
		}

		offset := (page.Number - 1) * page.Size
		return db.Offset(offset).Limit(page.Size)
	}
}

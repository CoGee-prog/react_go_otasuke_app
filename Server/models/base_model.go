package models

import (
	"strings"

	"github.com/jinzhu/gorm"
)

type Sort struct {
	IsDesc  bool
	OrderBy string
}

// ソートした結果を返す
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
	Number        int `json:"number"`
	Size          int `json:"size"`
	TotalElements int `json:"total_elements"`
	TotalPages    int `json:"total_pages"`
}

// ページネーションした結果を返す
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

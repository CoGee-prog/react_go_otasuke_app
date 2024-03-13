package database

import (
	"strings"

	"gorm.io/gorm"
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

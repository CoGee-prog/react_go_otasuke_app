package database

import (
	"react_go_otasuke_app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var d *gorm.DB

// DBに接続する
func Init() {
	c := config.GetConfig()
	var err error
	d, err = gorm.Open(c.GetString("db.provider"), c.GetString("db.url"))
	if err != nil {
		panic(err)
	}
}

// 自動マイグレーションを行う
func Migration(models ...interface{}) {
	d.AutoMigrate(models...)
}

// DBを取得する
func GetDB() *gorm.DB {
	if d == nil {
		Init()
	}

	return d
}

// DBの接続を切る
func Close() {
	d.Close()
}

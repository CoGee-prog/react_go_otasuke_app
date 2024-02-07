package database

import (
	"react_go_otasuke_app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var d *gorm.DB

// DBに接続する
func Init() {
	c := config.GetConfig()
	var err error
	d, err = gorm.Open(mysql.Open(c.GetString("db.url")), &gorm.Config{})
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
	db, err := d.DB()
	if err != nil{
		panic(err)
	}
	db.Close()
}

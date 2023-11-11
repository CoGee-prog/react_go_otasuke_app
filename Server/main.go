package main

import (
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/server"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {

}

func main() {
	// 設定を読み込み
	// config.Init(os.Getenv("APP_ENV"))

	// データベースの設定
	database.Init()
	models := []interface{}{
		&models.OpponentRecruiting{},
		&models.User{},
	}
	database.Migration(models)
	defer database.Close()

	// サーバー起動
	if err := server.Init(); err != nil {
		panic(err)
	}
}

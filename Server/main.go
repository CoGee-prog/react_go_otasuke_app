package main

import (
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/server"

	_ "gorm.io/driver/mysql"
)

func init() {

}

func main() {
	// 設定を読み込み
	config.Init()

	// データベースの設定
	database.Init()
	models := []interface{}{
		&models.OpponentRecruiting{},
		&models.Team{},
		&models.User{},
		&models.UserTeam{},
	}
	database.Migration(models...)
	defer database.Close()

	// サーバー起動
	if err := server.Init(); err != nil {
		panic(err)
	}
}

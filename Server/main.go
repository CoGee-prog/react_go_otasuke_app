package main

import (
	"flag"
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/server"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 設定を読み込み
	env := flag.String("e", "development", "")
	flag.Parse()
	config.Init(*env)

	// データベースの設定
	database.Init()
	database.Migration(&models.OpponentRecruiting{})
	defer database.Close()
	
	// サーバー起動
	if err := server.Init(); err != nil {
		panic(err)
	}
}

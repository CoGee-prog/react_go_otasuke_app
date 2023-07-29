package controllers

import "react_go_otasuke_app/database"

type BaseController struct {
	db *database.GormDatabase
}

// 基盤のコントローラーを作成する
func NewBaseController(db *database.GormDatabase) *BaseController {
	return &BaseController{
		db: db,
	}
}

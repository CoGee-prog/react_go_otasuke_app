package models

import "time"

type TeamRole int

const (
	// 0: 通常のチームメンバー
	TeamMember TeamRole = iota
	// 1: チーム管理者
	TeamAdmin
	// 2: チーム副管理者
	TeamSubAdmin
)

type UserTeam struct {
	UserID    string `gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeamID    string `gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Role      TeamRole   `gorm:"type:int;not null"`
	User      User       `gorm:"foreignKey:UserID;references:ID;"`
	Team      Team       `gorm:"foreignKey:TeamID;references:ID;"`
}

// ユーザーがチームの管理者または副管理者であるかどうかを確認する
func (ut *UserTeam) IsAdminOrSubAdmin() bool {
	return ut.Role == TeamAdmin || ut.Role == TeamSubAdmin
}

package models

type TeamRole int

const (
	// 0: 通常のチームメンバー
	TeamMember TeamRole = iota
	// 1: チーム管理者
	TeamAdmin
	// 2: チーム副管理者
	TeamSubAdmin
)

type UserTeams struct {
	UserID string   `gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeamID uint     `gorm:"primaryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role   TeamRole `gorm:"type:int;not null"`
	User   User     `gorm:"foreignKey:UserID;references:ID;"`
	Team   Team     `gorm:"foreignKey:TeamID;references:ID;"`
}

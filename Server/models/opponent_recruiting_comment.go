package models

import (
	"errors"

	"gorm.io/gorm"
)

type OpponentRecruitingComment struct {
	gorm.Model
	OpponentRecruitingID uint               `json:"opponent_recruiting_id" gorm:"index"`
	OpponentRecruiting   OpponentRecruiting `gorm:"foreignKey:OpponentRecruitingID;constraint:OnDelete:CASCADE;"`
	UserID               *string            `json:"user_id" gorm:"index"`
	User                 *User              `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;"`
	TeamID               *uint              `json:"team_id"`
	Team                 *Team              `gorm:"foreignKey:TeamID;constraint:OnDelete:SET NULL;"`
	Content              string             `json:"content"`
}

// 対戦相手募集のコメント作成時のバリデーション
func (orc *OpponentRecruitingComment) ValidateCreate() error {
	// コメント内容が空の場合はエラー
	if orc.Content == "" {
		return errors.New("コメントが空です")
	}

	// コメント内容の文字数制限
	if len(orc.Content) > 1000 {
		return errors.New("コメントは1000文字以内でなければなりません")
	}

	// OpponentRecruitingIDが0の場合はエラー
	if orc.OpponentRecruitingID == 0 {
		return errors.New("対戦相手募集IDは必須です")
	}

	// UserIDがnilの場合はエラー
	if orc.UserID == nil {
		return errors.New("ユーザーIDは必須です")
	}

	return nil
}

func (orc *OpponentRecruitingComment) ValidateUpdate() error {
	// コメント内容が空の場合はエラー
	if orc.Content == "" {
		return errors.New("コメントが空です")
	}

	// コメント内容の文字数制限
	if len(orc.Content) > 1000 {
		return errors.New("コメントは1000文字以内でなければなりません")
	}

	return nil
}
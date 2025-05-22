package repositories

import (
	"errors"
	"react_go_otasuke_app/models"
	"time"

	"gorm.io/gorm"
)

type TeamInviteRepository interface {
	GetByToken(tx *gorm.DB, token string) (*models.TeamInvite, error)
	CreateOrUpdateInvite(tx *gorm.DB, teamId string, token string) (*models.TeamInvite, error)
}

type teamInviteRepository struct{}

func NewTeamInviteRepository() TeamInviteRepository {
	return &teamInviteRepository{}
}

// トークンでチーム招待を取得する
func (r *teamInviteRepository) GetByToken(tx *gorm.DB, token string) (*models.TeamInvite, error) {
	var teamInvite models.TeamInvite
	result := tx.Where("token = ?", token).First(&teamInvite)
	// レコードが見つからない場合はnilを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
		// その他のエラーの場合
	} else if result.Error != nil {
		return nil, result.Error
	}
	// レコードが見つかった場合
	return &teamInvite, nil
}

// IDでチームを取得する(なければエラー)
func (r *teamInviteRepository) CreateOrUpdateInvite(tx *gorm.DB, teamId string, token string) (*models.TeamInvite, error) {
	var expiresAt = time.Now().Add(time.Hour * 24)
	err := tx.Exec(`
    INSERT INTO team_invites (team_id, token, expires_at)
    VALUES (?, ?, ?)
    ON DUPLICATE KEY UPDATE
        token = VALUES(token), expires_at = VALUES(expires_at)
		`, teamId, token, expiresAt).Error

	if err != nil {
		return nil, errors.New("チーム招待の作成に失敗しました")
	}

	teamInvite := models.TeamInvite{
		TeamID: teamId,
		Token: token,
		ExpiresAt: expiresAt,
	}

	return &teamInvite, nil
}
package repositories

import (
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type UserTeamRepository interface {
	AddTeamAdmin(tx *gorm.DB, userId string, team *models.Team) error
	GetUserTeam(tx *gorm.DB, userId string, teamId uint) (*models.UserTeam, error)
}

type userTeamRepository struct{}

func NewUserTeamRepository() UserTeamRepository {
	return &userTeamRepository{}
}


// チーム管理者としてユーザーを追加する
func (r *userTeamRepository) AddTeamAdmin(tx *gorm.DB, userId string, team *models.Team) error {
	// チーム作成者をチーム管理者としてチームに所属させる
	userTeam := models.UserTeam{
		UserID: userId,
		TeamID: team.ID,
		Role:   models.TeamAdmin,
	}
	if err := tx.Create(&userTeam).Error; err != nil {
		return err
	}

	return nil
}

// 指定したユーザーとチームの中間テーブルのレコードを取得する
func (r *userTeamRepository) GetUserTeam(tx *gorm.DB, userId string, teamId uint) (*models.UserTeam, error) { 
	var userTeam models.UserTeam 
	if err := tx.Where("user_id = ? AND team_id = ?", userId, teamId).First(&userTeam).Error; err != nil {
		return nil, err
	}
	return &userTeam, nil
}
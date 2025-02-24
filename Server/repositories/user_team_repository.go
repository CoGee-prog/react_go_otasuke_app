package repositories

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type UserTeamRepository interface {
	AddTeamAdmin(tx *gorm.DB, userId string, team *models.Team) error
	GetByUserIdAndTeamId(tx *gorm.DB, userId string, teamId string) (*models.UserTeam, error)
	FindByUserIdAndTeamId(tx *gorm.DB, userId string, teamId string) (*models.UserTeam, error)
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
func (r *userTeamRepository) GetByUserIdAndTeamId(tx *gorm.DB, userId string, teamId string) (*models.UserTeam, error) {
	var userTeam models.UserTeam
	result := tx.Where("user_id = ? AND team_id = ?", userId, teamId).First(&userTeam)
	// レコードが見つからない場合はnilを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
		// その他のエラーの場合
	} else if result.Error != nil {
		return nil, result.Error
	}
	// レコードが見つかった場合
	return &userTeam, nil
}

// 指定したユーザーとチームの中間テーブルのレコードを取得する(なければエラー)
func (r *userTeamRepository) FindByUserIdAndTeamId(tx *gorm.DB, userId string, teamId string) (*models.UserTeam, error) {
	var userTeam models.UserTeam
	if err := tx.Where("user_id = ? AND team_id = ?", userId, teamId).First(&userTeam).Error; err != nil {
		return nil, err
	}
	return &userTeam, nil
}

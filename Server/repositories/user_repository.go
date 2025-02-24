package repositories

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUserId(tx *gorm.DB, userId string) (*models.User, error)
	GetWithCurrentTeamByUserId(tx *gorm.DB, userId string) (*models.User, error)
	Create(tx *gorm.DB, user *models.User) error
	ChangeCurrentTeam(tx *gorm.DB, userId string, teamId string) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// ユーザーを取得する
func (r *userRepository) GetByUserId(tx *gorm.DB, userId string) (*models.User, error) {
	var user models.User

	result := tx.Where("id = ?", userId).First(&user)
	// レコードが見つからない場合はnilを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
		// その他のエラーの場合
	} else if result.Error != nil {
		return nil, result.Error
	}
	// レコードが見つかった場合
	return &user, nil
}

// ユーザーと現在のチームを取得する
func (r *userRepository) GetWithCurrentTeamByUserId(tx *gorm.DB, userId string) (*models.User, error) {
	var user models.User

	result := tx.Preload("CurrentTeam").Where("id = ?", userId).First(&user)
	// レコードが見つからない場合はnilを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
		// その他のエラーの場合
	} else if result.Error != nil {
		return nil, result.Error
	}
	// レコードが見つかった場合
	return &user, nil
}

// ユーザーを作成
func (r *userRepository) Create(tx *gorm.DB, user *models.User) error {
	return tx.Create(user).Error
}

// ユーザーの現在のチームを切り替える
func (r *userRepository) ChangeCurrentTeam(tx *gorm.DB, userId string, teamId string) error {
	return tx.Model(&models.User{}).Where("id = ?", userId).Update("current_team_id", teamId).Error
}

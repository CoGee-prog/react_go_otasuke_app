package repositories

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(tx *gorm.DB, userId string) (*models.User, error)
	CreateUser(tx *gorm.DB, user *models.User) error
	ChangeUserCurrentTeam(tx *gorm.DB, userId string, teamId uint) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// ユーザーを取得する
func (r *userRepository) GetUser(tx *gorm.DB, userId string) (*models.User, error) {
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
func (r *userRepository) CreateUser(tx *gorm.DB, user *models.User) error {
	return tx.Create(user).Error
}

// ユーザーの現在のチームを切り替える
func (r *userRepository) ChangeUserCurrentTeam(tx *gorm.DB, userId string, teamId uint) error {
	return tx.Model(&models.User{}).Where("id = ?", userId).Update("current_team_id", teamId).Error
}

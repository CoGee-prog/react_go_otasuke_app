package repositories

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

// ユーザーを取得する
func (r *UserRepository) Get(userId string)(*models.User, error) {
	var user models.User

	result := r.db.Preload("CurrentTeam").Where("id = ?", userId).First(&user)
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
func (r *UserRepository) Create(tx *gorm.DB, user *models.User) error {
  return tx.Create(user).Error
}

// ユーザーの現在のチームを切り替える
func (r *UserRepository) ChangeUserCurrentTeam(tx *gorm.DB, userId string, teamId uint) error {
	return tx.Model(&models.User{}).Where("id = ?", userId).Update("current_team_id", teamId).Error
}
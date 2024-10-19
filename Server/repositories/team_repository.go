package repositories

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type TeamRepository interface {
	GetTeam(tx *gorm.DB, userId string) (*models.Team, error)
	CreateTeam(tx *gorm.DB, team *models.Team) error
}

type teamRepository struct{}

func NewTeamRepository() TeamRepository {
	return &teamRepository{}
}

// チームを取得する
func (r *teamRepository) GetTeam(tx *gorm.DB, id string) (*models.Team, error) {
	var team models.Team
	result := tx.Where("id = ?", id).First(&team)
	// レコードが見つからない場合はnilを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
		// その他のエラーの場合
	} else if result.Error != nil {
		return nil, result.Error
	}
	// レコードが見つかった場合
	return &team, nil
}

// チームを作成する
func (r *teamRepository) CreateTeam(tx *gorm.DB, team *models.Team) error {
	// チームを作成する
	if err := tx.Create(team).Error; err != nil {
		return err
	}

	return nil
}

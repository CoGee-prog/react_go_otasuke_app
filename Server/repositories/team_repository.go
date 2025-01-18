package repositories

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type TeamRepository interface {
	GetByTeamId(tx *gorm.DB, teamId uint) (*models.Team, error)
	Create(tx *gorm.DB, team *models.Team) error
}

type teamRepository struct{}

func NewTeamRepository() TeamRepository {
	return &teamRepository{}
}

// チームを取得する
func (r *teamRepository) GetByTeamId(tx *gorm.DB, teamId uint) (*models.Team, error) {
	var team models.Team
	result := tx.Where("id = ?", teamId).First(&team)
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
func (r *teamRepository) Create(tx *gorm.DB, team *models.Team) error {
	// チームを作成する
	if err := tx.Create(team).Error; err != nil {
		return err
	}

	return nil
}

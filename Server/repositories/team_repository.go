package repositories

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type TeamRepository interface {
	GetById(tx *gorm.DB, teamId uint) (*models.Team, error)
	Create(tx *gorm.DB, team *models.Team) error
	UpdateById(tx *gorm.DB, id uint, team *models.Team) error
	FindById(tx *gorm.DB, id uint) (*models.Team, error)
}

type teamRepository struct{}

func NewTeamRepository() TeamRepository {
	return &teamRepository{}
}

// IDでチームを取得する
func (r *teamRepository) GetById(tx *gorm.DB, id uint) (*models.Team, error) {
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

// IDでチームを取得する(なければエラー)
func (r *teamRepository) FindById(tx *gorm.DB, id uint) (*models.Team, error) {
	var team models.Team
	result := tx.First(&team, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("データが見つかりません")
		}
		return nil, errors.New("データ取得に失敗しました")
	}
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

// チームを更新する
func (r *teamRepository) UpdateById(tx *gorm.DB, id uint, team *models.Team) error {
	result := tx.Model(&models.Team{}).Where("id = ?", id).Updates(team)

	if result.Error != nil {
		return result.Error
	}

	// 更新件数が0の場合
	if result.RowsAffected == 0 {
		return errors.New("更新対象のデータがありません")
	}

	return nil
}
package repositories

import (
	"errors"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type OpponentRecruitingRepository interface {
	FindByID(tx *gorm.DB, id uint) (*models.OpponentRecruiting, error)
	FindByIDWithComments(tx *gorm.DB, id uint) (*models.OpponentRecruiting, error)
	Create(tx *gorm.DB, opponentRecruiting *models.OpponentRecruiting) error
	GetTotalCount(tx *gorm.DB) (int64, error)
	GetListWithTeamByPaginate(tx *gorm.DB, page *database.Page, sort *database.Sort) ([]*models.OpponentRecruiting, error)
	UpdateByID(tx *gorm.DB, id uint, opponentRecruiting *models.OpponentRecruiting) error
	UpdateIsActive(tx *gorm.DB, id uint, opponentRecruiting *models.OpponentRecruiting) error
	DeleteByID(tx *gorm.DB, id uint) error
}

type opponentRecruitingRepository struct{}

func NewOpponentRecruitingRepository() OpponentRecruitingRepository {
	return &opponentRecruitingRepository{}
}

// IDで対戦相手募集を取得
func (r *opponentRecruitingRepository) FindByID(tx *gorm.DB, id uint) (*models.OpponentRecruiting, error) {
	var opponentRecruiting models.OpponentRecruiting
	result := tx.First(&opponentRecruiting, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("データが見つかりません")
		}
		return nil, errors.New("データ取得に失敗しました")
	}
	return &opponentRecruiting, nil
}

// 対戦相手募集とコメントを取得
func (r *opponentRecruitingRepository) FindByIDWithComments(tx *gorm.DB, id uint) (*models.OpponentRecruiting, error) {
	var opponentRecruiting models.OpponentRecruiting

	result := tx.Preload("Team").Preload("Comments.User").Preload("Comments.Team").First(&opponentRecruiting, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("データが見つかりません")
		}
		return nil, errors.New("データ取得に失敗しました")
	}

	return &opponentRecruiting, nil
}

// 対戦相手募集を作成する
func (r *opponentRecruitingRepository) Create(tx *gorm.DB, opponentRecruiting *models.OpponentRecruiting) error {
	result := tx.Create(opponentRecruiting)
	if result.Error != nil {
		return errors.New("対戦相手募集の作成に失敗しました")
	}
	return nil
}

func (r *opponentRecruitingRepository) GetTotalCount(tx *gorm.DB) (int64, error) {
	var totalCount int64
	if err := tx.Model(&models.OpponentRecruiting{}).Count(&totalCount).Error; err != nil {
		return 0, err
	}
	return totalCount, nil
}

// ページネーション有りで対戦相手募集一覧とそのチーム情報を取得する
func (r *opponentRecruitingRepository) GetListWithTeamByPaginate(tx *gorm.DB, page *database.Page, sort *database.Sort) ([]*models.OpponentRecruiting, error) {
	var opponentRecruitings []*models.OpponentRecruiting

	result := tx.Scopes(page.Paginate(), sort.Sort()).Preload("Team").Find(&opponentRecruitings)

	if result.Error != nil {
		return nil, result.Error
	}

	return opponentRecruitings, nil
}

// IDを指定してOpponentRecruitingを更新する
func (r *opponentRecruitingRepository) UpdateByID(tx *gorm.DB, id uint, opponentRecruiting *models.OpponentRecruiting) error {
	result := tx.Model(&models.OpponentRecruiting{}).Where("id = ?", id).Updates(opponentRecruiting)

	if result.Error != nil {
		return result.Error
	}

	// 更新件数が0の場合
	if result.RowsAffected == 0 {
		return errors.New("更新対象のデータがありません")
	}

	return nil
}

// 対戦相手募集の募集中かどうかを更新する
func (r *opponentRecruitingRepository) UpdateIsActive(tx *gorm.DB, id uint, opponentRecruiting *models.OpponentRecruiting) error {
	result := tx.Model(&models.OpponentRecruiting{}).Where("id = ?", id).Select("IsActive").Updates(opponentRecruiting)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("更新対象のデータがありません")
	}

	return nil
}

// IDを指定して対戦相手募集を削除する
func (r *opponentRecruitingRepository) DeleteByID(tx *gorm.DB, id uint) error {
	result := tx.Unscoped().Delete(&models.OpponentRecruiting{}, "id = ?", id)
	if result.Error != nil {
		return errors.New("データ削除に失敗しました")
	}

	// 削除件数が0の場合
	if result.RowsAffected == 0 {
		return errors.New("削除対象のデータがありません")
	}

	return nil
}

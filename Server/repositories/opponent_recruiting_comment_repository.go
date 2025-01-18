package repositories

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type OpponentRecruitingCommentRepository interface {
	FindById(tx *gorm.DB, id uint) (*models.OpponentRecruitingComment, error)
	Create(tx *gorm.DB, opponentRecruitingComment *models.OpponentRecruitingComment) error
	UpdateById(tx *gorm.DB, id uint, opponentRecruitingComment *models.OpponentRecruitingComment) error
	SetDeleteFlagById(tx *gorm.DB, id uint) error
}

type opponentRecruitingCommentRepository struct{}

func NewOpponentRecruitingCommentRepository() OpponentRecruitingCommentRepository {
	return &opponentRecruitingCommentRepository{}
}

// 対戦相手募集のコメントを取得する(なければエラー)
func (r *opponentRecruitingCommentRepository) FindById(tx *gorm.DB, id uint) (*models.OpponentRecruitingComment, error) {
	var opponentRecruitingComment models.OpponentRecruitingComment
	if err := tx.First(&opponentRecruitingComment, id).Error; err != nil {
		return nil, err
	}
	return &opponentRecruitingComment, nil
}

// 対戦相手募集のコメントを作成する
func (r *opponentRecruitingCommentRepository) Create(tx *gorm.DB, opponentRecruitingComment *models.OpponentRecruitingComment) error {
	if err := tx.Create(opponentRecruitingComment).Error; err != nil {
		return err
	}
	return nil
}

// 対戦相手募集のコメントを更新する
func (r *opponentRecruitingCommentRepository) UpdateById(tx *gorm.DB, id uint, opponentRecruitingComment *models.OpponentRecruitingComment) error {
	result := tx.Model(&models.OpponentRecruitingComment{}).Where("id = ?", opponentRecruitingComment.ID).Updates(opponentRecruitingComment)
	if result.Error != nil {
		return errors.New("更新に失敗しました")
	}
	// 更新したデータが0件の場合はエラー
	if result.RowsAffected == 0 {
		return errors.New("更新対象のデータがありません")
	}
	return nil
}

// 削除済みフラグを立てる
func (r *opponentRecruitingCommentRepository) SetDeleteFlagById(tx *gorm.DB, id uint) error {
	// コメントの内容を更新し、削除済みフラグを立てる
	deleteOpponentRecruitingComment := &models.OpponentRecruitingComment{
		Content: "削除済みコメント",
		Deleted: true,
	}

	// データを更新する
	result := tx.Model(&models.OpponentRecruitingComment{}).Where("id = ?", id).Updates(deleteOpponentRecruitingComment)
	if result.Error != nil {
		return errors.New("削除に失敗しました")
	}
	// 更新したデータが0件の場合はエラー
	if result.RowsAffected == 0 {
		return errors.New("削除対象のデータがありません")
	}

	return nil
}

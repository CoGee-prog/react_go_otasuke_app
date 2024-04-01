package services

import (
	"errors"
	"math"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OpponentRecruitingService struct {
	userTeamService *UserTeamService
}

// 対戦相手募集のサービスを作成する
func NewOpponentRecruitingService(uts *UserTeamService) *OpponentRecruitingService {
	return &OpponentRecruitingService{
		userTeamService: uts,
	}
}

// 対戦相手募集を作成する
func (ors *OpponentRecruitingService) CreateOpponentRecruiting(db *gorm.DB, userId string, opponentRecruiting *models.OpponentRecruiting) error {
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(db, userId, opponentRecruiting.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集を作成できます")
	}
	// 対戦相手募集を作成する
	result := db.Create(opponentRecruiting)
	if result.Error != nil {
		return errors.New("対戦相手募集の作成に失敗しました")
	}
	return nil
}

// 対戦相手募集を変更する
func (ors *OpponentRecruitingService) UpdateOpponentRecruiting(db *gorm.DB, userId string, opponentRecruiting *models.OpponentRecruiting, id uint) error {
	// 変更する対戦相手募集を取得する
	originalOpponentRecruiting, err := ors.FindOpponentRecruiting(db, id)
	if err != nil {
		return err
	}
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(db, userId, originalOpponentRecruiting.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集を変更できます")
	}

	// データを更新する
	result := db.Model(&models.OpponentRecruiting{}).Where("id = ?", id).Updates(opponentRecruiting)
	if result.Error != nil {
		return result.Error
	}
	// 更新したデータが0件の場合はエラー
	if result.RowsAffected == 0 {
		return errors.New("更新対象のデータがありません")
	}
	return nil
}

// 対戦相手募集を取得する(なければエラー)
func (ors *OpponentRecruitingService) FindOpponentRecruiting(db *gorm.DB, id uint) (*models.OpponentRecruiting, error) {
	var opponentRecruiting models.OpponentRecruiting
	result := db.First(&opponentRecruiting, id)
	if result.Error != nil {
		return nil, errors.New("データ取得に失敗しました")
	}
	return &opponentRecruiting, nil
}

// 対戦相手募集をコメントも含めて取得する(なければエラー)
func (ors *OpponentRecruitingService) FindOpponentRecruitingWithComment(db *gorm.DB, id uint) (*models.OpponentRecruiting, error) {
	var opponentRecruiting models.OpponentRecruiting
	result := db.Preload("Team").Preload("Comments.User").Preload("Comments.Team").First(&opponentRecruiting, id)
	if result.Error != nil {
		return nil, errors.New("データ取得に失敗しました")
	}
	return &opponentRecruiting, nil
}

// 対戦相手募集を削除する
func (ors *OpponentRecruitingService) DeleteOpponentRecruiting(db *gorm.DB, userId string, id uint) error {
	// 削除する対戦相手募集を取得する
	opponentRecruiting, err := ors.FindOpponentRecruiting(db, id)
	if err != nil {
		return err
	}
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(db, userId, opponentRecruiting.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集を削除できます")
	}

	// 対戦相手募集を削除する
	result := db.Unscoped().Delete(&models.OpponentRecruiting{}, "id = ?", id)
	if result.Error != nil {
		return errors.New("データ削除に失敗しました")
	}

	// 削除したデータが0件の場合はエラー
	if result.RowsAffected == 0 {
		return errors.New("削除対象のデータがありません")
	}
	return nil
}

// 対戦相手募集のリストとページ情報を返す
func (ors *OpponentRecruitingService) GetOpponentRecruitingList(c *gin.Context) ([]*models.OpponentRecruiting, *database.Page) {
	// リスト表示時の1ページあたりの要素数
	var pageSize int = 10

	// 対戦相手募集の構造体の配列
	var opponentRecruitings []*models.OpponentRecruiting
	db := c.MustGet("tx").(*gorm.DB)

	// クエリパラメータからフィルタリング条件を取得
	hasGroundQuery := c.Query("has_ground")
	prefectureId, _ := strconv.Atoi(c.Query("prefecture_id"))
	isActive := c.Query("is_active") == "true"
	date := c.Query("date")
	day := c.Query("day")

	// グラウンドの有無でフィルタリング
	if hasGroundQuery != "" {
		hasGround, err := strconv.ParseBool(hasGroundQuery)
		if err == nil {
			db = db.Where("has_ground = ?", hasGround)
		}
	}
	// 都道府県でフィルタリング
	if prefectureId > 0 {
		db = db.Where("prefecture_id = ?", prefectureId)
	}
	// 募集中かどうか
	if isActive {
		db = db.Where("is_active = ?", true)
	}
	// 日付でフィルタリング
	if date != "" && day == "" {
		db = db.Where("DATE(start_time) = ?", date)
		// 曜日でフィルタリング
	} else if day != "" && date == "" {
		db = db.Where("DAYNAME(start_time) = ?", day)
	}

	// 合計要素数
	var totalElements int64
	db.Model(&models.OpponentRecruiting{}).Count(&totalElements)

	// 合計要素数がページサイズより小さい場合はページサイズを合計要素数に合わせる
	if int(totalElements) < pageSize {
		pageSize = int(totalElements)
		if pageSize <= 0 {
			// 最低でも1ページは存在するようにする
			pageSize = 1
		}
	}

	// 合計ページ数
	totalPages := int(math.Ceil(float64(totalElements) / float64(pageSize)))
	pageNumber, _ := strconv.Atoi(c.Query("page"))

	// 指定されたページ数が合計ページ数を超えていたら合計ページ数に合わせる
	if pageNumber > totalPages {
		pageNumber = totalPages
	}

	page := &database.Page{
		Number:        pageNumber,
		Size:          pageSize,
		TotalElements: int(totalElements),
		TotalPages:    totalPages,
	}

	sort := &database.Sort{
		IsDesc:  true,
		OrderBy: "created_at",
	}

	// 対戦相手募集を指定されたページと作成順に並び替えて、チーム情報とまとめて返す
	db = db.Scopes(page.Paginate(), sort.Sort()).Preload("Team").Find(&opponentRecruitings)

	return opponentRecruitings, page
}

// 対戦相手募集のコメントを作成する
func (ors *OpponentRecruitingService) CreateOpponentRecruitingComment(db *gorm.DB, opponentRecruitingComment *models.OpponentRecruitingComment) error {
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(db, *opponentRecruitingComment.UserID, *opponentRecruitingComment.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集にコメントできます")
	}
	// 対戦相手募集のコメントを作成する
	result := db.Create(opponentRecruitingComment)
	if result.Error != nil {
		return errors.New("対戦相手募集のコメントに失敗しました")
	}
	return nil
}

// 対戦相手募集のコメントを変更する
func (ors *OpponentRecruitingService) UpdateOpponentRecruitingComment(db *gorm.DB, userId string, opponentRecruitingComment *models.OpponentRecruitingComment) error {
	// 変更する対戦相手募集のコメントを取得する
	originalOpponentRecruitingComment, err := ors.findOpponentRecruitingComment(db, opponentRecruitingComment.ID)
	if err != nil {
		return err
	}
	// そのユーザーのコメントでなければエラー
	if !ors.isUserOpponentRecruitingComment(*originalOpponentRecruitingComment, userId) {
		return errors.New("自分のコメントしか変更できません")
	}

	// データを更新する
	result := db.Model(&models.OpponentRecruitingComment{}).Updates(opponentRecruitingComment)
	if result.Error != nil {
		return result.Error
	}
	// 更新したデータが0件の場合はエラー
	if result.RowsAffected == 0 {
		return errors.New("更新対象のデータがありません")
	}
	return nil
}

// 対戦相手募集のコメントを取得する
func (ors *OpponentRecruitingService) findOpponentRecruitingComment(db *gorm.DB, id uint) (*models.OpponentRecruitingComment, error) {
	var opponentRecruitingComment models.OpponentRecruitingComment
	result := db.First(&opponentRecruitingComment, id)
	if result.Error != nil {
		return nil, errors.New("データ取得に失敗しました")
	}
	return &opponentRecruitingComment, nil
}

// そのユーザーの対戦相手募集のコメントかどうか
func (ors *OpponentRecruitingService) isUserOpponentRecruitingComment(opponentRecruitingComment models.OpponentRecruitingComment, userId string) bool {
	// コメントがそのユーザーのものでなければfalse
	if *opponentRecruitingComment.UserID != userId {
		return false
	}
	return true
}

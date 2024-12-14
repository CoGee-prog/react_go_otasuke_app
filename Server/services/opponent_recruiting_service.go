package services

import (
	"errors"
	"math"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OpponentRecruitingService interface {
	CreateOpponentRecruiting(tx *gorm.DB, userId string, opponentRecruiting *models.OpponentRecruiting) error
	UpdateOpponentRecruiting(tx *gorm.DB, userId string, id uint, opponentRecruiting *models.OpponentRecruiting) error
	UpdateStatusOpponentRecruiting(tx *gorm.DB, userId string, id uint, opponentRecruiting *models.OpponentRecruiting) error
	FindOpponentRecruiting(tx *gorm.DB, id uint) (*models.OpponentRecruiting, error)
	FindOpponentRecruitingWithComment(tx *gorm.DB, id uint) (*models.OpponentRecruiting, error)
	DeleteOpponentRecruiting(tx *gorm.DB, userId string, id uint) error
	GetOpponentRecruitingList(c *gin.Context, isMyTeam bool) ([]*models.OpponentRecruiting, *database.Page, error)
	CreateOpponentRecruitingComment(tx *gorm.DB, opponentRecruitingComment *models.OpponentRecruitingComment) error
	UpdateOpponentRecruitingComment(tx *gorm.DB, userId string, opponentRecruitingComment *models.OpponentRecruitingComment) error
	findOpponentRecruitingComment(tx *gorm.DB, id uint) (*models.OpponentRecruitingComment, error)
	isUserOpponentRecruitingComment(opponentRecruitingComment models.OpponentRecruitingComment, userId string) bool
	DeleteOpponentRecruitingComment(tx *gorm.DB, userId string, id uint) error
	softDeleteOpponentRecruitingComment(tx *gorm.DB, opponentRecruitingComment *models.OpponentRecruitingComment) error
}

type opponentRecruitingService struct {
	userTeamService                     UserTeamService
	userRepository                      repositories.UserRepository
	opponentRecruitingRepository        repositories.OpponentRecruitingRepository
	opponentRecruitingCommentRepository repositories.OpponentRecruitingCommentRepository
}

// 対戦相手募集のサービスを作成する
func NewOpponentRecruitingService(
	uts UserTeamService,
	userRepo repositories.UserRepository,
	opponentRecruitingRepo repositories.OpponentRecruitingRepository,
	opponentRecruitingCommentRepo repositories.OpponentRecruitingCommentRepository,
) OpponentRecruitingService {
	return &opponentRecruitingService{
		userTeamService:                     uts,
		userRepository:                      userRepo,
		opponentRecruitingRepository:        opponentRecruitingRepo,
		opponentRecruitingCommentRepository: opponentRecruitingCommentRepo,
	}
}

// 対戦相手募集を作成する
func (ors *opponentRecruitingService) CreateOpponentRecruiting(tx *gorm.DB, userId string, opponentRecruiting *models.OpponentRecruiting) error {
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(tx, userId, opponentRecruiting.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集を作成できます")
	}
	// 対戦相手募集を作成する
	return ors.opponentRecruitingRepository.Create(tx, opponentRecruiting)
}

// 対戦相手募集を変更する
func (ors *opponentRecruitingService) UpdateOpponentRecruiting(tx *gorm.DB, userId string, id uint, opponentRecruiting *models.OpponentRecruiting) error {
	// 変更する対戦相手募集を取得する
	originalOpponentRecruiting, err := ors.FindOpponentRecruiting(tx, id)
	if err != nil {
		return err
	}
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(tx, userId, originalOpponentRecruiting.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集を変更できます")
	}

	// データを更新する
	return ors.opponentRecruitingRepository.UpdateById(tx, id, opponentRecruiting)
}

// 対戦相手募集の状態(募集中かどうか)を変更する
func (ors *opponentRecruitingService) UpdateStatusOpponentRecruiting(tx *gorm.DB, userId string, id uint, opponentRecruiting *models.OpponentRecruiting) error {
	// 変更する対戦相手募集を取得する
	originalOpponentRecruiting, err := ors.FindOpponentRecruiting(tx, id)
	if err != nil {
		return err
	}
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(tx, userId, originalOpponentRecruiting.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集を変更できます")
	}

	// データを更新する
	err = ors.opponentRecruitingRepository.UpdateIsActive(tx, id, opponentRecruiting)

	if err != nil {
		return err
	}

	return nil
}

// 対戦相手募集を取得する(なければエラー)
func (ors *opponentRecruitingService) FindOpponentRecruiting(tx *gorm.DB, id uint) (*models.OpponentRecruiting, error) {
	return ors.opponentRecruitingRepository.FindById(tx, id)
}

// 対戦相手募集をコメントも含めて取得する(なければエラー)
func (ors *opponentRecruitingService) FindOpponentRecruitingWithComment(tx *gorm.DB, id uint) (*models.OpponentRecruiting, error) {
	return ors.opponentRecruitingRepository.FindByIdWithComments(tx, id)
}

// 対戦相手募集を削除する
func (ors *opponentRecruitingService) DeleteOpponentRecruiting(tx *gorm.DB, userId string, id uint) error {
	// 削除する対戦相手募集を取得する
	opponentRecruiting, err := ors.FindOpponentRecruiting(tx, id)
	if err != nil {
		return err
	}
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(tx, userId, opponentRecruiting.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集を削除できます")
	}

	// 対戦相手募集を削除する
	return ors.opponentRecruitingRepository.DeleteById(tx, id)
}

// 対戦相手募集のリストとページ情報を返す
func (ors *opponentRecruitingService) GetOpponentRecruitingList(c *gin.Context, isMyTeam bool) ([]*models.OpponentRecruiting, *database.Page, error) {
	// リスト表示時の1ページあたりの要素数
	var pageSize int = 10

	// 対戦相手募集の構造体の配列
	var opponentRecruitings []*models.OpponentRecruiting
	tx := c.MustGet("tx").(*gorm.DB)

	// 自チームフラグが立っている場合
	if isMyTeam {
		//ユーザーを取得する
		userId := c.MustGet("userId").(string)
		user, err := ors.userRepository.GetByUserId(tx, userId)

		// エラーが起きた場合
		if err != nil {
			return nil, nil, err
		}
		
		// 自チームの対戦相手募集に絞り込む
		myTeamId := user.CurrentTeamId

		// チームの副管理者以上でなければエラー
		if !ors.userTeamService.IsAdminOrSubAdmin(tx, userId, *myTeamId) {
			return nil, nil, errors.New("管理者または副管理者のみ自チームの対戦相手募集一覧を閲覧できます")
		}

		tx = tx.Where("team_id = ?", myTeamId)
	}

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
			tx = tx.Where("has_ground = ?", hasGround)
		}
	}
	// 都道府県でフィルタリング
	if prefectureId > 0 {
		tx = tx.Where("prefecture_id = ?", prefectureId)
	}
	// 募集中かどうか
	if isActive {
		tx = tx.Where("is_active = ?", true)
	}
	// 日付でフィルタリング
	if date != "" && day == "" {
		tx = tx.Where("DATE(start_time) = ?", date)
		// 曜日でフィルタリング
	} else if day != "" && date == "" {
		tx = tx.Where("DAYNAME(start_time) = ?", day)
	}

	// 対戦相手募集の合計要素数を取得
	totalElementsCount, err := ors.opponentRecruitingRepository.GetTotalCount(tx)
	//エラーが起きた場合
	if err != nil {
		return nil, nil, err
	}

	// 合計要素数がページサイズより小さい場合はページサイズを合計要素数に合わせる
	if int(totalElementsCount) < pageSize {
		pageSize = int(totalElementsCount)
		if pageSize <= 0 {
			// 最低でも1ページは存在するようにする
			pageSize = 1
		}
	}

	// 合計ページ数
	totalPages := int(math.Ceil(float64(totalElementsCount) / float64(pageSize)))
	pageNumber, _ := strconv.Atoi(c.Query("page"))

	// 指定されたページ数が合計ページ数を超えていたら合計ページ数に合わせる
	if pageNumber > totalPages {
		pageNumber = totalPages
	}

	page := &database.Page{
		Number:        pageNumber,
		Size:          pageSize,
		TotalElements: int(totalElementsCount),
		TotalPages:    totalPages,
	}

	sort := &database.Sort{
		IsDesc:  true,
		OrderBy: "created_at",
	}

	// 対戦相手募集を指定されたページと作成順に並び替えて、チーム情報とまとめて返す
	opponentRecruitings, err = ors.opponentRecruitingRepository.GetListWithTeamByPaginate(tx, page, sort)

	if err != nil {
		return nil, nil, err
	}

	return opponentRecruitings, page, nil
}

// 対戦相手募集のコメントを作成する
func (ors *opponentRecruitingService) CreateOpponentRecruitingComment(tx *gorm.DB, opponentRecruitingComment *models.OpponentRecruitingComment) error {
	// チームの管理者または副管理者でなければエラー
	if !ors.userTeamService.IsAdminOrSubAdmin(tx, *opponentRecruitingComment.UserID, *opponentRecruitingComment.TeamID) {
		return errors.New("管理者または副管理者のみ対戦相手募集にコメントできます")
	}
	// 対戦相手募集のコメントを作成する
	err := ors.opponentRecruitingCommentRepository.Create(tx, opponentRecruitingComment)
	if err != nil {
		return errors.New("対戦相手募集のコメントに失敗しました")
	}
	return nil
}

// 対戦相手募集のコメントを変更する
func (ors *opponentRecruitingService) UpdateOpponentRecruitingComment(tx *gorm.DB, userId string, opponentRecruitingComment *models.OpponentRecruitingComment) error {
	// 変更する対戦相手募集のコメントを取得する
	originalOpponentRecruitingComment, err := ors.findOpponentRecruitingComment(tx, opponentRecruitingComment.ID)
	if err != nil {
		return err
	}
	// そのユーザーのコメントでなければエラー
	if !ors.isUserOpponentRecruitingComment(*originalOpponentRecruitingComment, userId) {
		return errors.New("自分のコメントしか変更できません")
	}

	// データを更新する
	return ors.opponentRecruitingCommentRepository.UpdateById(tx, originalOpponentRecruitingComment.ID, opponentRecruitingComment)
}

// 対戦相手募集のコメントを取得する
func (ors *opponentRecruitingService) findOpponentRecruitingComment(tx *gorm.DB, id uint) (*models.OpponentRecruitingComment, error) {
	return ors.opponentRecruitingCommentRepository.FindById(tx, id)
}

// そのユーザーの対戦相手募集のコメントかどうか
func (ors *opponentRecruitingService) isUserOpponentRecruitingComment(opponentRecruitingComment models.OpponentRecruitingComment, userId string) bool {
	// コメントがそのユーザーのものでなければfalse
	if *opponentRecruitingComment.UserID != userId {
		return false
	}
	return true
}

// 対戦相手募集のコメントを削除する
func (ors *opponentRecruitingService) DeleteOpponentRecruitingComment(tx *gorm.DB, userId string, id uint) error {
	// 削除する対戦相手募集を取得する
	opponentRecruitingComment, err := ors.findOpponentRecruitingComment(tx, id)
	if err != nil {
		return err
	}

	// そのユーザーのコメントでなければエラー
	if !ors.isUserOpponentRecruitingComment(*opponentRecruitingComment, userId) {
		return errors.New("自分のコメントしか削除できません")
	}

	// 対戦相手募集を削除済みにする
	if err = ors.softDeleteOpponentRecruitingComment(tx, opponentRecruitingComment); err != nil {
		return err
	}
	return nil
}

// 対戦相手募集コメントの削除済みフラグを立てて内容を削除する
func (ors *opponentRecruitingService) softDeleteOpponentRecruitingComment(tx *gorm.DB, opponentRecruitingComment *models.OpponentRecruitingComment) error {
	return ors.opponentRecruitingCommentRepository.SetDeleteFlagById(tx, opponentRecruitingComment.ID)
}

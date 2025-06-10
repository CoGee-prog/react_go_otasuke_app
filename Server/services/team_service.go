package services

import (
	"errors"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/repositories"
	"react_go_otasuke_app/utils"
	"time"

	"gorm.io/gorm"
)

type TeamService interface {
	GetTeam(tx *gorm.DB, teamId string) (*models.Team, error)
	CreateTeamWithAdmin(tx *gorm.DB, userId string, team *models.Team) error
	UpdateTeam(tx *gorm.DB, userId string, teamId string, team *models.Team) error
	CreateInviteToken(tx *gorm.DB, userId string, team *models.Team) (*models.TeamInvite, error)
	GetTeamByInviteToken(tx *gorm.DB, token string) (*models.Team, error)
	JoinTeamByInviteToken(tx *gorm.DB, userId string, token string) (*models.Team, error)
}

type teamService struct {
	userTeamService      UserTeamService
	teamRepository       repositories.TeamRepository
	userTeamRepository   repositories.UserTeamRepository
	teamInviteRepository repositories.TeamInviteRepository
}

// チームサービスを作成する
func NewTeamService(uts UserTeamService, teamRepo repositories.TeamRepository, userTeamRepo repositories.UserTeamRepository, teamInviteRepo repositories.TeamInviteRepository) TeamService {
	return &teamService{
		userTeamService:      uts,
		teamRepository:       teamRepo,
		userTeamRepository:   userTeamRepo,
		teamInviteRepository: teamInviteRepo,
	}
}

// チームを取得する
func (ts *teamService) GetTeam(tx *gorm.DB, teamId string) (*models.Team, error) {
	return ts.teamRepository.GetById(tx, teamId)
}

// チームを作成し、チーム作成者をチーム管理者とする
func (ts *teamService) CreateTeamWithAdmin(tx *gorm.DB, userId string, team *models.Team) error {
	// チームIDとして20文字のランダムな文字列を生成
	team.ID, _ = utils.GenerateRandomString(20)

	var err error
	// チーム作成できたかどうか
	isTeamCreated := false

	// 万が一チームIDが被った時に備えて3回までリトライする
	for attempt := 1; attempt <= 3; attempt++ {
		// チームを作成する
		if err = ts.teamRepository.Create(tx, team); err != nil {
			continue
		}
		isTeamCreated = true
		break
	}

	// チーム作成できなかった場合
	if !isTeamCreated {
		return err
	}

	// チーム作成者をチーム管理者としてチームに所属させる
	if err := ts.userTeamRepository.AddTeamAdmin(tx, userId, team); err != nil {
		return err
	}

	return nil
}

// チームを更新する
func (ts *teamService) UpdateTeam(tx *gorm.DB, userId string, teamId string, team *models.Team) error {
	// チームの管理者または副管理者でなければエラー
	if !ts.userTeamService.IsAdminOrSubAdmin(tx, userId, teamId) {
		return errors.New("管理者または副管理者のみ対戦相手募集を変更できます")
	}

	// データを更新する
	return ts.teamRepository.UpdateById(tx, teamId, team)
}

func (ts *teamService) CreateInviteToken(tx *gorm.DB, userId string, team *models.Team) (*models.TeamInvite, error) {
	// チームの管理者または副管理者でなければエラー
	if !ts.userTeamService.IsAdminOrSubAdmin(tx, userId, team.ID) {
		return nil, errors.New("管理者または副管理者のみ招待トークンを発行できます")
	}

	// トークンを生成する
	token, _ := utils.GenerateRandomString(30)

	// 招待トークンを作成する
	return ts.teamInviteRepository.CreateOrUpdateInvite(tx, team.ID, token)
}

func (ts *teamService) GetTeamByInviteToken(tx *gorm.DB, token string) (*models.Team, error) {
	invite, err := ts.teamInviteRepository.GetByToken(tx, token)
	if err != nil {
		return nil, err
	}
	if invite == nil || invite.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("招待URLが無効です")
	}

	return ts.teamRepository.GetById(tx, invite.TeamID)
}

func (ts *teamService) JoinTeamByInviteToken(tx *gorm.DB, userId string, token string) (*models.Team, error) {
	team, err := ts.GetTeamByInviteToken(tx, token)
	if err != nil || team == nil {
		return nil, err
	}

	// すでに所属しているか確認
	userTeam, err := ts.userTeamRepository.GetByUserIdAndTeamId(tx, userId, team.ID)
	if err != nil {
		return nil, err
	}
	if userTeam != nil {
		return nil, errors.New("既にチームに所属しています")
	}

	if err := ts.userTeamRepository.AddTeamMember(tx, userId, team); err != nil {
		return nil, err
	}

	return team, nil
}

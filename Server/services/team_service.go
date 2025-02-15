package services

import (
	"errors"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/repositories"

	"gorm.io/gorm"
)

type TeamService interface {
	GetTeam(tx *gorm.DB, id uint) (*models.Team, error)
	CreateTeamWithAdmin(tx *gorm.DB, userId string, team *models.Team) error
	UpdateTeam(tx *gorm.DB, userId string, id uint, team *models.Team) error
}

type teamService struct {
	userTeamService    UserTeamService
	teamRepository     repositories.TeamRepository
	userTeamRepository repositories.UserTeamRepository
}

// チームサービスを作成する
func NewTeamService(uts UserTeamService, teamRepo repositories.TeamRepository, userTeamRepo repositories.UserTeamRepository) TeamService {
	return &teamService{
		userTeamService:    uts,
		teamRepository:     teamRepo,
		userTeamRepository: userTeamRepo,
	}
}

// チームを取得する
func (ts *teamService) GetTeam(tx *gorm.DB, id uint) (*models.Team, error) {
	return ts.teamRepository.GetById(tx, id)
}

// チームを作成し、チーム作成者をチーム管理者とする
func (ts *teamService) CreateTeamWithAdmin(tx *gorm.DB, userId string, team *models.Team) error {
	// チームを作成する
	if err := ts.teamRepository.Create(tx, team); err != nil {
		return err
	}

	// チーム作成者をチーム管理者としてチームに所属させる
	if err := ts.userTeamRepository.AddTeamAdmin(tx, userId, team); err != nil {
		return err
	}

	return nil
}

// チームを更新する
func (ts *teamService) UpdateTeam(tx *gorm.DB, userId string, teamId uint, team *models.Team) error {
	// チームの管理者または副管理者でなければエラー
	if !ts.userTeamService.IsAdminOrSubAdmin(tx, userId, teamId) {
		return errors.New("管理者または副管理者のみ対戦相手募集を変更できます")
	}

	// データを更新する
	return ts.teamRepository.UpdateById(tx, teamId, team)
}

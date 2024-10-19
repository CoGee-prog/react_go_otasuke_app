package services

import (
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/repositories"

	"gorm.io/gorm"
)

type TeamService interface {
	GetTeam(tx *gorm.DB, id string) (*models.Team, error)
	CreateTeamWithAdmin(tx *gorm.DB, userId string, team *models.Team) error
}

type teamService struct {
	teamRepository repositories.TeamRepository
	userTeamRepository repositories.UserTeamRepository
}

// チームサービスを作成する
func NewTeamService(teamRepo repositories.TeamRepository, userTeamRepo repositories.UserTeamRepository) TeamService {
	return &teamService{
		teamRepository: teamRepo,
		userTeamRepository: userTeamRepo,
	}
}

// チームを取得する
func (ts *teamService) GetTeam(tx *gorm.DB, id string) (*models.Team, error) {
	return ts.teamRepository.GetTeam(tx, id)
}

// チームを作成し、チーム作成者をチーム管理者とする
func (ts *teamService) CreateTeamWithAdmin(tx *gorm.DB, userId string, team *models.Team) error {
	// チームを作成する
	if err := ts.teamRepository.CreateTeam(tx, team); err != nil {
		return err
	}

	// チーム作成者をチーム管理者としてチームに所属させる
	if err := ts.userTeamRepository.AddTeamAdmin(tx, userId, team); err != nil {
		return err
	}

	return nil
}

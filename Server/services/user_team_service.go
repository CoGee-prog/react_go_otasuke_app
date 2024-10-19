package services

import (
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/repositories"

	"gorm.io/gorm"
)

type UserTeamService interface {
	GetUserTeam(db *gorm.DB, userID string, teamId uint) (*models.UserTeam, error)
	IsAdminOrSubAdmin(db *gorm.DB, userId string, teamId uint) bool
}

type userTeamService struct {
	userTeamRepository repositories.UserTeamRepository
}

// ユーザーチームサービスを作成する
func NewUserTeamService(userTeamRepo repositories.UserTeamRepository) UserTeamService {
	return &userTeamService{
		userTeamRepository: userTeamRepo,
	}
}

// ユーザーチームを取得する
func (uts *userTeamService) GetUserTeam(db *gorm.DB, userId string, teamId uint) (*models.UserTeam, error) {
	return uts.userTeamRepository.GetUserTeam(db, userId, teamId)
}

// ユーザーが管理者または副管理者かどうか
func (uts *userTeamService) IsAdminOrSubAdmin(db *gorm.DB, userId string, teamId uint) bool {
	// ユーザーのチームを取得する
	userTeam, err := uts.userTeamRepository.GetUserTeam(db, userId, teamId)
	// チームに所属していなければfalse
	if err != nil {
		return false
	}

	return userTeam.IsAdminOrSubAdmin()
}

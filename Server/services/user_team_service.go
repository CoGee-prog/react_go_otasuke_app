package services

import (
	"react_go_otasuke_app/repositories"

	"gorm.io/gorm"
)

type UserTeamService interface {
	IsAdminOrSubAdmin(db *gorm.DB, userId string, teamId string) bool
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

// ユーザーが管理者または副管理者かどうか
func (uts *userTeamService) IsAdminOrSubAdmin(db *gorm.DB, userId string, teamId string) bool {
	// ユーザーのチームを取得する
	userTeam, err := uts.userTeamRepository.GetByUserIdAndTeamId(db, userId, teamId)
	// レコードが見つからないかチームに所属していなければfalse
	if userTeam == nil || err != nil {
		return false
	}

	return userTeam.IsAdminOrSubAdmin()
}

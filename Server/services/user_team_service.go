package services

import (
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/utils"
)

type UserTeamService struct {
	db *database.GormDatabase
}

// ユーザーチームサービスを作成する
func NewUserTeamService(db *database.GormDatabase) *UserTeamService {
	return &UserTeamService{
		db: db,
	}
}

// ユーザーチームを取得する
func (uts *UserTeamService) GetUserTeam(userID string, teamID uint) (*models.UserTeam, error) {
	var userTeam models.UserTeam
	err := uts.db.DB.Where("user_id = ? AND team_id = ?", userID, teamID).First(&userTeam).Error
	if err != nil {
		return nil, err
	}
	return &userTeam, nil
}

// ユーザーが管理者または副管理者かどうか
func (uts *UserTeamService) IsAdminOrSubAdmin(teamID uint) bool {
	// ユーザーのチームを取得する
	userTeam, err := uts.GetUserTeam(utils.GetUserID(), teamID)
	// チームに所属していなければfalse
	if err != nil {
		return false
	}

	return userTeam.IsAdminOrSubAdmin()
}

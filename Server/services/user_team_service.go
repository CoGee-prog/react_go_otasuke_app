package services

import (
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/utils"

	"gorm.io/gorm"
)

type UserTeamService struct{}

// ユーザーチームサービスを作成する
func NewUserTeamService() *UserTeamService {
	return &UserTeamService{}
}

// ユーザーチームを取得する
func (uts *UserTeamService) GetUserTeam(db *gorm.DB, userID string, teamID uint) (*models.UserTeam, error) {
	var userTeam models.UserTeam
	err := db.Where("user_id = ? AND team_id = ?", userID, teamID).First(&userTeam).Error
	if err != nil {
		return nil, err
	}
	return &userTeam, nil
}

// ユーザーが管理者または副管理者かどうか
func (uts *UserTeamService) IsAdminOrSubAdmin(db *gorm.DB, teamID uint) bool {
	// ユーザーのチームを取得する
	userTeam, err := uts.GetUserTeam(db, utils.GetUserID(), teamID)
	// チームに所属していなければfalse
	if err != nil {
		return false
	}

	return userTeam.IsAdminOrSubAdmin()
}

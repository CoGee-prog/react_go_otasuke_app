package services

import (
	"errors"
	"react_go_otasuke_app/models"

	"gorm.io/gorm"
)

type TeamService struct{}

// チームサービスを作成する
func NewTeamService() *TeamService {
	return &TeamService{}
}

// チームを取得する
func (ts *TeamService) GetTeam(db *gorm.DB, id string) (*models.Team, error) {
	var team models.Team
	result := db.Where("id = ?", id).First(&team)
	// レコードが見つからない場合はnilを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
		// その他のエラーの場合
	} else if result.Error != nil {
		return nil, result.Error
	}
	// レコードが見つかった場合
	return &team, nil
}

// チームを作成する
func (ts *TeamService) CreateTeam(db *gorm.DB, userId string, team *models.Team) error {
	// チームを作成する
	if err := db.Create(team).Error; err != nil {
		return err
	}

	// チーム作成者をチーム管理者としてチームに所属させる
	userTeam := models.UserTeam{
		UserID: userId,
		TeamID: team.ID,
		Role:   models.TeamAdmin,
	}
	if err := db.Create(&userTeam).Error; err != nil {
		return err
	}

	return nil
}

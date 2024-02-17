package services

import (
	"errors"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/utils"

	"gorm.io/gorm"
)

type TeamService struct {
	db *database.GormDatabase
}

// チームサービスを作成する
func NewTeamService(db *database.GormDatabase) *TeamService {
	return &TeamService{
		db: db,
	}
}

// チームを取得する
func (ts *TeamService) GetTeam(id string) (*models.Team, error) {
	var team models.Team
	result := ts.db.DB.Where("id = ?", id).First(&team)
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
func (ts *TeamService) CreateTeam(team *models.Team) error {
	// チームを作成する
	if err := ts.db.DB.Create(team).Error; err != nil {
		return err
	}

	// チーム作成者をチーム管理者としてチームに所属させる
	userTeam := models.UserTeam{
		UserID: utils.GetUserID(),
		TeamID: team.ID,
		Role:   models.TeamAdmin,
	}
	if err := ts.db.DB.Create(&userTeam).Error; err != nil {
		return err
	}

	return nil
}

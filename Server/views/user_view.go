package views

import (
	"react_go_otasuke_app/models"
)

type UserView struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	CurrentTeamId   *string          `json:"current_team_id"`
	CurrentTeamName string           `json:"current_team_name"`
	CurrentTeamRole *models.TeamRole `json:"current_team_role"`
}

// Userの構造体から必要なキーのみ返す
func CreateUserView(user *models.User) *UserView {
	userView := &UserView{
		ID:            user.ID,
		Name:          user.Name,
		CurrentTeamId: user.CurrentTeamId,
	}

	// CurrentTeamがnilでなければチーム情報を設定
	if user.CurrentTeam != nil {
		userView.CurrentTeamName = user.CurrentTeam.Name
		userView.CurrentTeamRole = user.CurrentTeamRole
	}

	return userView
}

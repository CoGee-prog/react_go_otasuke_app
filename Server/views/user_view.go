package views

import (
	"react_go_otasuke_app/models"
)

type UserView struct {
	Name          string `json:"name"`
	CurrentTeamId *int   `json:"current_team_id"`
}

// Userの構造体から必要なキーのみ返す
func CreateUserView(user *models.User) *UserView {
	return &UserView{
		Name:          user.Name,
		CurrentTeamId: user.CurrentTeamId,
	}
}

package views

import (
	"react_go_otasuke_app/models"
)

type OpponentRecruitingCommentView struct {
	OpponentRecruitingID uint    `json:"opponent_recruiting_id"`
	UserID               *string `json:"user_id"`
	UserName             *string `json:"user_name"`
	TeamID               *uint   `json:"team_id"`
	TeamName             *string `json:"team_name"`
	Content              string  `json:"content"`
}

// 対戦相手募集のコメント構造体から必要なキーのみ返す
func CreateOpponentRecruitingCommentView(opponentRecruitingComments []*models.OpponentRecruitingComment) []*OpponentRecruitingCommentView {
	newArray := make([]*OpponentRecruitingCommentView, len(opponentRecruitingComments))
	for i, v := range opponentRecruitingComments {
		newArray[i] = &OpponentRecruitingCommentView{
			OpponentRecruitingID: v.ID,
			UserID:               v.UserID,
			UserName:             &v.User.Name,
			TeamID:               v.TeamID,
			TeamName:             &v.Team.Name,
			Content:              v.Content,
		}
	}

	return newArray
}

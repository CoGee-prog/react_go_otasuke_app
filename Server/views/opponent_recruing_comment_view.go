package views

import (
	"react_go_otasuke_app/models"
)

type OpponentRecruitingCommentView struct {
	ID                   uint    `json:"id"`
	OpponentRecruitingID uint    `json:"opponent_recruiting_id"`
	UserID               *string `json:"user_id"`
	UserName             *string `json:"user_name"`
	TeamID               *uint   `json:"team_id"`
	TeamName             *string `json:"team_name"`
	Content              string  `json:"content"`
	Edited               bool    `json:"edited"`
	Deleted              bool    `json:"deleted"`
}

// 対戦相手募集のコメント構造体から必要なキーのみ返す
func CreateOpponentRecruitingCommentView(opponentRecruitingComments []*models.OpponentRecruitingComment) []*OpponentRecruitingCommentView {
	newArray := make([]*OpponentRecruitingCommentView, len(opponentRecruitingComments))
	for i, v := range opponentRecruitingComments {
		newArray[i] = &OpponentRecruitingCommentView{
			ID:                   v.ID,
			OpponentRecruitingID: v.OpponentRecruitingID,
			UserID:               v.UserID,
			UserName:             &v.User.Name,
			TeamID:               v.TeamID,
			TeamName:             &v.Team.Name,
			Content:              v.Content,
			Edited:               v.CreatedAt != v.UpdatedAt, // 作成日時と更新日時が違う場合は編集済みのコメント
			Deleted:              v.Deleted,
		}
	}

	return newArray
}

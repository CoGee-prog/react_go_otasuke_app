package views

import (
	"react_go_otasuke_app/models"
	"time"
)

type OpponentRecruitingView struct {
	ID       uint      `json:"id"`
	TeamId   uint       `json:"team_id"`
	DateTime time.Time `json:"date_time"`
	AreaId   uint       `json:"area"`
	Detail   *string   `json:"detail"`
}

// 対戦相手募集の構造体から必要なキーのみ返す
func CreateOpponentRecruitingView(opponentRecruitings []*models.OpponentRecruiting) []*OpponentRecruitingView {
	newArray := make([]*OpponentRecruitingView, len(opponentRecruitings))
	for i, v := range opponentRecruitings {
		newArray[i] = &OpponentRecruitingView{
			ID:       v.ID,
			TeamId:   v.TeamId,
			DateTime: v.DateTime,
			AreaId:   v.AreaId,
			Detail:   v.Detail,
		}
	}

	return newArray
}

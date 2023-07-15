package views

import (
	"react_go_otasuke_app/models"
	"time"
)

type OpponentRecruitingView struct {
	ID       uint      `json:"id"`
	TeamId   int       `json:"team_id"`
	DateTime time.Time `json:"date_time"`
	AreaId   int       `json:"area"`
	Detail   *string   `json:"detail"`
}

func IndexOpponentRecruitingView(opponentRecruiting []*models.OpponentRecruiting) []*OpponentRecruitingView {
	newArray := make([]*OpponentRecruitingView, len(opponentRecruiting))
	for i, v := range opponentRecruiting {
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

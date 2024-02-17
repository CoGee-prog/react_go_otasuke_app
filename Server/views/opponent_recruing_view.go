package views

import (
	"react_go_otasuke_app/models"
	"time"
)

type OpponentRecruitingView struct {
	ID         uint      `json:"id"`
	Team       *TeamView `json:"team"`
	DateTime   time.Time `json:"date_time"`
	Prefecture string    `json:"prefecture"`
	Detail     *string   `json:"detail"`
}

// 対戦相手募集の構造体から必要なキーのみ返す
func CreateOpponentRecruitingView(opponentRecruitings []*models.OpponentRecruiting) []*OpponentRecruitingView {
	newArray := make([]*OpponentRecruitingView, len(opponentRecruitings))
	for i, v := range opponentRecruitings {
		newArray[i] = &OpponentRecruitingView{
			ID:         v.ID,
			Team:       CreateTeamView(v.Team),
			DateTime:   v.DateTime,
			Prefecture: v.PrefectureId.ToString(),
			Detail:     v.Detail,
		}
	}

	return newArray
}

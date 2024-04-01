package views

import (
	"react_go_otasuke_app/models"
	"time"
)

type OpponentRecruitingIndexView struct {
	ID         uint      `json:"id"`
	Team       *TeamView `json:"team"`
	Title      string    `json:"title"`
	HasGround  bool      `json:"has_ground"`
	GroundName string    `json:"ground_name"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Prefecture string    `json:"prefecture"`
	Detail     string    `json:"detail"`
	IsActive   bool      `json:"is_active"`
}

// 対戦相手募集の構造体から対戦相手募集一覧表示に必要なキーのみ返す
func CreateOpponentRecruitingIndexView(opponentRecruitings []*models.OpponentRecruiting) []*OpponentRecruitingIndexView {
	newArray := make([]*OpponentRecruitingIndexView, len(opponentRecruitings))
	for i, v := range opponentRecruitings {
		newArray[i] = &OpponentRecruitingIndexView{
			ID:         v.ID,
			Team:       CreateTeamView(v.Team),
			Title:      v.Title,
			HasGround:  v.HasGround,
			GroundName: v.GroundName,
			StartTime:  v.StartTime,
			EndTime:    v.EndTime,
			Prefecture: v.PrefectureID.ToString(),
			Detail:     v.Detail,
			IsActive:   v.IsActive,
		}
	}

	return newArray
}

type OpponentRecruitingGetView struct {
	ID         uint                             `json:"id"`
	Team       *TeamView                        `json:"team"`
	Title      string                           `json:"title"`
	HasGround  bool                             `json:"has_ground"`
	GroundName string                           `json:"ground_name"`
	StartTime  time.Time                        `json:"start_time"`
	EndTime    time.Time                        `json:"end_time"`
	Prefecture string                           `json:"prefecture"`
	Detail     string                           `json:"detail"`
	IsActive   bool                             `json:"is_active"`
	Comments   []*OpponentRecruitingCommentView `json:"comments"`
}

// 対戦相手募集の構造体から対戦相手募集一覧表示に必要なキーのみ返す
func CreateOpponentRecruitingGetView(opponentRecruiting *models.OpponentRecruiting) *OpponentRecruitingGetView {
	return &OpponentRecruitingGetView{
		ID:         opponentRecruiting.ID,
		Team:       CreateTeamView(opponentRecruiting.Team),
		Title:      opponentRecruiting.Title,
		HasGround:  opponentRecruiting.HasGround,
		GroundName: opponentRecruiting.GroundName,
		StartTime:  opponentRecruiting.StartTime,
		EndTime:    opponentRecruiting.EndTime,
		Prefecture: opponentRecruiting.PrefectureID.ToString(),
		Detail:     opponentRecruiting.Detail,
		IsActive:   opponentRecruiting.IsActive,
		Comments:   CreateOpponentRecruitingCommentView(opponentRecruiting.Comments),
	}
}

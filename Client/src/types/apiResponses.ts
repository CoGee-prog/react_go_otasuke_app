import { OpponentRecruiting, OpponentRecruitingWithComments } from "./opponentRecruiting";
import { Page } from "./page";
import { User } from "./user";

// ログインのレスポンス
export interface LoginApiResponse {
  user: User;
}

// チーム取得のレスポンス
export interface GetTeamApiResponse {
	name: string,
	prefecture_id: number,
	level_id: number,
	home_page_url: string,
	other: string,
}

// チーム作成のレスポンス
export interface CreateTeamsApiResponse {
  current_team_id: number
	current_team_name: string
	current_team_role: number
}

// チーム更新のレスポンス
export interface UpdateTeamApiResponse {
	name: string,
	prefecture_id: number,
	level_id: number,
	home_page_url: string,
	other: string,
}

// 対戦相手募集一覧のレスポンス
export interface GetOpponentRecruitingsApiResponse {
  opponent_recruitings: OpponentRecruiting[]
	page: Page
}

// 対戦相手募集詳細のレスポンス
export interface GetOpponentRecruitingApiResponse { 
	opponent_recruiting: OpponentRecruitingWithComments
}
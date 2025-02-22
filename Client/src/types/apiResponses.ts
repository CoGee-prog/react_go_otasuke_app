import { OpponentRecruiting, OpponentRecruitingWithComments } from "./opponentRecruiting";
import { Page } from "./page";
import { Team } from "./team";
import { User } from "./user";

// ログインのレスポンス
export interface LoginApiResponse {
  user: User;
}

// チーム取得のレスポンス
export interface GetTeamApiResponse {
	team: Team
}

// チーム作成のレスポンス
export interface CreateTeamsApiResponse {
  current_team_id: number
	current_team_name: string
	current_team_role: number
}

// チーム更新のレスポンス
export interface UpdateTeamApiResponse {
	team: Team
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
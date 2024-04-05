import { OpponentRecruiting, OpponentRecruitingWithComments } from "./opponentRecruiting";
import { Page } from "./page";
import { User } from "./user";

export interface LoginApiResponse {
  user: User;
}

export interface GetOpponentRecruitingsApiResponse {
  opponent_recruitings: OpponentRecruiting[]
	page: Page
}

// チーム作成のレスポンス
export interface CreateTeamsApiResponse {
  current_team_id: number
	current_team_name: string
	current_team_role: number
}

export interface GetOpponentRecruitingApiResponse { 
	opponent_recruiting: OpponentRecruitingWithComments
}
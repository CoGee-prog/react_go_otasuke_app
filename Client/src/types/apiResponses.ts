import { OpponentRecruiting } from "./opponentRecruitings";
import { Page } from "./page";
import { User } from "./user";

export interface loginApiResponse {
  user: User;
}

export interface getOpponentRecruitingsApiResponse {
  opponent_recruitings: OpponentRecruiting
	page: Page
}
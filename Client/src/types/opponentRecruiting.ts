import { Team } from "./team";

export interface OpponentRecruiting {
  id: string;
	team: Team;
	title: string; 
	has_ground: boolean;
	ground_name: string;
	start_time: string;
	end_time: string;
	prefecture: string;
	detail: string;
	is_active: boolean;
}	

export interface OpponentRecruitingWithComments extends OpponentRecruiting {
	comments: OpponentRecruitingComment[]
}

export interface OpponentRecruitingComment {
	id: number
  opponent_recruiting_id: number;
  user_id: string;
  user_name: string;
  team_id: number;
  team_name: string;
  content: string;
  edited: boolean;
  deleted: boolean; 
}
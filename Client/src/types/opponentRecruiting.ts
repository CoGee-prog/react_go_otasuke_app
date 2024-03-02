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
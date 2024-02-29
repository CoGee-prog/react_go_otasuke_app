import { Team } from "./team";

export interface OpponentRecruiting {
  id: string;
	team: Team;
	start_time: string;
	end_time: string;
	prefecture: string;
	detail: string;
}	
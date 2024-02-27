import { Team } from "./team";

export interface OpponentRecruiting {
  id: string;
	team: Team;
	date_time: string;
	prefecture: string;
	detail: string;
}	
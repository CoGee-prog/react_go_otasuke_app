import { TeamRole } from "./teamRole";

export interface User {
  name: string;
  current_team_id: number;
  current_team_name: string;
  current_team_role: TeamRole;
}
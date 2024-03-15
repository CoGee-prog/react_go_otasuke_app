import { TeamRole } from "./teamRole";

export interface User {
  name: string|undefined;
  current_team_id: number|undefined; 
  current_team_name: string|undefined;
  current_team_role: TeamRole|undefined;
}
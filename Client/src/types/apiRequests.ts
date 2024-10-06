export interface CreateOpponentRecruitingsApiRequest {
  title: string
  has_ground: boolean
  ground_name: string
  prefecture_id: string
  start_time: string
  end_time: string
  detail: string
}

export interface CreateTeamsApiRequest {
  name: string
	prefecture_id: number | ''
  level_id: number | ''
  home_page_url: string
  other: string
}
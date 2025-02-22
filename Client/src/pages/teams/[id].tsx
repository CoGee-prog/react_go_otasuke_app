import { GetServerSideProps, GetServerSidePropsContext, NextPage } from 'next'
import OpponentRecruitingDetail from 'src/components/opponent_recruitings/OpponentRecruitingDetail'
import TeamDetail from 'src/components/teams/TeamDetail'
import { GetTeamApiResponse } from 'src/types/apiResponses'
import { Team } from 'src/types/team'
import fetchAPI from 'src/utils/fetchApi'

export const getServerSideProps: GetServerSideProps = async (
  context: GetServerSidePropsContext,
) => {
  const options: RequestInit = {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  }
  const { id } = context.params as { id: string }
  const response = await fetchAPI<GetTeamApiResponse>(`/teams/${id}`, options)
  return {
    props: {
      initialTeam: response.result.team,
    },
  }
}

const TeamDetailPage: NextPage<{
  initialTeam: Team
}> = ({ initialTeam }) => {
  return <TeamDetail initialTeam={initialTeam} />
}

export default TeamDetailPage

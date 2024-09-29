import type { GetServerSideProps, NextPage } from 'next'
import OpponentRecruitingList from 'src/components/opponent_recruitings/OpponentRecruitingList'
import { GetOpponentRecruitingsApiResponse } from 'src/types/apiResponses'
import { OpponentRecruiting } from 'src/types/opponentRecruiting'
import { Page } from 'src/types/page'
import fetchAPI from 'src/utils/fetchApi'

export const getServerSideProps: GetServerSideProps = async () => {
  const options: RequestInit = {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  }
  const response = await fetchAPI<GetOpponentRecruitingsApiResponse>(
    '/opponent_recruitings/my_team?page=1',
    options,
  )
  return {
    props: {
      initialRecruitings: response.result.opponent_recruitings,
      initialPage: response.result.page,
    },
  }
}

const OpponentRecruitingListPage: NextPage<{
  initialRecruitings: OpponentRecruiting[]
  initialPage: Page
}> = ({ initialRecruitings, initialPage }) => {
  return (
    <OpponentRecruitingList initialRecruitings={initialRecruitings} initialPage={initialPage} isMyTeam={true} />
  )
}

export default OpponentRecruitingListPage

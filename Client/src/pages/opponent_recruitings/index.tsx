import type { GetServerSideProps, NextPage } from 'next'
import { OpponentRecruitingList } from 'src/components/opponent_recruitings/OpponentRecruitingList'
import { getOpponentRecruitingsApiResponse } from 'src/types/apiResponses'
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
  // APIサーバーにトークンを送信
  const response = await fetchAPI<getOpponentRecruitingsApiResponse>(
    '/opponent_recruitings?page=1',
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
    <OpponentRecruitingList initialRecruitings={initialRecruitings} initialPage={initialPage} />
  )
}

export default OpponentRecruitingListPage

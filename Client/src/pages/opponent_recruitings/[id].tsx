import { GetServerSideProps, GetServerSidePropsContext, NextPage } from 'next'
import OpponentRecruitingDetail from 'src/components/opponent_recruitings/OpponentRecruitingDetail'
import { GetOpponentRecruitingApiResponse } from 'src/types/apiResponses'
import { OpponentRecruitingWithComments } from 'src/types/opponentRecruiting'
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
  const response = await fetchAPI<GetOpponentRecruitingApiResponse>(
    `/opponent_recruitings/${id}`,
    options,
  )
  return {
    props: {
      initialRecruitings: response.result.opponent_recruiting,
    },
  }
}

const OpponentRecruitingListPage: NextPage<{
  initialRecruitings: OpponentRecruitingWithComments
}> = ({ initialRecruitings }) => {
  return <OpponentRecruitingDetail opponentRecruitingWithComments={initialRecruitings} />
}

export default OpponentRecruitingListPage

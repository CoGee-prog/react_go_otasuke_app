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
      id: id,
    },
  }
}

const OpponentRecruitingListPage: NextPage<{
  initialRecruitings: OpponentRecruitingWithComments
  id: string
}> = ({ initialRecruitings, id }) => {
  return <OpponentRecruitingDetail opponentRecruitingWithComments={initialRecruitings} id={id} />
}

export default OpponentRecruitingListPage

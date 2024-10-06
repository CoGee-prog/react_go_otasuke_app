import { useRouter } from 'next/router'

export function useNavigateOpponentRecruitingsCreate() {
  const router = useRouter()

  const navigateOpponentRecruitingsCreate = () => router.push('/opponent_recruitings/create')

  return navigateOpponentRecruitingsCreate
}

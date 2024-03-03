import { useRouter } from 'next/router'

export function useNavigateOpponentRecruitingsIndex() {
  const router = useRouter()

  const navigateOpponentRecruitingsIndex = () => router.push('/opponent_recruitings')

  return navigateOpponentRecruitingsIndex
}

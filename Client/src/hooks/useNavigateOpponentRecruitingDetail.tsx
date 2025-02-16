import { useRouter } from 'next/router'

export function useNavigateOpponentRecruitingDetail(id: string) {
  const router = useRouter()

  const navigateOpponentRecruitingDetail = () => router.push(`/opponent_recruitings/${id}`)

  return navigateOpponentRecruitingDetail
}

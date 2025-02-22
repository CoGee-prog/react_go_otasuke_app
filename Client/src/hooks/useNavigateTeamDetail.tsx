import { useRouter } from 'next/router'

export function useNavigateTeamDetail(id: number) {
  const router = useRouter()

  const navigateTeamDetail = () => router.push(`/teams/${id}`)

  return navigateTeamDetail
}

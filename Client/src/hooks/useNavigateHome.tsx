import { useRouter } from 'next/router'

export function useNavigateHome() {
  const router = useRouter()

  const navigateHome = () => router.push('/')

  return navigateHome
}

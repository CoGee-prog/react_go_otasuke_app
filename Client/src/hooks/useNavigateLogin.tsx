import { useRouter } from 'next/router'

export function useNavigateLogin() {
  const router = useRouter()

  const navigateLogin = () => router.push('/login')

  return navigateLogin
}

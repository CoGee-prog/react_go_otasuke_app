import { useRouter } from 'next/router'
import SignInScreen from 'src/components/layouts/SignInScreen'

const Login: React.FC = () => {
  const router = useRouter()
  const { from } = router.query
  return (
    <div>
      <SignInScreen redirectPath={from as string} />
    </div>
  )
}

export default Login

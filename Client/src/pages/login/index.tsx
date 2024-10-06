import { useRouter } from 'next/router'
import { useState, useEffect } from 'react'
import SignInScreen from 'src/components/layouts/SignInScreen'

const Login: React.FC = () => {
  const router = useRouter()

  useEffect(() => {
    // ログインページに遷移した際にリダイレクトパスをクリア
    sessionStorage.removeItem('redirectPath')

    if (router.isReady) {
      const fromPath = router.query.from as string | undefined
      if (fromPath) {
        sessionStorage.setItem('redirectPath', fromPath)
      }
    }
  }, [router.isReady, router.query.from])

  return (
    <div>
      <SignInScreen />
    </div>
  )
}

export default Login

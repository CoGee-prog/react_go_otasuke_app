import React, { useContext } from 'react'
import { AuthContext } from 'src/contexts/AuthContext'
import { useFlashMessage } from 'src/contexts/FlashMessageContext'
import { useNavigateHome } from 'src/hooks/useNavigateHome'
import { LoginApiResponse } from 'src/types/apiResponses'
import { User } from 'src/types/user'
import fetchAPI from 'src/utils/fetchApi'
import { saveDataWithExpiry } from 'src/utils/localStorageHelper'

const LocalLoginButton = () => {
  const { showFlashMessage } = useFlashMessage()
  const navigateHome = useNavigateHome()
  const { setIsLoggedIn, setUser } = useContext(AuthContext)
  const login = (userData: User) => {
    setIsLoggedIn(true)
    setUser(userData)
  }

  // ボタンクリック時に実行される関数
  const handleLogin = async () => {
    try {
      const options: RequestInit = {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      }
      // APIサーバーにトークンを送信
      fetchAPI<LoginApiResponse>('/login', options)
        .then((responseData) => {
          // ユーザー情報をローカルストレージにキャッシュ
          saveDataWithExpiry<User>('user', responseData.result.user, 3600)
          login(responseData.result.user)
          showFlashMessage({ message: responseData.message, type: 'success' })
        })
        .catch((error) => {
          showFlashMessage({
            message:
              error instanceof Error && error.message ? error.message : 'エラーが発生しました',
            type: 'error',
          })
          // ホーム画面に戻す
          navigateHome()
        })
        .finally()
    } catch (error) {
      console.error('Login failed:', error)
    }
  }

  // REACT_APP_FIREBASE_AUTH_DOMAINがlocalhostの場合のみボタンを表示
  if (process.env.REACT_APP_FIREBASE_AUTH_DOMAIN === 'localhost') {
    return <button onClick={handleLogin}>Localhost Login</button>
  }

  // 条件に一致しない場合は何も表示しない
  return null
}

export default LocalLoginButton

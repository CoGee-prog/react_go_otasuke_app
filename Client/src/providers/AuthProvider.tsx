import React, { useState, useEffect } from 'react'
import { AuthContext } from 'src/contexts/AuthContext'
import fetchAPI from 'src/utils/fetchApi'
import { User } from 'src/types/user'
import { onAuthStateChanged } from 'firebase/auth'
import { LoginApiResponse } from 'src/types/apiResponses'
import { useNavigateHome } from 'src/hooks/useNavigateHome'
import { useNavigateLogin } from 'src/hooks/useNavigateLogin'
import { auth } from 'config/firebaseApp'
import { useFlashMessage } from 'src/contexts/FlashMessageContext'
import { loadDataWithExpiry, saveDataWithExpiry } from 'src/utils/localStorageHelper'

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const [user, setUser] = useState<User | null>(null)
  const { showFlashMessage } = useFlashMessage()
  const navigateHome = useNavigateHome()
  const navigateLogin = useNavigateLogin()

  useEffect(() => {
    const unregisterAuthObserver = onAuthStateChanged(auth, (user) => {
      if (user && !isLoggedIn) {
        // キャッシュしたデータがあればそれを返す
        const cachedUser = loadDataWithExpiry<User>('user')
        if (cachedUser) {
          login(cachedUser)
          return
        }

        setIsLoading(true)
        user.getIdToken().then((idToken) => {
          const options: RequestInit = {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              Authorization: `${idToken}`,
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
            .finally(() => setIsLoading(false))
        })
      }
    })

    return () => unregisterAuthObserver()
  }, [isLoggedIn])

  const login = (userData: User) => {
    setIsLoggedIn(true)
    setUser(userData)
  }

  const logout = () => {
    auth.signOut()
    const options: RequestInit = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    }
    fetchAPI('/logout', options).then((responseData) => {
      // レスポンスステータスが200の場合
      if (responseData.status === 200) {
        showFlashMessage({ message: responseData.message, type: 'success' })
        // ホーム画面に戻す
        navigateHome()
      } else {
        showFlashMessage({ message: responseData.message, type: 'error' })
        // 認証エラーの場合
        if (responseData.status === 401) {
          // ログイン画面に戻す
          navigateLogin()
        } else {
          // ホーム画面に戻す
          navigateHome()
        }
      }
      localStorage.removeItem('user')
      setIsLoggedIn(false)
      setUser(null)
    })
  }

  return (
    <AuthContext.Provider value={{ isLoggedIn, setIsLoggedIn ,isLoading, login, logout, user, setUser }}>
      {children}
    </AuthContext.Provider>
  )
}

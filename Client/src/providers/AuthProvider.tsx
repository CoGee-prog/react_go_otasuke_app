import React, { useState, useEffect } from 'react'
import { AuthContext } from 'src/contexts/AuthContext'
import fetchAPI from 'src/utils/fetchApi'
import { User } from 'src/types/user'
import { getAuth, onAuthStateChanged } from 'firebase/auth'
import { loginApiResponse } from 'src/types/apiResponses'
import { useNavigateHome } from 'src/hooks/useNavigateHome'
import firebaseApp from 'config/firebaseApp'
import { useFlashMessage } from 'src/contexts/FlashMessageContext'

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const firebaseAuth = getAuth(firebaseApp)
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const [user, setUser] = useState<User | null>(null)
  const { showFlashMessage } = useFlashMessage()
  const navigateHome = useNavigateHome()

  useEffect(() => {
    const unregisterAuthObserver = onAuthStateChanged(firebaseAuth, (user) => {
      if (user && !isLoggedIn) {
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
          fetchAPI<loginApiResponse>('/login', options)
            .then((responseData) => {
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
    firebaseAuth.signOut()
    const options: RequestInit = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    }
    fetchAPI('/logout', options).then((responseData) => {
      showFlashMessage({ message: responseData.message, type: 'success' })
      setIsLoggedIn(false)
      setUser(null)
      // ホーム画面に戻す
      navigateHome()
    })
  }

  return (
    <AuthContext.Provider value={{ isLoggedIn, isLoading, login, logout, user }}>
      {children}
    </AuthContext.Provider>
  )
}

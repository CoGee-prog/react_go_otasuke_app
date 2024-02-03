import React, { useState, useEffect } from 'react'
import firebaseConfig from 'config/firebaseConfig'
import { AuthContext } from 'src/contexts/AuthContext'
import fetchAPI from 'src/helpers/apiService'
import { User } from 'src/types/user'
import { getAuth, onAuthStateChanged } from 'firebase/auth'
import { ResponseData } from 'src/types/responseData'
import { loginApiResponse } from 'src/types/apiResponses'

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const firebaseAuth = getAuth(firebaseConfig)
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const [user, setUser] = useState<User | null>(null)

  useEffect(() => {
    const unregisterAuthObserver = onAuthStateChanged(firebaseAuth, (user) => {
      if (user && !isLoggedIn) {
        setIsLoading(true);
        user.getIdToken().then((idToken) => {
          // APIサーバーにトークンを送信
          fetchAPI('/login', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              // トークンをAuthorizationヘッダーに含める
              Authorization: `${idToken}`,
            },
            credentials: 'include',
          })
            .then((data: ResponseData<loginApiResponse>) => {
							login(data.result.user)
              setIsLoading(false)
            })
            .catch((error) => {
              setIsLoading(false)
              console.error('Error:', error)
            })
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
    fetchAPI('/logout', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    }).then(() => {
      setIsLoggedIn(false)
      setUser(null)
    })
  }

  return (
    <AuthContext.Provider value={{ isLoggedIn, isLoading, login, logout, user }}>
      {children}
    </AuthContext.Provider>
  )
}

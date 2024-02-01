/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/anchor-is-valid */
// Import FirebaseAuth and firebase.
import React, { useContext, useEffect, useState } from 'react'
import firebaseConfig from 'config/firebaseConfig'
import { Card, CardContent } from '@mui/material'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'
import {
  getAuth,
  onAuthStateChanged,
  EmailAuthProvider,
  GoogleAuthProvider,
  FacebookAuthProvider,
} from 'firebase/auth'
import fetchAPI from 'src/helpers/apiService'
import { AuthContext } from 'src/contexts/AuthContext'
import { ResponseData } from 'src/types/responseData'
import { loginApiResponse } from 'src/types/apiResponses'

const firebaseAuth = getAuth(firebaseConfig)

function SignInScreen() {
  const { login, isLoggedIn } = useContext(AuthContext)

  useEffect(() => {
		// ログイン時のみ動くようにする
    const unregisterAuthObserver = onAuthStateChanged(firebaseAuth, (user) => {
      if (user) {
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
            })
            .catch((error) => {
              console.error('Error:', error)
            })
        })
      }
    })

    return () => unregisterAuthObserver()
  }, [])

  const uiConfig = {
    signInFlow: 'popup',
    signInOptions: [
      GoogleAuthProvider.PROVIDER_ID,
      FacebookAuthProvider.PROVIDER_ID,
      EmailAuthProvider.PROVIDER_ID,
    ],
    callbacks: {
      signInSuccessWithAuthResult: () => false,
    },
  }

  return (
    <div>
      {isLoggedIn ? (
        <>
          <button onClick={() => firebaseAuth.signOut()}>Sign out</button>
        </>
      ) : (
        <Card sx={{ minWidth: 275 }}>
          <CardContent>
            <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebaseAuth} />
          </CardContent>
        </Card>
      )}
    </div>
  )
}

export default SignInScreen

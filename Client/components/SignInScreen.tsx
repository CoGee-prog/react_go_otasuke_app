/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/anchor-is-valid */
// Import FirebaseAuth and firebase.
import React, { useEffect, useState } from 'react'
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
import fetchAPI from 'helpers/apiService'

const auth = getAuth(firebaseConfig)

function SignInScreen() {
  const [isSignedIn, setIsSignedIn] = useState<boolean>(false);
	const [serverMessage, setServerMessage] = useState<string>('')

  useEffect(() => {
    const unregisterAuthObserver = onAuthStateChanged(auth, user => {
      
      if (user) {
        user.getIdToken().then(idToken => {
          // APIサーバーにトークンを送信
          fetchAPI('/login', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              // トークンをAuthorizationヘッダーに含める
              Authorization: `${idToken}`,
            },
          })
            .then((data) => {
							// サインイン状態にする
							setIsSignedIn(true);
              // サーバーからのレスポンスを状態に保存
              setServerMessage(data.message)
            })
            .catch((error) => {
              console.error('Error:', error)
            })
        });
      }
    });

    return () => unregisterAuthObserver();
  }, []);

  const uiConfig = {
    signInFlow: 'popup',
    signInOptions: [
      GoogleAuthProvider.PROVIDER_ID,
      FacebookAuthProvider.PROVIDER_ID,
      EmailAuthProvider.PROVIDER_ID
    ],
    callbacks: {
      signInSuccessWithAuthResult: () => false
    }
  };

  return (
    <div>
      {isSignedIn ? (
          <>
            <p>{serverMessage}</p>
            <button onClick={() => auth.signOut()}>Sign out</button>
          </>
        ) : (
          <Card sx={{ minWidth: 275 }}>
            <CardContent>
              <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={auth} />
            </CardContent>
          </Card>
        )
      }
    </div>
  )
};

export default SignInScreen

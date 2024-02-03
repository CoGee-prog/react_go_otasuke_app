/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/anchor-is-valid */
// Import FirebaseAuth and firebase.
import React from 'react'
import firebaseConfig from 'config/firebaseConfig'
import { Card, CardContent } from '@mui/material'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'
import {
  getAuth,
  EmailAuthProvider,
  GoogleAuthProvider,
  FacebookAuthProvider,
} from 'firebase/auth'

const firebaseAuth = getAuth(firebaseConfig)

function SignInScreen() {
  const uiConfig = {
    signInFlow: 'redirect',
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
			<Card sx={{ minWidth: 275 }}>
				<CardContent>
					<StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebaseAuth} />
				</CardContent>
			</Card>
    </div>
  )
}

export default SignInScreen

/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/anchor-is-valid */
// Import FirebaseAuth and firebase.
import { auth } from 'config/firebaseApp'
import { Box, Card, CardContent, Typography, useTheme } from '@mui/material'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'
import { EmailAuthProvider, GoogleAuthProvider, FacebookAuthProvider } from 'firebase/auth'
import { useRouter } from 'next/router'
import React, { useContext, useEffect } from 'react'
import { AuthContext } from 'src/contexts/AuthContext'
import LoadingScreen from '../commons/LoadingScreen'
import LocalLoginButton from '../commons/LocalLoginButton'

const SignInScreen: React.FC = () => {
  const theme = useTheme()
  const router = useRouter()
  const { isLoggedIn, isLoading } = useContext(AuthContext)

  useEffect(() => {
    // ログイン状態であればリダイレクト
    if (isLoggedIn) {
        const redirectPath = sessionStorage.getItem('redirectPath')
        router.push(redirectPath || '/opponent_recruitings')
        sessionStorage.removeItem('redirectPath')
    }
  }, [isLoggedIn, router])

  const uiConfig = {
    signInFlow: 'redirect',
    signInOptions: [
      GoogleAuthProvider.PROVIDER_ID,
      FacebookAuthProvider.PROVIDER_ID,
      EmailAuthProvider.PROVIDER_ID,
    ],
    callbacks: {
      signInSuccessWithAuthResult: () => {
        return false
      },
    },
  }

  return (
    <Box
      sx={{
        maxWidth: 400,
        mx: 'auto',
        my: 4,
        px: 2,
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        gap: 2,
      }}
    >
      {isLoading ? (
        <LoadingScreen />
      ) : (
        <Card
          sx={{
            width: '100%',
            boxShadow: theme.shadows[3],
            borderRadius: theme.shape.borderRadius,
            overflow: 'hidden',
            bgcolor: 'background.paper',
          }}
        >
          <CardContent
            sx={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
              p: 3,
            }}
          >
            <Typography
              variant='h5'
              component='h2'
              gutterBottom
              sx={{
                mb: 2,
                fontWeight: 'bold',
              }}
            >
              ログイン
            </Typography>
            <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={auth} />
						<LocalLoginButton />
          </CardContent>
        </Card>
      )}
    </Box>
  )
}

export default SignInScreen

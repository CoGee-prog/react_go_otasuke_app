/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/anchor-is-valid */
// Import FirebaseAuth and firebase.
import { auth } from 'config/firebaseApp'
import { Box, Card, CardContent, Typography, useTheme } from '@mui/material'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'
import { getAuth, EmailAuthProvider, GoogleAuthProvider, FacebookAuthProvider } from 'firebase/auth'
import { useRouter } from 'next/router'
import React from 'react'

interface SignInScreenProps {
  redirectPath?: string
}

const SignInScreen: React.FC<SignInScreenProps> = ({ redirectPath }) => {
  const theme = useTheme()
  const router = useRouter()

  const uiConfig = {
    signInFlow: 'redirect',
    signInOptions: [
      GoogleAuthProvider.PROVIDER_ID,
      FacebookAuthProvider.PROVIDER_ID,
      EmailAuthProvider.PROVIDER_ID,
    ],
    callbacks: {
      signInSuccessWithAuthResult: () => {
        console.log(redirectPath)
        router.push(redirectPath || '/opponent_recruitings')
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
        </CardContent>
      </Card>
    </Box>
  )
}

export default SignInScreen

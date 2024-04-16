/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/anchor-is-valid */
// Import FirebaseAuth and firebase.
import firebaseApp from 'config/firebaseApp'
import { Box, Card, CardContent, Typography, useTheme } from '@mui/material'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'
import { getAuth, EmailAuthProvider, GoogleAuthProvider, FacebookAuthProvider } from 'firebase/auth'

const firebaseAuth = getAuth(firebaseApp)

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

function SignInScreen() {
  const theme = useTheme()

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
          <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={firebaseAuth} />
        </CardContent>
      </Card>
    </Box>
  )
}

export default SignInScreen

import AppBar from '@mui/material/AppBar'
import Box from '@mui/material/Box'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import React, { useContext } from 'react'
import SignInBackdrop from './SignInBackdrop'
import { AuthContext } from 'src/contexts/AuthContext'
import AccountMenu from './AccountMenu'
import LoadingScreen from './LoadingScreen'

export default function Header() {
  const {isLoggedIn, isLoading} = useContext(AuthContext)
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position='static'>
        <Toolbar>
          <Typography variant='h6' component='div' sx={{ flexGrow: 1 }}>
            おたスケ
          </Typography>
          <div>
            {isLoading ? <LoadingScreen/>: isLoggedIn ? <AccountMenu /> : <SignInBackdrop />}
          </div>
        </Toolbar>
      </AppBar>
    </Box>
  )
}

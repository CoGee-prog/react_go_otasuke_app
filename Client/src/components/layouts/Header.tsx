import AppBar from '@mui/material/AppBar'
import Box from '@mui/material/Box'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import React, { useContext } from 'react'
import SignInBackdrop from './SignInBackdrop'
import { AuthContext } from 'src/contexts/AuthContext'
import AccountMenu from './AccountMenu'
import LoadingScreen from '../commons/LoadingScreen'
import { useNavigateHome } from 'src/hooks/useNavigateHome'

export default function Header() {
  const { isLoggedIn, isLoading } = useContext(AuthContext)
  const navigateHome = useNavigateHome()

  const appBarColor = '#333'

	return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position='static' style={{ backgroundColor: appBarColor }}>
        <Toolbar style={{ position: 'relative' }}>
          <Typography
            variant='h6'
            component='div'
            style={{ cursor: 'pointer', flexGrow: 1 }}
            onClick={navigateHome}
          >
            おたスケ
          </Typography>
          <div>
            {isLoading ? <LoadingScreen /> : isLoggedIn ? <AccountMenu /> : <SignInBackdrop />}
          </div>
        </Toolbar>
      </AppBar>
    </Box>
  )
}

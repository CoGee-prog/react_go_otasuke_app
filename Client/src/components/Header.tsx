import MenuIcon from '@mui/icons-material/Menu'
import AppBar from '@mui/material/AppBar'
import Box from '@mui/material/Box'
import IconButton from '@mui/material/IconButton'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import React, { useContext } from 'react'
import SignInBackdrop from './SignInBackdrop'
import { AuthContext } from 'src/contexts/AuthContext'

export default function Header() {
  const authContext = useContext(AuthContext)
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position='static'>
        <Toolbar>
          <Typography variant='h6' component='div' sx={{ flexGrow: 1 }}>
            おたスケ
          </Typography>
          <div>
            {authContext.isLoggedIn ? (
              <div style={{ display: 'flex', alignItems: 'center' }}>
                <div>{authContext.user?.name}</div>
                <div>{authContext.user?.current_team_name}</div>
                <IconButton size='large' edge='start' color='inherit' aria-label='menu' sx={{ mr: 2 }}>
                  <MenuIcon />
                </IconButton>
              </div>
            ) : (
              <SignInBackdrop />
            )}
          </div>
        </Toolbar>
      </AppBar>
    </Box>
  )
}

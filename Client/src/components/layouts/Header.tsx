import AppBar from '@mui/material/AppBar'
import Box from '@mui/material/Box'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import React, { useContext } from 'react'
import SignInBackdrop from './SignInBackdrop'
import { AuthContext } from 'src/contexts/AuthContext'
import AccountMenu from './AccountMenu'
import LoadingScreen from '../commons/LoadingScreen'
import { FlashMessageContext } from 'src/contexts/FlashMessageContext'
import { useNavigateHome } from 'src/hooks/useNavigateHome'

export default function Header() {
  const { flashMessage } = useContext(FlashMessageContext)
  const { isLoggedIn, isLoading } = useContext(AuthContext)
  const navigateHome = useNavigateHome()

  const appBarColor = '#333' // ダークグレー
  const successColor = '#4CAF50' // 明るいグリーン
  const errorColor = '#F44336' // 明るいレッド
  const flashMessageBackgroundColor =
    flashMessage.type === 'success'
      ? successColor
      : flashMessage.type === 'error'
      ? errorColor
      : 'transparent'

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
          {flashMessage.message && (
            <div
              style={{
                backgroundColor: flashMessageBackgroundColor,
                color: 'white',
                padding: '10px',
                borderRadius: '5px',
                position: 'absolute',
                left: '50%',
                transform: 'translateX(-50%)',
                boxSizing: 'border-box',
              }}
            >
              {flashMessage.message}
            </div>
          )}
          <div>
            {isLoading ? <LoadingScreen /> : isLoggedIn ? <AccountMenu /> : <SignInBackdrop />}
          </div>
        </Toolbar>
      </AppBar>
    </Box>
  )
}

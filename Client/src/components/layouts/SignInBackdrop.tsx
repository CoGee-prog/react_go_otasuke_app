import Backdrop from '@mui/material/Backdrop'
import Button from '@mui/material/Button'
import * as React from 'react'
import SignInScreen from './SignInScreen'

export default function SignInBackdrop() {
  const [open, setOpen] = React.useState(false)
  const handleClose = () => {
    setOpen(false)
  }
  const handleToggle = () => {
    setOpen(!open)
  }

  return (
    <div>
      <Button color='inherit' onClick={handleToggle}>
        ログイン
      </Button>
      <Backdrop
        sx={{ color: '#fff', zIndex: (theme) => theme.zIndex.drawer + 1 }}
        open={open}
        onClick={handleClose}
      >
        <SignInScreen />
      </Backdrop>
    </div>
  )
}

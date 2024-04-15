import Backdrop from '@mui/material/Backdrop'
import Button from '@mui/material/Button'
import * as React from 'react'
import SignInScreen from './SignInScreen'
import Link from 'next/link'

export default function SignInBackdrop() {
  return (
    <div>
      <Link href='/login' passHref>
        <Button color='inherit'>
          ログイン
        </Button>
      </Link>
    </div>
  )
}

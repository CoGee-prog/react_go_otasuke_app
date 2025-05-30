import Button from '@mui/material/Button'
import * as React from 'react'
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

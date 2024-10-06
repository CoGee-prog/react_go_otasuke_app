import React from 'react'
import { Button, ButtonProps } from '@mui/material'

const PrimaryButton: React.FC<ButtonProps> = (props) => {
  return (
    <Button
      type='submit'
      variant='contained'
      sx={{
        backgroundColor: '#009688',
        color: '#fff',
        padding: '10px 10px',
        fontSize: '16px',
        fontWeight: 'bold',
        boxShadow: '0 3px 5px rgba(0, 0, 0, 0.2)',
        '&:hover': {
          backgroundColor: '#00796B',
          boxShadow: '0 5px 7px rgba(0, 0, 0, 0.3)',
        },
      }}
      fullWidth
      {...props}
    >
      {props.children}
    </Button>
  )
}

export default PrimaryButton

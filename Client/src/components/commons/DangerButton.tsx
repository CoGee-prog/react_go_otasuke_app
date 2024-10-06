import React from 'react';
import { Button, ButtonProps } from '@mui/material';

const DangerButton: React.FC<ButtonProps> = (props) => {
  return (
    <Button
      variant='contained'
      sx={{
        backgroundColor: '#e53935',
        color: '#fff',
        padding: '10px 10px',
        fontSize: '16px',
        fontWeight: 'bold',
        boxShadow: '0 3px 5px rgba(0, 0, 0, 0.2)',
        '&:hover': {
          backgroundColor: '#d32f2f',
          boxShadow: '0 5px 7px rgba(0, 0, 0, 0.3)',
        },
      }}
      fullWidth
      {...props}
    >
      {props.children}
    </Button>
  )
};

export default DangerButton;

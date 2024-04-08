import React from 'react'
import { Box, Typography } from '@mui/material'
import { useFlashMessage } from 'src/contexts/FlashMessageContext'

const FlashMessage = () => {
  const { flashMessage } = useFlashMessage()

  if (!flashMessage.message) return null

  const successColor = '#4CAF50' // 明るいグリーン
  const errorColor = '#F44336' // 明るいレッド
  const flashMessageBackgroundColor =
    flashMessage.type === 'success'
      ? successColor
      : flashMessage.type === 'error'
      ? errorColor
      : 'transparent'

  return (
    <Box
      sx={{
        position: 'fixed',
        top: 5,
        left: '50%',
        transform: 'translateX(-50%)',
        bgcolor: 'primary.main',
        color: 'white',
        p: 2,
        zIndex: 1000,
        textAlign: 'center',
        backgroundColor: flashMessageBackgroundColor,
        padding: '10px',
        borderRadius: '5px',
        boxSizing: 'border-box',
      }}
    >
      <Typography variant='subtitle1'>{flashMessage.message}</Typography>
    </Box>
  )
}

export default FlashMessage

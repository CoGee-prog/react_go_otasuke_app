import React, { createContext, useContext } from 'react'

export interface FlashMessage {
  message: string
  type: 'success' | 'error'
}

interface FlashMessageContextType {
  flashMessage: FlashMessage
  setFlashMessage: (message: FlashMessage) => void
  clearFlashMessage: () => void
}

const defaultFlashMessageContext: FlashMessageContextType = {
  flashMessage: { message: '', type: 'success' },
  setFlashMessage: (message: FlashMessage) => {},
  clearFlashMessage: () => {},
}

export const FlashMessageContext = createContext<FlashMessageContextType>(
  defaultFlashMessageContext,
)

export const useFlashMessage = () => {
  const context = useContext(FlashMessageContext)
  if (!context) {
    throw new Error('useFlashMessage must be used within a FlashMessageProvider')
  }
  return context
}

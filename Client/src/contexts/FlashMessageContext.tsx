import { createContext, useContext } from 'react'

export interface FlashMessage {
  message: string
  type: 'success' | 'error' | null
}

interface FlashMessageContextType {
  flashMessage: FlashMessage
  showFlashMessage: (flashMessage: FlashMessage) => void
}

const defaultFlashMessageContext: FlashMessageContextType = {
  flashMessage: { message: '', type: 'success' },
  showFlashMessage: (flashMessage: FlashMessage) => {},
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

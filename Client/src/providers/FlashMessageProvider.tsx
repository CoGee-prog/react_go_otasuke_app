import { useState } from "react"
import { FlashMessage, FlashMessageContext } from "src/contexts/FlashMessageContext"

export const FlashMessageProvider: React.FC = ({ children }) => {
  const [flashMessage, setFlashMessage] = useState<FlashMessage>({message: '', type: 'success'})

  const clearFlashMessage = () => setFlashMessage({message: '', type: 'success'})

  return (
    <FlashMessageContext.Provider value={{ flashMessage, setFlashMessage, clearFlashMessage }}>
      {children}
    </FlashMessageContext.Provider>
  )
}

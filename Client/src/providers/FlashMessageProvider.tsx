import { useState } from "react"
import { FlashMessage, FlashMessageContext } from "src/contexts/FlashMessageContext"

export const FlashMessageProvider: React.FC = ({ children }) => {
  const [flashMessage, setFlashMessage] = useState<FlashMessage>({message: '', type: 'success'})

	  const showFlashMessage = (flashMessage: FlashMessage) => {
      setFlashMessage(flashMessage)
      // 数秒後にメッセージを消去
      const timer = setTimeout(() => setFlashMessage({ message: '', type: 'success' }), 5000)

      // クリーンアップ関数でタイマーをクリア
      return () => clearTimeout(timer)
    }

  return (
    <FlashMessageContext.Provider value={{ flashMessage, showFlashMessage }}>
      {children}
    </FlashMessageContext.Provider>
  )
}

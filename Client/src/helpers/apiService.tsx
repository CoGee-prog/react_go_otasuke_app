import { error } from "console"
import { useContext, useEffect } from "react"
import { FlashMessage, FlashMessageContext, useFlashMessage } from "src/contexts/FlashMessageContext"
import { ResponseData } from "src/types/responseData"

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || ''

async function fetchAPI<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const {setFlashMessage} = useFlashMessage()
  const url = `${BASE_URL}${endpoint}`
  const response = await fetch(url, options)

  // フラッシュメッセージを設定し、一定時間後にクリアする関数
  const showMessageWithTimeout = (
    message: string,
    type: 'success' | 'error',
    timeout: number = 5000,
  ) => {
    setFlashMessage({ message, type })

    // 一定時間後にメッセージをクリア
    setTimeout(() => setFlashMessage({ message: '', type: 'success' }), timeout)
  }
	
	const responseData: ResponseData<T> = await response.json()
  if (!response.ok) {
    // エラーハンドリングをここで行う
    showMessageWithTimeout(responseData.message || 'エラーが発生しました', 'error')
		throw new Error('API call failed: ' + response.statusText)
  }
	showMessageWithTimeout(responseData.message, 'success')
  return responseData.result
}



export default fetchAPI

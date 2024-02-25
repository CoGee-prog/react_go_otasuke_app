import { useState } from 'react'
import { useFlashMessage } from 'src/contexts/FlashMessageContext'
import fetchAPI from 'src/utils/fetchApi'

function useApiWithFlashMessage<T>() {
  const [data, setData] = useState<T | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)
  const { showFlashMessage } = useFlashMessage()

  const request = async (endpoint: string, options: RequestInit = {}) => {
    setIsLoading(true)
    try {
      const responseData = await fetchAPI<T>(endpoint, options)
      setData(responseData.result)
      showFlashMessage({ message: responseData.message, type: 'success' })
    } catch (error) {
			setError(error instanceof Error ? error.message : 'エラーが発生しました')
      showFlashMessage({
        message: error instanceof Error && error.message ? error.message : 'エラーが発生しました',
        type: 'error',
      })
    } finally {
      setIsLoading(false)
    }
  }

  return { data, isLoading, error, request }
}

export default useApiWithFlashMessage

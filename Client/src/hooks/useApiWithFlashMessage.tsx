import { useState } from 'react'
import { useFlashMessage } from 'src/contexts/FlashMessageContext'
import fetchAPI from 'src/utils/fetchApi'

function useApiWithFlashMessage<T>() {
  const [data, setData] = useState<T | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [errorMessage, setErrorMessage] = useState<string | null>(null)
  const { showFlashMessage } = useFlashMessage()

  const request = async (endpoint: string, options: RequestInit = {}) => {
    setIsLoading(true)
    try {
      const responseData = await fetchAPI<T>(endpoint, options)
      setData(responseData.result)
			// HTTPステータスが200以外はエラー
			if(responseData.status !== 200){
				throw new Error(responseData.message)
			}
			showFlashMessage({ message: responseData.message, type: 'success' })
    } catch (error) {
			setErrorMessage(error instanceof Error ? error.message : 'エラーが発生しました')
      showFlashMessage({
        message: error instanceof Error && error.message ? error.message : 'エラーが発生しました',
        type: 'error',
      })
			throw new Error(errorMessage!)
    } finally {
      setIsLoading(false)
    }
  }

  return { data, isLoading, error: errorMessage, request }
}

export default useApiWithFlashMessage

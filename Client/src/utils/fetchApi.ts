import { ResponseData } from 'src/types/responseData'

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || ''

async function fetchAPI<T>(endpoint: string, options: RequestInit = {}): Promise<ResponseData<T>> {
  const url = `${BASE_URL}${endpoint}`
  const response = await fetch(url, options)

  const responseData: ResponseData<T> = await response.json()
  if (!response.ok) {
    throw new Error(responseData.message)
  }
  return responseData
}

export default fetchAPI

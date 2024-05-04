import { ResponseData } from 'src/types/responseData'

const BASE_URL_CLIENT = process.env.NEXT_PUBLIC_API_BASE_URL_CLIENT || ''
const BASE_URL_SERVER = process.env.NEXT_PUBLIC_API_BASE_URL_SERVER || ''

async function fetchAPI<T>(endpoint: string, options: RequestInit = {}): Promise<ResponseData<T>> {
	// 実行環境に応じてベースURLを選択
  const BASE_URL = typeof window === 'undefined' ? BASE_URL_SERVER : BASE_URL_CLIENT;
  const url = new URL(endpoint, BASE_URL).toString();
  const response = await fetch(url, options)

  const responseData: ResponseData<T> = await response.json()
  return responseData
}

export default fetchAPI

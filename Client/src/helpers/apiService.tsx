const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || ''

async function fetchAPI(endpoint: string, options: RequestInit = {}) {
  const url = `${BASE_URL}${endpoint}`
  const response = await fetch(url, options)
  if (!response.ok) {
    // エラーハンドリングをここで行う
    throw new Error('API call failed: ' + response.statusText)
  }
  return response.json()
}

export default fetchAPI

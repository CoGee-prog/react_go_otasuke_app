import { ResponseData } from 'src/types/responseData'
import { loadDataWithExpiry } from './localStorageHelper';
import { User } from 'src/types/user';

const BASE_URL_CLIENT = process.env.NEXT_PUBLIC_API_BASE_URL_CLIENT || ''
const BASE_URL_SERVER = process.env.NEXT_PUBLIC_API_BASE_URL_SERVER || ''

async function fetchAPI<T>(endpoint: string, options: RequestInit = {}): Promise<ResponseData<T>> {
	// 実行環境に応じてベースURLを選択
  const BASE_URL = typeof window === 'undefined' ? BASE_URL_SERVER : BASE_URL_CLIENT;
	// ローカル環境の場合
	if (process.env.REACT_APP_FIREBASE_AUTH_DOMAIN === 'localhost') {
		var cachedUser = loadDataWithExpiry<User>('user')
		var userId: string;
		if (cachedUser){
			userId = cachedUser.id as string
		}else{
			userId = 'test-user-id'
		}

		// ヘッダーが未定義なら新たに初期化する
		if (!options.headers) {
				options.headers = {};
		}
		const headers = options.headers as Record<string, string>;
		headers['X-User-Id'] = userId;
	}
  const url = new URL(endpoint, BASE_URL).toString();
  const response = await fetch(url, options)

  const responseData: ResponseData<T> = await response.json()
  return responseData
}

export default fetchAPI

import { createContext } from 'react'
import { User } from 'src/types/user'

interface AuthContextType {
  isLoggedIn: boolean
  setIsLoggedIn: (boolean: boolean) => void
  isLoading: boolean
  user: User | null
  setUser: (user: User | null) => void
  login: (userData: User) => void
  logout: () => void
}

const defaultAuthContext: AuthContextType = {
  isLoggedIn: false,
  setIsLoggedIn: () => {},
  isLoading: false,
  user: null,
  setUser: () => {},
  login: () => {},
  logout: () => {},
}

export const AuthContext = createContext<AuthContextType>(defaultAuthContext)

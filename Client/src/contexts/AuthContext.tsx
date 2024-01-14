import { createContext } from 'react'
import { User } from 'src/types/user'

interface AuthContextType {
  isLoggedIn: boolean
  user: User | null
  login: (userData: User) => void
  logout: () => void
}

const defaultAuthContext: AuthContextType = {
  isLoggedIn: false,
  user: null,
  login: () => {},
  logout: () => {},
}

export const AuthContext = createContext<AuthContextType>(defaultAuthContext)

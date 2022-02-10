import {
  UserWithAuth,
  useSignInWithEmailMutation,
  useSignInWithUsernameMutation,
  useSignUpMutation,
} from '@/generated/graphql'
import { useLocalStorage } from '@/hooks'
import { getItem } from '@/utils/storage'
import { useRouter } from 'next/router'
import React, { useCallback, useEffect, useState } from 'react'
import { useContext } from 'react'
import { z } from 'zod'
import { GITHUB_AUTH_CALLBACK_URL, GITHUB_CLIENT_ID } from './constants'

export const SignUpCredentials = z.object({
  email: z.string().email(),
  username: z.string().min(3).max(50),
  password: z.string().min(8).max(24),
})
export type SignUpCredentialsType = z.infer<typeof SignUpCredentials>

type SignInOptions =
  | {
      type: 'github'
    }
  | {
      type: 'username'
      username: string
      password: string
    }
  | {
      type: 'email'
      email: string
      password: string
    }

type IAuthContext = {
  loading: boolean
  signUp: (credentials: SignUpCredentialsType) => Promise<Error | null>
  signIn: (options: SignInOptions) => Promise<Error | null>
  signOut: () => void
  user?: UserWithAuth
}

const AuthContext = React.createContext<IAuthContext>({
  loading: false,
  signUp: async () => null,
  signIn: async () => null,
  signOut: () => {},
})

export const AuthProvider: React.FC = ({ children }) => {
  const { push } = useRouter()
  const [loading, setLoading] = useState<boolean>(true)
  const [user, setUser] = useLocalStorage<UserWithAuth | undefined>(
    'slides-token',
    undefined
  )
  const [_resultEmail, _signInWithEmail] = useSignInWithEmailMutation()
  const [_resultUsername, _signInWithUsername] = useSignInWithUsernameMutation()
  const [_resultSignUp, _signUp] = useSignUpMutation()

  // init state from localStorage
  useEffect(() => {
    const { data } = getItem<UserWithAuth | undefined>('token')
    if (data) setUser(data)
    setLoading(false)
  }, [])

  const signUp: IAuthContext['signUp'] = useCallback(
    async (credentials) => {
      setLoading(true)
      const { data, error } = await _signUp({ input: credentials })
      if (error || !data) {
        setLoading(false)
        return Error('email or username is already in use')
      }
      setUser(data.SignUp)
      setLoading(false)
      push('/files')
      return null
    },
    [setLoading, _signUp, setUser]
  )

  const signIn: IAuthContext['signIn'] = useCallback(async (options) => {
    setLoading(true)
    switch (options.type) {
      case 'github':
        window.location.href = `https://github.com/login/oauth/authorize?client_id=${GITHUB_CLIENT_ID}&redirect_uri=${GITHUB_AUTH_CALLBACK_URL}`
        return null
      case 'username': {
        const { username, password } = options
        const { data, error } = await _signInWithUsername({
          input: { username, password },
        })
        if (error || !data) {
          setLoading(false)
          return Error('username or password is incorrect')
        }
        setUser(data.SignInWithUsername)
        setLoading(false)
        push('/files')
        return null
      }
      case 'email': {
        const { email, password } = options
        const { data, error } = await _signInWithEmail({
          input: { email, password },
        })
        if (error || !data) {
          setLoading(false)
          return Error('email or password is incorrect')
        }
        setUser(data.SignInWithEmail)
        setLoading(false)
        push('/files')
        return null
      }
    }
  }, [])

  const signOut: IAuthContext['signOut'] = () => {
    setLoading(true)
    setUser(undefined)
    setLoading(false)
  }

  return (
    <AuthContext.Provider value={{ signUp, signIn, signOut, loading, user }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => useContext(AuthContext)

export const CheckAuth: React.FC = ({ children }) => {
  const { loading, user } = useAuth()
  const { replace } = useRouter()

  useEffect(() => {
    if (!loading && !user) replace('/')
  }, [loading, user])

  return <>{children}</>
}

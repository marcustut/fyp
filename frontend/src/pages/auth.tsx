import React, { useEffect, useState } from 'react'
import { Spinner } from '@/components/Spinner'
import { regex } from '@/utils/regex'
import { Icon } from '@iconify/react'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { useLocalStorage } from '@/hooks'
import { UserWithAuth, useUserByAccessTokenQuery } from '@/generated/graphql'
import { z } from 'zod'
import { useAuth } from '@/lib/auth'

type AuthPageView = 'loading' | 'failed'

const AuthPageQuery = z.object({
  access_token: z.string().regex(regex.jwt, 'not a valid JWT token'),
})

const Auth: NextPage = () => {
  const { query, replace } = useRouter()
  const { user } = useAuth()
  const [result] = useUserByAccessTokenQuery({
    variables: { token: query.access_token as string },
  })
  const [view, setView] = useState<AuthPageView>('loading')
  const [_, setUser] = useLocalStorage<UserWithAuth | undefined>(
    'slides-token',
    undefined
  )

  const handleError = (err: unknown) => {
    console.error(err)
    setView('failed')
  }

  // validate access_token
  useEffect(() => {
    // skip when query string is undefined
    if (!query.access_token) return

    // check if access_token exists and is valid JWT
    const res = AuthPageQuery.safeParse(query)
    if (!res.success) handleError(res.error)

    // skip when fetching
    if (result.fetching) return

    // if no user, redirect back to landing
    if (!result.data) {
      replace('/')
      return
    }

    // else bring to app
    setUser(result.data.UserByAccessToken)
    replace('/files')
  }, [query, result])

  return (
    <div className="flex h-[100vh] flex-col items-center justify-center">
      {view === 'loading' ? (
        <>
          <Spinner className="h-8 w-8 text-indigo-500" />
          <p className="mt-2 text-sm text-indigo-500">signing in</p>
        </>
      ) : (
        <>
          <Icon icon="mdi:alert-circle" className="h-8 w-8 text-red-500" />
          <p className="mt-2 text-sm text-red-500">failed sign in</p>
        </>
      )}
    </div>
  )
}

export default Auth

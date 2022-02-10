import '@/styles/globals.css'
import 'react-toastify/dist/ReactToastify.css'
import React from 'react'
import NextNProgress from 'nextjs-progressbar'
import type { AppProps } from 'next/app'
import { withUrqlClient } from 'next-urql'
import { AuthProvider } from '@/lib/auth'
import { AUTH_API_URL } from '@/lib/constants'
import { ToastContainer } from 'react-toastify'

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <NextNProgress color="#6366f1" options={{ showSpinner: false }} />
      <AuthProvider>
        <Component {...pageProps} />
        <ToastContainer />
      </AuthProvider>
    </>
  )
}

export default withUrqlClient((_ssrExchange, ctx) => ({
  url: AUTH_API_URL,
}))(MyApp)

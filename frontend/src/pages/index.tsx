import Head from 'next/head'
import { withUrqlClient } from 'next-urql'
import { useDarkMode } from '@/hooks'
import { AUTH_API_URL } from '@/lib/constants'
import { gql, useQuery } from 'urql'
import { useState } from 'react'
import { AuthDialog } from '@/features/auth'

const VALIDATE_ACCESS_TOKEN_QUERY = gql`
  query ValidateAccessToken($access_token: String!) {
    ValidateAccessToken(token: $access_token)
  }
`

const IndexPage = () => {
  const [authOpen, setAuthOpen] = useState<boolean>(false)
  const { darkMode, setDarkMode } = useDarkMode()
  // const [summarizedText, setSummarizedText] = useState<string>('not summarized')
  const [result] = useQuery({
    query: VALIDATE_ACCESS_TOKEN_QUERY,
    variables: { access_token: 'hahdahshsahdh' },
  })

  const handleSignIn = () => {
    setAuthOpen(true)
  }

  return (
    <>
      <Head>
        <title>Landing Page</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <AuthDialog open={authOpen} setOpen={setAuthOpen} />

      <div className="space-y-20 overflow-hidden dark:text-white">
        <header className="px-4 sm:px-6 md:px-8">
          <div className="flex w-full items-center justify-between py-6">
            <button className="text-2xl font-bold">AI Presentation Tool</button>
            <div className="flex space-x-6">
              <nav className="flex items-center space-x-8 font-medium">
                <a
                  className="hover:text-indigo-500 dark:hover:text-indigo-400"
                  href="/docs"
                >
                  Docs
                </a>
                <a
                  className="hover:text-indigo-500 dark:hover:text-indigo-400"
                  href="/about"
                >
                  About
                </a>
              </nav>
              <div className="border-l border-slate-200 dark:border-slate-800"></div>
              <button
                className="text-slate-200 dark:text-slate-800"
                onClick={() => setDarkMode(!darkMode)}
              >
                {darkMode ? 'go light' : 'go dark'}
              </button>
            </div>
          </div>
        </header>

        <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
          <div className="max-w-full sm:max-w-xl md:max-w-2xl lg:max-w-3xl">
            <h1 className="text-center text-4xl font-extrabold tracking-tight sm:text-5xl lg:text-6xl">
              Generate{' '}
              <span className="bg-gradient-to-r from-pink-500 to-violet-500 bg-clip-text text-transparent">
                presentable slides
              </span>{' '}
              from your documents.
            </h1>

            <p className="mt-3 text-2xl">
              Get started by{' '}
              <button
                className="relative rounded-md bg-gray-200 px-2.5 py-1 font-mono text-lg transition duration-150 ease-in-out hover:bg-gray-300 dark:bg-gray-700 hover:dark:bg-gray-600"
                onClick={handleSignIn}
              >
                <span className="absolute -top-0.5 -right-0.5 flex h-2.5 w-2.5">
                  <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-indigo-400 opacity-75"></span>
                  <span className="relative inline-flex h-2.5 w-2.5 rounded-full bg-indigo-500"></span>
                </span>
                signing in
              </button>
            </p>
          </div>
        </main>

        <footer className="flex h-24 w-full items-center justify-center border-t border-slate-200 dark:border-slate-800">
          <a
            className="flex items-center justify-center"
            href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
            target="_blank"
            rel="noopener noreferrer"
          >
            Powered by{' '}
            <img src="/vercel.svg" alt="Vercel Logo" className="ml-2 h-4" />
          </a>
        </footer>
      </div>
    </>
  )
}

export default withUrqlClient((_ssrExchange, ctx) => ({
  url: AUTH_API_URL,
}))(IndexPage)

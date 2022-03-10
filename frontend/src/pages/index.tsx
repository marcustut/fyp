import Head from 'next/head'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { NextPage } from 'next'

import { useDarkMode } from '@/hooks'
import { AuthDialog } from '@/features/auth'
import { Icon } from '@iconify/react'
import { useAuth } from '@/lib/auth'

const IndexPage: NextPage = () => {
  const [authOpen, setAuthOpen] = useState<boolean>(false)
  const { darkMode, setDarkMode } = useDarkMode()
  const { user } = useAuth()

  const handleSignIn = () => setAuthOpen(true)

  return (
    <>
      <Head>
        <title>SliGen</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <AuthDialog open={authOpen} setOpen={setAuthOpen} />

      <div className="space-y-20 overflow-hidden dark:text-white">
        <header className="px-4 sm:px-6 md:px-8">
          <div className="flex w-full items-center justify-between py-6">
            <button className="flex items-center text-3xl font-bold text-indigo-600">
              <div className="mr-4 rounded-xl bg-indigo-600 p-2">
                <img src="/SliGenOutline.png" className="h-8 w-8" />
              </div>
              SliGen
            </button>
            <div className="flex space-x-6">
              {user && (
                <>
                  <nav className="flex items-center space-x-8 font-medium">
                    <a
                      className="hover:text-indigo-500 dark:hover:text-indigo-400"
                      href="/files"
                    >
                      Go to App
                    </a>
                  </nav>
                  <div className="border-l border-slate-700 dark:border-slate-200" />
                </>
              )}
              <button
                className="text-slate-700 dark:text-slate-200"
                onClick={() => setDarkMode(!darkMode)}
              >
                <Icon
                  icon={
                    darkMode
                      ? 'heroicons-outline:sun'
                      : 'heroicons-outline:moon'
                  }
                />
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

        <footer className="flex h-24 w-full items-center justify-center whitespace-pre-wrap border-t border-slate-200 dark:border-slate-800">
          built by{' '}
          <Link href="https://github.com/marcustut">
            <a className="font-medium text-slate-500">@marcustut </a>
          </Link>
          and{' '}
          <Link href="https://github.com/lianaling">
            <a className="font-medium text-slate-500">@lianaling</a>
          </Link>
        </footer>
      </div>
    </>
  )
}

export default IndexPage

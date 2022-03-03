import React from 'react'
import { NextPage } from 'next'
import Head from 'next/head'

import { SideNav, Spinner, AppLayout, InfoTab } from '@/components'
import { FileBrowser, Summary } from '@/features/storage'
import { CheckAuth, useAuth } from '@/lib/auth'

const FilesPage: NextPage = () => {
  const { user } = useAuth()

  return (
    <>
      <Head>
        <title>AI Summarizer - Files</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <CheckAuth>
        {!user ? (
          <Spinner />
        ) : (
          <AppLayout>
            <SideNav title="AI Summarizer" />
            <FileBrowser user={user} />
            <InfoTab
              mainRender={() => <Summary user={user} variant="visual" />}
              subRender={() => <Summary user={user} variant="detailed" />}
            />
          </AppLayout>
        )}
      </CheckAuth>
    </>
  )
}

export default FilesPage

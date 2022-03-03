import React from 'react'
import { NextPage } from 'next'
import Head from 'next/head'

import { CheckAuth, useAuth } from '@/lib/auth'
import { AppLayout, InfoTab, SideNav, Spinner } from '@/components'
import { FileBrowser, Summary } from '@/features/storage'

const SharedPage: NextPage = () => {
  const { user } = useAuth()

  return (
    <>
      <Head>
        <title>AI Summarizer - Shared Files</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <CheckAuth>
        {!user ? (
          <Spinner />
        ) : (
          <AppLayout>
            <SideNav title="AI Summarizer" />
            <FileBrowser
              user={user}
              variant="shared"
              emptyMessage="No files shared"
            />
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

export default SharedPage

import React, { useState } from 'react'
import dayjs from 'dayjs'
import { Icon } from '@iconify/react'
import { Column } from 'react-table'
import { TW } from 'tailwindcss-classnames'

import {
  Notification,
  Avatar,
  Button,
  AppContainer,
  Table,
  GlobalSearch,
} from '@/components'
import { UserWithAuth } from '@/generated/graphql'

type FileBrowserProps = {
  user: UserWithAuth
}

export const FileBrowser: React.FC<FileBrowserProps> = ({ user }) => {
  const [setGlobalFilter, setSetGlobalFilter] = useState<
    (filterValue: any) => void
  >(() => () => {})

  const data = [
    {
      title: 'Machine Learning Thesis',
      owner: '',
      sharedWith: ['USER_1231ASX132'],
      fileName: 'machine_learning_thesis.md',
      fileSize: 12344123,
      updatedAt: '2022-02-07T22:28:27+08:00',
      createdAt: '2022-02-06T22:28:27+08:00',
    },
    {
      title: 'Relationship between Tree and Trie asdhajdhkasdhklajsdlkas',
      owner: '',
      sharedWith: ['USER_1231ASX132', 'USER_1231ASX132'],
      fileName: 'tree_and_trie.pdf',
      fileSize: 1344123,
      updatedAt: '2022-02-08T23:28:27+08:00',
      createdAt: '2022-02-08T22:28:27+08:00',
    },
    {
      title: 'SPM Recording',
      owner: '',
      sharedWith: [],
      fileName: 'spm_recording.mp4',
      fileSize: 344123,
      updatedAt: '2022-02-16T22:28:27+08:00',
      createdAt: '2022-02-08T22:28:27+08:00',
    },
    {
      title: 'Screenshot (Firebase)',
      owner: '',
      sharedWith: [],
      fileName: 'firebase_screenshot.png',
      fileSize: 2444123,
      updatedAt: '2022-02-08T22:28:27+08:00',
      createdAt: '2022-02-08T22:28:27+08:00',
    },
  ]

  const columns =
    data.length === 0
      ? []
      : ([
          {
            id: 'Name',
            Header: 'Name',
            accessor: 'title',
            Cell: ({ value }) => {
              const type = data.filter((d) => d.title === value).shift()!.title
              return (
                <div className="flex w-96 items-center pr-4 font-medium text-slate-900">
                  <Icon
                    icon={
                      type === 'pdf'
                        ? 'icon-park-outline:file-pdf-one'
                        : type === 'mp4'
                        ? 'icon-park-outline:image-files'
                        : type === 'png'
                        ? 'icon-park-outline:video-file'
                        : 'icon-park-outline:file-doc'
                    }
                    className="mr-3 h-8 w-8"
                  />
                  <p className="w-full truncate">{value}</p>
                </div>
              )
            },
          },
          {
            Header: 'Member',
            accessor: (row) =>
              row.sharedWith.length === 0
                ? 'Only you'
                : `${row.sharedWith.length + 1} members`,
            Cell: ({ value }: { value: string }) => (
              <p className="pr-4">{value}</p>
            ),
          },
          {
            Header: 'Last Modified',
            accessor: (row) => dayjs(row.updatedAt).format('MMM D, YYYY'),
            Cell: ({ value }: { value: string }) => <p>{value}</p>,
          },
          {
            id: 'action',
            Header: '',
            Cell: () => (
              <button>
                <Icon icon="heroicons-solid:dots-horizontal" />
              </button>
            ),
          },
        ] as Column<typeof data[0]>[])

  return (
    <AppContainer>
      <div className="flex w-full items-center">
        <GlobalSearch
          setGlobalFilter={setGlobalFilter}
          placeholder="Search file..."
          icon={(defaultClasses) => (
            <Icon icon="heroicons-outline:search" className={defaultClasses} />
          )}
        />
        <Notification dot className="mr-8 h-8 w-8" />
        <Avatar
          src={user.user.avatar_url}
          gender={'male'}
          name={user.user.full_name ?? user.user.username}
          outline
          className="h-8 w-8"
        />
      </div>

      <div className="mt-12 flex flex-col">
        <div className="flex items-center justify-between">
          <h1 className="text-4xl font-semibold">My files</h1>
          <Button
            tw={(defaultClasses) => ({
              ...defaultClasses,
              display: TW.display('flex'),
              alignItems: TW.alignItems('items-center'),
              borderRadius: TW.borderRadius('rounded-full'),
              backgroundColor: TW.backgroundColor('bg-indigo-600'),
            })}
          >
            <Icon
              icon="heroicons-outline:plus-circle"
              className="mr-3 h-5 w-5"
            />
            New Slide
          </Button>
        </div>
      </div>

      <Table
        data={data}
        columns={columns}
        updateGlobalFilter={setSetGlobalFilter}
        className="mt-12 w-full"
      />
    </AppContainer>
  )
}

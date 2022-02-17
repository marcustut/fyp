import React, { useState } from 'react'
import dayjs from 'dayjs'
import { toast } from 'react-toastify'
import { Column } from 'react-table'
import { TW, TTailwindString } from 'tailwindcss-classnames'
import { Icon } from '@iconify/react'

import {
  Notification,
  Avatar,
  Button,
  AppContainer,
  Table,
  GlobalSearch,
  Menu,
} from '@/components'
import { UserWithAuth } from '@/generated/graphql'
import { AvatarDropdown } from '@/features/auth'
import { NewSlide } from '@/features/slide'

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
              const type = data
                .filter((d) => d.title === value)
                .shift()!
                .fileName.split('.')
                .reverse()[0]
              return (
                <div
                  className="flex w-full items-center py-3 pr-4 font-medium text-slate-900 hover:cursor-pointer dark:text-slate-50"
                  onClick={() => toast(value)}
                >
                  <Icon
                    icon={
                      type === 'pdf'
                        ? 'mdi:file-pdf'
                        : type === 'mp4'
                        ? 'mdi:file-video'
                        : type === 'png'
                        ? 'mdi:file-image'
                        : 'mdi:file-document'
                    }
                    className="mx-3 h-8 w-8"
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
            Cell: ({ value }: { value: string }) => <p className="">{value}</p>,
          },
          {
            id: 'action',
            Header: '',
            Cell: () => (
              <Menu
                tw={{ margin: TW.margin('mr-3') }}
                buttonRender={() => (
                  <Icon icon="heroicons-solid:dots-horizontal" />
                )}
                items={[
                  {
                    render: () => <>Open in New Tab</>,
                    icon: (defaultClasses) => (
                      <Icon
                        icon="heroicons-outline:external-link"
                        className={`${defaultClasses}`}
                      />
                    ),
                    onClick: () => toast.warn('not implemented'),
                  },
                  {
                    render: () => <>Delete</>,
                    icon: (defaultClasses) => (
                      <Icon
                        icon="heroicons-outline:trash"
                        className={defaultClasses}
                      />
                    ),
                    onClick: () => toast.warn('not implemented'),
                  },
                ]}
              />
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
          inputContainerClassName={'dark:focus-within:text-indigo-50'}
        />
        <Notification
          dot
          className="mr-8 h-8 w-8"
          onClick={() => toast.warn('not implemented')}
        />
        <AvatarDropdown
          avatarProps={{
            src: user.user.avatar_url,
            gender: 'male',
            name: user.user.full_name ?? user.user.username,
            outline: true,
            width: 'w-8',
            height: 'h-8',
          }}
        />
      </div>

      <div className="mt-12 flex flex-col">
        <div className="flex items-center justify-between">
          <h1 className="text-4xl font-semibold">My files</h1>
          <NewSlide
            buttonProps={{
              tw: (defaultClasses) => ({
                ...defaultClasses,
                display: TW.display('flex'),
                alignItems: TW.alignItems('items-center'),
                borderRadius: TW.borderRadius('rounded-full'),
                backgroundColor: TW.backgroundColor('bg-indigo-600'),
              }),
            }}
          />
        </div>
      </div>

      <Table
        data={data}
        columns={columns}
        updateGlobalFilter={setSetGlobalFilter}
        tableRowProps={{
          tw: {
            margin: TW.margin('mx-8'),
            textColor: TW.textColor('hover:text-indigo-50'),
          },
          className: 'group',
        }}
        tableCellProps={{
          tw: {
            padding: TW.padding('py-1'),
            backgroundColor: TW.backgroundColor(
              'group-hover:bg-indigo-200',
              'dark:group-hover:bg-indigo-500' as TTailwindString
            ),
            borderRadius: TW.borderRadius(
              'first:rounded-l-3xl',
              'last:rounded-r-3xl'
            ),
            transitionProperty: TW.transitionProperty('transition'),
            transitionDuration: TW.transitionDuration('duration-200'),
            transitionTimingFunction:
              TW.transitionTimingFunction('ease-in-out'),
          },
        }}
        className="mt-12 w-full"
      />
    </AppContainer>
  )
}

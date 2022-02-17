import React, { Fragment, useState } from 'react'
import { Icon } from '@iconify/react'
import { Tab } from '@headlessui/react'
import Dropzone from 'react-dropzone'
import { classnames, TW, TTailwindString } from 'tailwindcss-classnames'

import { Button, ButtonProps, Dialog } from '@/components'

type NewSlideProps = {
  buttonProps: ButtonProps
}

const types = ['Upload File', 'From URL', 'Type Text']

export const NewSlide: React.FC<NewSlideProps> = ({ buttonProps }) => {
  const [open, setOpen] = useState<boolean>(false)

  return (
    <>
      <Button {...buttonProps} onClick={() => setOpen(true)}>
        <Icon icon="heroicons-outline:plus-circle" className="mr-3 h-5 w-5" />
        New Slide
      </Button>

      <Dialog
        open={open}
        setOpen={setOpen}
        render={(_, setOpen) => (
          <>
            <Tab.Group>
              <Tab.List className="flex w-full space-x-1 rounded-xl bg-indigo-100 p-1 dark:bg-indigo-500">
                {types.map((type) => (
                  <Tab key={type} as={Fragment}>
                    {({ selected }) => (
                      <button
                        className={classnames(
                          TW.width('w-full'),
                          TW.padding('px-4', 'py-3'),
                          TW.borderRadius('rounded-lg'),
                          TW.textColor({
                            'text-slate-500': !selected,
                            'dark:text-slate-300': !selected,
                            ['hover:dark:text-slate-50' as TTailwindString]:
                              !selected,
                            'text-slate-50': selected,
                          }),
                          TW.fontSize('text-sm'),
                          TW.backgroundColor({
                            'bg-indigo-400': selected,
                            'bg-transparent': !selected,
                            'hover:bg-indigo-50': !selected,
                            ['dark:hover:bg-indigo-400' as TTailwindString]:
                              !selected,
                          }),
                          TW.outlineStyle('focus:outline-none'),
                          TW.ringWidth('focus:ring-2'),
                          TW.ringColor(
                            'ring-indigo-400',
                            'dark:ring-indigo-400'
                          ),
                          TW.ringOffsetWidth('ring-offset-2'),
                          TW.ringOffsetColor(
                            'ring-offset-indigo-100',
                            'dark:ring-offset-indigo-200'
                          ),
                          TW.ringOpacity('ring-opacity-60')
                        )}
                      >
                        {type}
                      </button>
                    )}
                  </Tab>
                ))}
              </Tab.List>
              <Tab.Panels className="mt-2 w-full rounded-xl bg-gradient-to-t from-indigo-200 to-indigo-100 p-2 dark:from-indigo-500 dark:to-indigo-400">
                <Tab.Panel>
                  <Dropzone
                    accept="application/pdf, text/plain"
                    onDrop={(acceptedFiles, fileRejections, event) => {
                      console.log(acceptedFiles)
                      console.log(fileRejections)
                      console.log(event)
                    }}
                  >
                    {({ getRootProps, getInputProps, acceptedFiles }) => (
                      <div>
                        <div
                          {...getRootProps()}
                          className={`flex flex-col items-center justify-center rounded-lg border-2 border-dashed border-indigo-300 px-8 dark:border-indigo-200 ${
                            acceptedFiles.length !== 0 ? 'py-8' : 'py-16'
                          } text-indigo-500 dark:text-indigo-50`}
                        >
                          {acceptedFiles.length === 0 ? (
                            <>
                              <input {...getInputProps()} />
                              <Icon
                                icon="heroicons-outline:upload"
                                className="h-16 w-16"
                              />
                              <p className="font-medium">
                                Drag 'n' drop your files here, or click to
                                select files
                              </p>
                              <p className="text-sm text-indigo-400 dark:text-indigo-200">
                                only .PDF and .TXT files are accepted
                              </p>
                            </>
                          ) : (
                            <pre>{JSON.stringify(acceptedFiles, null, 2)}</pre>
                          )}
                        </div>
                      </div>
                    )}
                  </Dropzone>
                </Tab.Panel>
                <Tab.Panel>From URL</Tab.Panel>
                <Tab.Panel>Type Text</Tab.Panel>
              </Tab.Panels>
            </Tab.Group>
          </>
        )}
      />
    </>
  )
}

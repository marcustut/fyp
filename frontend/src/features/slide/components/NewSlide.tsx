import React, { Fragment, useCallback, useMemo, useState } from 'react'
import { Icon } from '@iconify/react'
import { Listbox, Tab, Transition } from '@headlessui/react'
import Dropzone from 'react-dropzone'
import { classnames, TW, TTailwindString } from 'tailwindcss-classnames'

import { Button, ButtonProps, Dialog, Input } from '@/components'
import { toast } from 'react-toastify'
import { handleFileRejections } from '@/utils/handler'
import { SUMMARIZE_API_URL } from '@/lib/constants'
import { SummarizeClient, SummarizeOptions } from '@/lib/summarize'
import { isValidHttpUrl } from '@/utils/validate'
import {
  useCreateSlideWithTextMutation,
  UserWithAuth,
} from '@/generated/graphql'
import { generateName } from '@/utils/name'

type NewSlideProps = {
  user: UserWithAuth
  buttonProps: ButtonProps
}

enum UploadType {
  File = 'Upload File',
  URL = 'From URL',
  Text = 'Type Text',
}

type UserSummarizeOptions = Omit<
  Omit<Omit<SummarizeOptions, 'type'>, 'input'>,
  'outputName'
>

const modes = ['ext', 'abs']

export const NewSlide: React.FC<NewSlideProps> = ({ user, buttonProps }) => {
  const summarize = new SummarizeClient({ ServerURL: SUMMARIZE_API_URL })
  const [open, setOpen] = useState<boolean>(false)
  const [loading, setLoading] = useState<boolean>(false)
  const [uploadType, setUploadType] = useState<UploadType>(UploadType.File)
  const [url, setUrl] = useState<string>('')
  const [text, setText] = useState<string>('')
  const [summarizeOpts, setSummarizeOpts] = useState<UserSummarizeOptions>({
    mode: 'ext',
    theme: 'apple-basic',
    maxChunk: 500,
    maxCharPerSlide: 500,
  })
  const [stagedFile, setStagedFile] =
    useState<{ blob: Blob; fileName: string }>()
  const [_, createSlideWithText] = useCreateSlideWithTextMutation()

  const summarizationMode = useMemo(
    () =>
      summarizeOpts.mode === 'ext' ? 'Original Text' : 'AI Generated Text',
    [summarizeOpts]
  )

  const handleCreateSlide = useCallback(
    async (mdText: string, fileName: string) => {
      const { data, error } = await createSlideWithText({
        createSlideWithTextInput: {
          name: fileName,
          user_id: user.user.id,
        },
        text: mdText,
      })
      if (error || !data) {
        toast.error('An error occured while creating slide')
        return
      }
      toast.success('redirecting you to the slide...')
      setTimeout(
        () => window.open(`slide/${data.CreateSlideWithText.id}`),
        1000
      )
    },
    [user, createSlideWithText]
  )

  const handleSummarize = useCallback(async () => {
    setLoading(true)
    const outputName = `Slide ${generateName()}.md`
    switch (uploadType) {
      case UploadType.File: {
        if (!stagedFile) {
          toast.error('no file is selected')
          setLoading(false)
          return
        }
        await summarize.uploadFile(stagedFile.blob, stagedFile.fileName)
        try {
          const mdText = await summarize.summarize({
            type: 'pdf',
            input: stagedFile.fileName,
            outputName,
            ...summarizeOpts,
            maxChunk: 250,
          })
          await handleCreateSlide(mdText, outputName)
        } catch (err) {
          console.error(err)
          toast.error((err as Error).message)
        }
        break
      }
      case UploadType.URL: {
        if (!isValidHttpUrl(url)) {
          toast.error(`'${url}' is not a valid URL`)
          setLoading(false)
          return
        }
        try {
          const mdText = await summarize.summarize({
            type: 'url',
            input: url,
            outputName,
            ...summarizeOpts,
          })
          await handleCreateSlide(mdText, outputName)
        } catch (err) {
          console.error(err)
          toast.error((err as Error).message)
        }
        break
      }
      case UploadType.Text: {
        if (text === '') {
          toast.error('text cannot be empty')
          setLoading(false)
          return
        }
        const fileName = 'input.txt'
        await summarize.uploadFile(new Blob([text]), fileName)
        try {
          const mdText = await summarize.summarize({
            type: 'txt',
            input: fileName,
            outputName,
            ...summarizeOpts,
          })
          await handleCreateSlide(mdText, outputName)
        } catch (err) {
          console.error(err)
          toast.error((err as Error).message)
        }
        break
      }
    }
    setLoading(false)
  }, [uploadType, stagedFile, url, text, setLoading])

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
            <Tab.Group
              onChange={(index) =>
                setUploadType(
                  index === 0
                    ? UploadType.File
                    : index === 1
                    ? UploadType.URL
                    : UploadType.Text
                )
              }
            >
              <Tab.List className="flex w-full space-x-1 rounded-xl bg-indigo-100 p-1 dark:bg-indigo-500">
                {Object.values(UploadType).map((type) => (
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
                    multiple={false}
                    maxFiles={1}
                    maxSize={10000000}
                    onDropAccepted={(acceptedFiles, event) => {
                      setLoading(true)
                      setStagedFile({
                        blob: acceptedFiles[0],
                        fileName: acceptedFiles[0].name.replaceAll(' ', '_'),
                      })
                      setLoading(false)
                    }}
                    onDropRejected={handleFileRejections}
                  >
                    {({ getRootProps, getInputProps, acceptedFiles }) => (
                      <div>
                        <div
                          {...getRootProps({
                            className: `flex flex-col items-center justify-center rounded-lg border-2 border-dashed border-indigo-300 px-8 dark:border-indigo-200 ${
                              acceptedFiles.length !== 0 ? 'py-8' : 'py-16'
                            } text-indigo-500 dark:text-indigo-50`,
                          })}
                        >
                          {acceptedFiles.length === 0 ? (
                            <>
                              <input {...getInputProps()} />
                              <Icon
                                icon="heroicons-outline:upload"
                                className="h-16 w-16"
                              />
                              <p className="text-center font-medium">
                                Drag 'n' drop your files here, or click to
                                select files
                              </p>
                              <p className="text-center text-sm text-indigo-400 dark:text-indigo-200">
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
                <Tab.Panel as="div" className="px-4 py-3">
                  <p className="pb-4 text-slate-500 dark:text-slate-50">
                    💡 This is most suitable for summarizing web articles.
                  </p>
                  <Input
                    variant="primary"
                    defaultValue="https://"
                    className="border border-indigo-400 dark:border-slate-300 dark:bg-transparent dark:text-slate-400 dark:focus-within:text-indigo-50"
                    icon={(defaultClasses) => (
                      <Icon
                        icon="heroicons-outline:globe-alt"
                        className={defaultClasses}
                      />
                    )}
                    onChange={(event) => setUrl(event.target.value)}
                  />
                </Tab.Panel>
                <Tab.Panel as="div" className="px-4 py-3">
                  <p className="pb-4 text-slate-500 dark:text-slate-50">
                    You can type into the textbox below for summarization
                  </p>
                  <Input
                    textarea
                    variant="primary"
                    placeholder="Enter your text here..."
                    className="border border-indigo-400 dark:border-slate-300 dark:bg-transparent dark:text-slate-400 dark:placeholder:text-slate-300 dark:focus-within:text-indigo-50"
                    onChange={(event) => setText(event.target.value)}
                  />
                </Tab.Panel>
              </Tab.Panels>
              <Listbox
                value={summarizeOpts.mode}
                onChange={(mode) =>
                  setSummarizeOpts({ ...summarizeOpts, mode })
                }
              >
                <p className="mt-4 mb-1 self-start text-left text-sm font-medium">
                  Summarization Mode
                </p>
                <Listbox.Button className="relative w-full cursor-default rounded-lg bg-indigo-200 py-2 pl-3 pr-10 text-left focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm">
                  <span className="block truncate">{summarizationMode}</span>
                  <span className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                    <Icon
                      icon="heroicons-outline:selector"
                      className="h-5 w-5 text-gray-700"
                      aria-hidden="true"
                    />
                  </span>
                </Listbox.Button>
                <Transition
                  as={Fragment}
                  leave="transition ease-in duration-100"
                  leaveFrom="opacity-100"
                  leaveTo="opacity-0"
                >
                  <Listbox.Options className="mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                    {modes.map((mode, idx) => (
                      <Listbox.Option
                        key={idx}
                        className={({ active }) =>
                          `relative cursor-default select-none py-2 pl-10 pr-4 ${
                            active
                              ? 'bg-indigo-100 text-indigo-700'
                              : 'text-gray-900'
                          }`
                        }
                        value={mode}
                      >
                        {({ selected }) => (
                          <>
                            <span
                              className={`block truncate ${
                                selected ? 'font-medium' : 'font-normal'
                              }`}
                            >
                              {mode === 'ext'
                                ? 'Original Text'
                                : 'AI Generated Text'}
                            </span>
                            {selected ? (
                              <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-indigo-600">
                                <Icon
                                  icon="heroicons-outline:check"
                                  className="h-5 w-5"
                                  aria-hidden="true"
                                />
                              </span>
                            ) : null}
                          </>
                        )}
                      </Listbox.Option>
                    ))}
                  </Listbox.Options>
                </Transition>
              </Listbox>
              <Button
                loading={loading}
                variant="primary"
                tw={(defaultClasses) => ({
                  ...defaultClasses,
                  margin: TW.margin('mt-4'),
                  width: TW.width('w-full'),
                  display: TW.display('flex'),
                  alignItems: TW.alignItems('items-center'),
                  justifyContent: TW.justifyContent('justify-center'),
                  borderRadius: TW.borderRadius('rounded-lg'),
                  backgroundColor: TW.backgroundColor(
                    'bg-indigo-500',
                    'hover:bg-indigo-400'
                  ),
                })}
                onClick={handleSummarize}
              >
                Summarize
              </Button>
            </Tab.Group>
          </>
        )}
      />
    </>
  )
}

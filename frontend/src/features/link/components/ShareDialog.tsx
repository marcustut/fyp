import { z } from 'zod'
import React, { useMemo, useState } from 'react'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { toast } from 'react-toastify'

import {
  Slide,
  useAddUsersToSlideMutation,
  UserWithAuth,
  useShortenUrlMutation,
  useUsersQuery,
} from '@/generated/graphql'
import { Dialog, Button, Input, Spinner } from '@/components'
import { LINK_API_URL } from '@/lib/constants'
import { Icon } from '@iconify/react'

const SlideShareSchema = z.object({
  emails: z.string().min(4),
})

type ShareDialogProps = {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  slide: Slide
  user: UserWithAuth
}

export const ShareDialog: React.FC<ShareDialogProps> = ({
  open,
  setOpen,
  slide,
  user,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof SlideShareSchema>>({
    resolver: zodResolver(SlideShareSchema),
  })
  const [loading, setLoading] = useState<boolean>(false)
  const [_shortenUrl, shortenUrl] = useShortenUrlMutation()
  const [_addUsersToSlide, addUsersToSlide] = useAddUsersToSlideMutation()
  const [sharedUsersState] = useUsersQuery({
    variables: { where: { idIn: slide.shared_with } },
  })
  const [generatedUrl, setGeneratedUrl] = useState<string>()

  const originalUrl = useMemo(
    () => `http://${window.location.host}/slide/${slide.id}`,
    [slide]
  )

  const sharedUsers = useMemo(() => {
    if (sharedUsersState.data && sharedUsersState.data.Users.edges)
      return [
        user.user,
        ...sharedUsersState.data.Users.edges.map((e) => e?.node),
      ]
    return [user.user]
  }, [sharedUsersState])

  const handleCopyLink = async () => {
    setLoading(true)
    const { data, error } = await shortenUrl({
      input: {
        original_url: originalUrl,
        owner_id: slide.user.id,
      },
    })
    if (error) toast.error(error.message)
    else if (!data) toast.error(`unexpected error`)
    else {
      const url = `${LINK_API_URL}/${data.CreateLinkOptionalLinkID.link_id}`
      setGeneratedUrl(url)
      navigator.clipboard.writeText(url)
      toast.success(`link for ${slide.name} is copied to clipboard`)
    }
    setLoading(false)
  }

  const handleSendInvite = async (values: z.infer<typeof SlideShareSchema>) => {
    setLoading(true)
    const emails = values.emails.split(',').map((email) => email.trim())
    const { data, error } = await addUsersToSlide({ id: slide.id, emails })
    if (error) toast.error(error.message)
    else if (!data) toast.error(`unexpected error`)
    else toast.success(`invite had been sent to ${emails.join(', ')}`)
    setLoading(false)
  }

  return (
    <Dialog
      open={open}
      setOpen={setOpen}
      render={(open, setOpen) => (
        <>
          <h2 className="w-full text-sm tracking-tight">
            Share {user.user.username}'s{' '}
            <span className="font-medium">{slide.name}</span>
          </h2>

          <div className="mt-2 mb-3 w-full border-b border-slate-300 dark:border-slate-200" />

          <div className="flex w-full flex-col">
            <div className="flex w-full items-center">
              <form className="w-full">
                <label className="ml-1 text-sm font-medium"></label>
                <Input
                  register={register('emails')}
                  variant="primary"
                  type="text"
                  className="mt-1 w-full"
                  placeholder="Email, comma-separated"
                />
                <p className="mt-1 ml-1 text-xs text-red-500">
                  {errors.emails && errors.emails.message}
                </p>
              </form>

              <Button
                loading={loading}
                className="ml-4 w-36 py-2.5 text-xs"
                onClick={handleSubmit(handleSendInvite)}
              >
                Send Invite
              </Button>
            </div>

            <div className="mt-3 mb-5 w-full border-b border-slate-300 dark:border-slate-200" />
            <p className="flex items-center px-2 text-sm text-slate-400">
              <Icon icon="heroicons-outline:users" className="mr-1" />
              Shared users
            </p>
            <div className="flex w-full items-center px-2">
              {sharedUsersState.fetching ? (
                <Spinner />
              ) : (
                sharedUsers.map((u) => {
                  if (!u) return <></>
                  return (
                    <div
                      key={u.username}
                      className="mt-2 flex w-full items-center text-sm"
                    >
                      <img
                        src={
                          u.avatar_url ??
                          `https://avatars.dicebear.com/api/male/${
                            u.full_name ?? u.username
                          }.svg?mood[]=happy`
                        }
                        className="mr-2 h-6 w-6 rounded-full"
                      />
                      {u.username}{' '}
                      {u.username === user.user.username && '(you)'}
                    </div>
                  )
                })
              )}
            </div>

            <div className="mt-5 mb-3 w-full border-b border-slate-300 dark:border-slate-200" />
            <button
              className="flex items-center text-left text-sm text-indigo-500"
              onClick={handleCopyLink}
            >
              <Icon icon="heroicons-outline:link" className="mr-1 h-4 w-4" />
              Copy Link
            </button>
          </div>
        </>
      )}
    />
  )
}

import { z } from 'zod'
import React, { useEffect, useState } from 'react'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { toast } from 'react-toastify'

import { Slide, useShortenUrlMutation } from '@/generated/graphql'
import { Dialog, Button, Input } from '@/components'
import { LINK_API_URL } from '@/lib/constants'

const SlideShareSchema = z.object({
  link_id: z.string().min(4).max(8),
})

type ShareDialogProps = {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  slide: Slide
}

export const ShareDialog: React.FC<ShareDialogProps> = ({
  open,
  setOpen,
  slide,
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
  const [generatedUrl, setGeneratedUrl] = useState<string>()

  // TODO: Fix the url not valid issue
  useEffect(() => {
    ;(async () => {
      setLoading(true)
      const { data, error } = await shortenUrl({
        input: {
          original_url: `http://${window.location.host}/slide/${slide.id}`,
          owner_id: slide.user.id,
        },
      })
      if (error) toast.error(error.message)
      else if (!data) toast.error(`unexpected error`)
      else {
        const url = `${LINK_API_URL}/${data.CreateLinkOptionalLinkID.link_id}`
        console.log(url)
        setGeneratedUrl(url)
        navigator.clipboard.writeText(url)
        toast.success(`successfully shortened the link`)
      }
      setLoading(false)
    })()
  }, [])

  const onSubmit = async (values: z.infer<typeof SlideShareSchema>) => {
    setLoading(true)
    console.log(values)
    setLoading(false)
    setOpen(false)
  }

  return (
    <Dialog
      open={open}
      setOpen={setOpen}
      render={(open, setOpen) => (
        <>
          <form className="w-full">
            <label className="ml-1 text-sm font-medium"></label>
            <Input
              register={register('link_id')}
              variant="primary"
              type="text"
              className="mt-1 w-full"
              placeholder="enter here..."
              defaultValue={generatedUrl}
            />
            <p className="mt-1 ml-1 text-xs text-red-500">
              {errors.link_id && errors.link_id.message}
            </p>
          </form>

          <Button
            loading={loading}
            className="mt-4 w-full"
            onClick={handleSubmit(onSubmit)}
          >
            Rename
          </Button>
        </>
      )}
    />
  )
}

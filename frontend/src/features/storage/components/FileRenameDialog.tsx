import { Button, Dialog, Input } from '@/components'
import { Slide, useUpdateSlideMutation } from '@/generated/graphql'
import { zodResolver } from '@hookform/resolvers/zod'
import React, { useState } from 'react'
import { useForm } from 'react-hook-form'
import { toast } from 'react-toastify'
import { z } from 'zod'

const SlideRenameSchema = z.object({
  name: z.string().nonempty(),
})

type FileRenameDialogProps = {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  originalFile: Slide
}

export const FileRenameDialog: React.FC<FileRenameDialogProps> = ({
  open,
  setOpen,
  originalFile,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof SlideRenameSchema>>({
    resolver: zodResolver(SlideRenameSchema),
  })
  const [loading, setLoading] = useState<boolean>(false)
  const [_updateSlideState, updateSlide] = useUpdateSlideMutation()

  const onSubmit = async (values: z.infer<typeof SlideRenameSchema>) => {
    setLoading(true)
    const { data, error } = await updateSlide({
      input: { id: originalFile.id, name: values.name },
    })
    if (error) toast.error(error.message)
    else if (!data) toast.error(`unexpected error`)
    else toast.success(`successfully renamed to ${data.UpdateSlide.name}`)
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
            <label className="ml-1 text-sm font-medium">New Name</label>
            <Input
              register={register('name')}
              variant="primary"
              type="text"
              className="mt-1 w-full"
              placeholder="enter here..."
              defaultValue={originalFile.name}
            />
            <p className="mt-1 ml-1 text-xs text-red-500">
              {errors.name && errors.name.message}
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

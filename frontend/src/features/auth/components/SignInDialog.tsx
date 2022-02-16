import React from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { toast } from 'react-toastify'
import { Icon } from '@iconify/react'
import { zodResolver } from '@hookform/resolvers/zod'

import { AuthView } from '@/features/auth'
import { regex } from '@/utils/regex'
import { useAuth } from '@/lib/auth'
import { Input, Button, Dialog } from '@/components'

const LoginCredentials = z.object({
  emailOrUsername: z.string().nonempty(),
  password: z.string().min(8).max(24),
})

type SignInDialogProps = {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  setAuthView: React.Dispatch<React.SetStateAction<AuthView>>
}

export const SignInDialog: React.FC<SignInDialogProps> = ({
  open,
  setOpen,
  setAuthView,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof LoginCredentials>>({
    resolver: zodResolver(LoginCredentials),
  })
  const { loading, signIn } = useAuth()

  const onSubmit = async (creds: z.infer<typeof LoginCredentials>) => {
    if (creds.emailOrUsername.toLowerCase().match(regex.email)) {
      const err = await signIn({
        type: 'email',
        email: creds.emailOrUsername,
        password: creds.password,
      })
      if (err) toast(err.message, { type: 'error' })
    } else {
      const err = await signIn({
        type: 'username',
        username: creds.emailOrUsername,
        password: creds.password,
      })
      if (err) toast(err.message, { type: 'error' })
    }
  }

  return (
    <Dialog
      open={open}
      setOpen={setOpen}
      render={(_, setOpen) => (
        <>
          <h2 className="w-full text-sm font-bold uppercase tracking-tight">
            Log in to your account
          </h2>

          <form className="mt-4 w-full">
            <label className="ml-1 text-sm font-medium">Email / Username</label>
            <Input
              register={register('emailOrUsername')}
              variant="primary"
              type="text"
              className="mt-1 w-full"
              placeholder="enter here..."
              icon={(defaultClasses) => (
                <Icon
                  icon="heroicons-outline:user"
                  className={`${defaultClasses}`}
                />
              )}
            />
            <p className="mt-1 ml-1 text-xs text-red-500">
              {errors.emailOrUsername && errors.emailOrUsername.message}
            </p>
            <label className="ml-1 text-sm font-medium">Password</label>
            <Input
              register={register('password')}
              variant="primary"
              type="password"
              className="mt-1 w-full"
              placeholder="enter here..."
              icon={(defaultClasses) => (
                <Icon
                  icon="heroicons-outline:lock-closed"
                  className={`${defaultClasses}`}
                />
              )}
            />
            <p className="mt-1 ml-1 text-xs text-red-500">
              {errors.password && errors.password.message}
            </p>
          </form>

          <Button
            loading={loading}
            className="mt-4 w-full"
            onClick={handleSubmit(onSubmit)}
          >
            Log in
          </Button>
          <Button
            variant="secondary"
            className="mt-2 flex w-full items-center justify-center"
            onClick={() => signIn({ type: 'github' })}
          >
            <Icon icon="mdi:github" className="mr-2 h-6 w-6" /> Continue with
            GitHub
          </Button>
          <button
            className="mt-3 text-sm tracking-tight text-indigo-400"
            onClick={() => setAuthView(AuthView.SignUp)}
          >
            Register an account
          </button>
        </>
      )}
    />
  )
}

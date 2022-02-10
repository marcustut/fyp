import React from 'react'
import { useForm } from 'react-hook-form'
import { Icon } from '@iconify/react'
import { Dialog } from '@/components/Dialog'
import { zodResolver } from '@hookform/resolvers/zod'
import { AuthView } from '@/features/auth'
import { z } from 'zod'
import { Button } from '@/components/Button'

const SignUpCredentials = z.object({
  email: z.string().email(),
  username: z.string().min(3).max(50),
  password: z.string().min(8).max(24),
})

type SignUpDialogProps = {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  setAuthView: React.Dispatch<React.SetStateAction<AuthView>>
}

export const SignUpDialog: React.FC<SignUpDialogProps> = ({
  open,
  setOpen,
  setAuthView,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof SignUpCredentials>>({
    resolver: zodResolver(SignUpCredentials),
  })

  const onSubmit = (creds: any) => {
    console.log(creds)
  }

  return (
    <Dialog
      open={open}
      setOpen={setOpen}
      render={(_, setOpen) => (
        <>
          <h2 className="w-full text-sm font-bold uppercase tracking-tight">
            Sign up an account
          </h2>

          <form className="mt-4 w-full">
            <label className="ml-1 text-sm font-medium">Email</label>
            <input
              {...register('email')}
              type="text"
              className="mt-1 w-full rounded-md border bg-slate-100 px-3 py-2 text-sm text-black transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-400 dark:bg-slate-300 dark:placeholder:text-slate-500 dark:focus:ring-indigo-500"
              placeholder="enter here..."
            />
            <p className="mt-1 ml-1 text-xs text-red-500">
              {errors.email && errors.email.message}
            </p>
            <label className="ml-1 text-sm font-medium">Username</label>
            <input
              {...register('username')}
              type="text"
              className="mt-1 w-full rounded-md border bg-slate-100 px-3 py-2 text-sm text-black transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-400 dark:bg-slate-300 dark:placeholder:text-slate-500 dark:focus:ring-indigo-500"
              placeholder="enter here..."
            />
            <p className="mt-1 ml-1 text-xs text-red-500">
              {errors.username && errors.username.message}
            </p>
            <label className="ml-1 text-sm font-medium">Password</label>
            <input
              {...register('password')}
              type="password"
              className="mt-1 w-full rounded-md border bg-slate-100 px-3 py-2 text-sm tracking-tight text-black transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-400 dark:bg-slate-300 dark:placeholder:text-slate-500 dark:focus:ring-indigo-500"
              placeholder="enter here..."
            />
            <p className="mt-1 ml-1 text-xs text-red-500">
              {errors.password && errors.password.message}
            </p>
          </form>

          <Button className="mt-4 w-full" onClick={handleSubmit(onSubmit)}>
            Sign Up
          </Button>
          <button
            className="mt-3 text-sm tracking-tight text-indigo-400"
            onClick={() => setAuthView(AuthView.SignIn)}
          >
            Back to Log in
          </button>
        </>
      )}
    />
  )
}

import React, { Fragment, useState } from 'react'
import { Dialog, Transition } from '@headlessui/react'
import { useForm } from 'react-hook-form'
import { Icon } from '@iconify/react'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'

type AuthDialogProps = {
  initialView?: 'signin' | 'signup'
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

const LoginCredentials = z.object({
  emailOrUsername: z.string().nonempty(),
  password: z.string().nonempty(),
})

export const AuthDialog: React.FC<AuthDialogProps> = ({
  initialView = 'signin',
  open,
  setOpen,
}) => {
  const [view, setView] = useState<'signin' | 'signup'>(initialView)
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<z.infer<typeof LoginCredentials>>({
    resolver: zodResolver(LoginCredentials),
  })

  console.log(errors)

  const onSubmit = (creds: any) => {
    console.log(creds)
  }

  return (
    <Transition appear show={open} as={Fragment}>
      <Dialog onClose={() => setOpen(false)}>
        <Transition.Child
          as={Dialog.Overlay}
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
          className="fixed inset-0 bg-black bg-opacity-30 backdrop-blur-[2px]"
        />

        <Transition.Child
          as="div"
          enter="ease-out duration-300"
          enterFrom="opacity-0 scale-95"
          enterTo="opacity-100 scale-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100 scale-100"
          leaveTo="opacity-0 scale-95"
          className="fixed top-1/2 left-1/2 z-10 flex min-w-[300px] -translate-x-1/2 -translate-y-1/2 transform flex-col items-center rounded-lg bg-slate-50 px-6 py-4 text-black shadow-2xl dark:bg-slate-700 dark:text-slate-100"
        >
          <Dialog.Title className="w-full text-sm font-bold uppercase tracking-tight">
            Log in to your account
          </Dialog.Title>

          <form className="mt-4 w-full">
            <label className="ml-1 text-sm font-medium">Email / Username</label>
            <input
              {...register('emailOrUsername')}
              type="text"
              className="mt-1 w-full rounded-md border bg-slate-100 px-3 py-2 text-sm text-black transition duration-150 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-400 dark:bg-slate-300 dark:placeholder:text-slate-500 dark:focus:ring-indigo-500"
              placeholder="enter here..."
            />
            <p className="mt-1 ml-1 text-xs text-red-500">
              {errors.emailOrUsername && errors.emailOrUsername.message}
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

          <button
            className="mt-4 w-full rounded-md bg-indigo-400 py-3 text-sm font-medium text-white transition duration-200 ease-in-out hover:bg-indigo-500 dark:bg-indigo-500 dark:hover:bg-indigo-400"
            type="submit"
            onClick={handleSubmit(onSubmit)}
          >
            Log in
          </button>
          <button
            className="mt-2 flex w-full items-center justify-center rounded-md bg-slate-200 py-3 text-sm text-slate-500 transition duration-200 ease-in-out hover:bg-slate-300 dark:bg-slate-600 dark:text-slate-50 dark:hover:bg-slate-500"
            onClick={() => setOpen(false)}
          >
            <Icon icon="mdi:github" className="mr-2 h-6 w-6" /> Continue with
            GitHub
          </button>
          <button className="mt-3 text-sm tracking-tight text-indigo-400">
            Register an account
          </button>
        </Transition.Child>
      </Dialog>
    </Transition>
  )
}

import React from 'react'
import { UseFormRegisterReturn } from 'react-hook-form'

export type InputProps = React.DetailedHTMLProps<
  React.InputHTMLAttributes<HTMLInputElement>,
  HTMLInputElement
> & {
  variant: 'primary' | 'transparent'
  register?: UseFormRegisterReturn
  icon?: (defaultClasses: string) => JSX.Element
  textarea?: boolean
  inputClassName?: string
}

export const Input = ({
  variant,
  register,
  icon,
  textarea = false,
  className = '',
  inputClassName = '',
  ...props
}: InputProps) => {
  const inputCn = `w-full bg-transparent placeholder:text-slate-400 focus-within:outline-none dark:placeholder:text-slate-500`
  const cn = `${inputCn} rounded-md ${
    variant !== 'transparent' ? 'border dark:border-slate-50' : ''
  } px-3 py-2 text-sm text-black transition duration-150 ease-in-out text-slate-400 focus-within:text-indigo-500 ${
    variant === 'primary' ? 'focus-within:ring-2' : 'focus-within:ring-0'
  } focus-within:ring-indigo-400 ${
    variant !== 'transparent' ? 'dark:bg-slate-300' : 'dark:bg-transparent'
  } dark:focus-within:ring-indigo-500`

  const renderInput = () => {
    const _cn = `${
      icon ? `${inputCn} ${inputClassName}` : `${cn} ${className}`
    }`
    if (register && textarea)
      return <textarea {...register} {...(props as unknown)} className={_cn} />
    else if (register) return <input {...register} {...props} className={_cn} />
    else if (textarea)
      return <textarea {...(props as unknown)} className={_cn} />
    else return <input {...props} className={_cn} />
  }

  return (
    <>
      {icon ? (
        <div className={`flex items-center ${cn} ${className}`}>
          {icon('h-4 w-4 mr-2')}
          {renderInput()}
        </div>
      ) : (
        <>{renderInput()}</>
      )}
    </>
  )
}

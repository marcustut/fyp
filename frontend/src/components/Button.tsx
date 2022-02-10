import React from 'react'
import { Icon } from '@iconify/react'
import Link from 'next/link'

type ButtonProps = React.DetailedHTMLProps<
  React.ButtonHTMLAttributes<HTMLButtonElement>,
  HTMLButtonElement
> & {
  href?: string
  variant?: 'primary' | 'secondary'
  loading?: boolean
  className?: string
}

export const Button: React.FC<ButtonProps> = ({
  href,
  variant = 'primary',
  loading = false,
  className = '',
  children,
  ...props
}) => {
  const bg =
    variant === 'primary'
      ? `bg-indigo-400 hover:bg-indigo-500 dark:bg-indigo-500 dark:hover:bg-indigo-400`
      : `bg-slate-200 hover:bg-slate-300 dark:bg-slate-600 dark:hover:bg-slate-500`
  const text =
    variant === 'primary' ? `text-white` : `text-slate-500 dark:text-slate-50`
  const cn = `${className} ${bg} ${text} rounded-md py-3 text-sm font-medium transition duration-200 ease-in-out`
  const inner = (
    <>
      {loading ? (
        <Icon icon="mdi:loading" className="mx-auto h-5 w-5 animate-spin" />
      ) : (
        children
      )}
    </>
  )

  return (
    <>
      {href ? (
        <Link href={href}>
          <a className={cn}>{inner}</a>
        </Link>
      ) : (
        <button className={cn} {...props}>
          {inner}
        </button>
      )}
    </>
  )
}

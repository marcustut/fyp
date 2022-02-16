import React, { useMemo } from 'react'
import Link from 'next/link'
import { classnames, TW, TTailwindString } from 'tailwindcss-classnames'

import { Spinner } from '@/components/Spinner'
import { TailwindClasses } from '@/types/utils'

type ButtonVariant = 'primary' | 'secondary' | 'transparent'

type ButtonProps = React.DetailedHTMLProps<
  React.ButtonHTMLAttributes<HTMLButtonElement>,
  HTMLButtonElement
> & {
  href?: string
  variant?: ButtonVariant
  loading?: boolean
  tw?: (defaultClasses: TailwindClasses) => TailwindClasses
  className?: string
}

const defaultClasses = {
  padding: TW.padding('px-4', 'py-3'),
  borderRadius: TW.borderRadius('rounded-md'),
}

export const Button: React.FC<ButtonProps> = ({
  href,
  variant = 'primary',
  loading = false,
  tw = (defaultClasses) => defaultClasses,
  className = '',
  children,
  ...props
}) => {
  const classes = {
    common: classnames(
      TW.fontSize('text-sm'),
      TW.fontWeight('font-medium'),
      TW.transitionProperty('transition'),
      TW.transitionDuration('duration-200'),
      TW.transitionTimingFunction('ease-in-out'),
      ...Object.values(tw(defaultClasses))
    ),
    primary: classnames(
      TW.backgroundColor(
        'bg-indigo-400',
        'hover:bg-indigo-500',
        'dark:bg-indigo-500',
        'dark:hover:bg-indigo-400' as TTailwindString
      ),
      TW.textColor('text-white')
    ),
    secondary: classnames(
      TW.backgroundColor(
        'bg-slate-200',
        'hover:bg-slate-300',
        'dark:bg-slate-600',
        'dark:hover:bg-slate-500' as TTailwindString
      ),
      TW.textColor('text-slate-500', 'dark:text-slate-50')
    ),
    transparent: classnames(TW.backgroundColor('bg-transparent')),
  }
  const variantClasses = useMemo(
    () =>
      classnames(
        classes.common,
        variant === 'primary'
          ? classes.primary
          : variant === 'secondary'
          ? classes.secondary
          : classes.transparent
      ),
    [variant]
  )
  const cn = `${variantClasses} ${className}`
  const inner = <>{loading ? <Spinner /> : children}</>

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

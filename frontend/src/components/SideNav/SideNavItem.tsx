import React from 'react'
import { useRouter } from 'next/router'
import { Button } from '@/components'

type SideNavItemProps = {
  name: string
  icon: (defaultClass: string, active: boolean) => JSX.Element
  path: string
  active?: boolean
  className?: string
}

export const SideNavItem: React.FC<SideNavItemProps> = ({
  name,
  icon,
  path,
  active = false,
  className = '',
}) => {
  const { push } = useRouter()
  return (
    <Button
      className={`flex items-center justify-start font-medium ${className} ${
        active
          ? 'bg-indigo-100/50 text-indigo-600 hover:bg-indigo-100 dark:bg-indigo-800 dark:text-indigo-50 dark:hover:bg-indigo-700'
          : 'bg-transparent text-slate-400 hover:bg-indigo-100/50 hover:text-indigo-600 dark:bg-transparent dark:hover:bg-indigo-800 dark:hover:text-indigo-50'
      }`}
      onClick={() => push(path)}
    >
      {icon('mr-3 h-6 w-6', active)}
      {name}
    </Button>
  )
}

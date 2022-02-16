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
      className={`flex items-center justify-start bg-transparent font-medium ${className} ${
        active
          ? 'bg-indigo-100/50 text-indigo-600 hover:bg-indigo-100'
          : 'text-slate-400 hover:bg-indigo-100/50 hover:text-indigo-600'
      }`}
      onClick={() => push(path)}
    >
      {icon('mr-3 h-6 w-6', active)}
      {name}
    </Button>
  )
}

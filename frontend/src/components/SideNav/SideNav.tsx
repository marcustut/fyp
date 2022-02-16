import React from 'react'
import { useRouter } from 'next/router'
import { Icon } from '@iconify/react'

import { CTA, SideNavItem } from '@/components/SideNav'

const mainNavItems = {
  'My files': {
    icon: (defaultClasses: string, active: boolean) => (
      <Icon
        icon={`heroicons-${active ? 'solid' : 'outline'}:home`}
        className={defaultClasses}
      />
    ),
    path: '/files',
  },
  'Shared file': {
    icon: (defaultClasses: string, active: boolean) => (
      <Icon
        icon={`heroicons-${active ? 'solid' : 'outline'}:users`}
        className={defaultClasses}
      />
    ),
    path: '/shared',
  },
  Starred: {
    icon: (defaultClasses: string, active: boolean) => (
      <Icon
        icon={`heroicons-${active ? 'solid' : 'outline'}:star`}
        className={defaultClasses}
      />
    ),
    path: '/starred',
  },
  'Recycle bin': {
    icon: (defaultClasses: string, active: boolean) => (
      <Icon
        icon={`heroicons-${active ? 'solid' : 'outline'}:trash`}
        className={defaultClasses}
      />
    ),
    path: '/bin',
  },
}

const subNavItems = {
  Help: {
    icon: (defaultClasses: string, active: boolean) => (
      <Icon
        icon={`heroicons-outline:question-${'mark'}-circle`}
        className={defaultClasses}
      />
    ),
    path: '/help',
  },
  Settings: {
    icon: (defaultClasses: string, active: boolean) => (
      <Icon
        icon={`heroicons-${active ? 'solid' : 'outline'}:cog`}
        className={defaultClasses}
      />
    ),
    path: '/settings',
  },
}

type SideNavProps = {
  title: string
  className?: string
}

export const SideNav: React.FC<SideNavProps> = ({ title, className = '' }) => {
  const { pathname } = useRouter()

  return (
    <div className={`flex w-72 flex-col ${className}`}>
      <div className="mb-8 flex items-center">
        <div className="mr-4 rounded-xl bg-indigo-600 p-3">
          <Icon icon="mdi:cloud-outline" className="h-6 w-6 text-indigo-50" />
        </div>
        <h1 className="text-2xl font-bold">{title}</h1>
      </div>

      {Object.entries(mainNavItems).map(([name, { icon, path }], idx) => (
        <SideNavItem
          key={name}
          name={name}
          icon={icon}
          path={path}
          active={pathname === path}
          className={idx !== 0 ? 'mt-3' : ''}
        />
      ))}

      {Object.entries(subNavItems).map(([name, { icon, path }], idx) => (
        <SideNavItem
          key={name}
          name={name}
          icon={icon}
          path={path}
          active={pathname === path}
          className={idx !== 0 ? 'mt-3' : 'mt-6'}
        />
      ))}

      <CTA />
    </div>
  )
}

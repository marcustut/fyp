import React from 'react'
import { Icon } from '@iconify/react'
import { TW } from 'tailwindcss-classnames'

import { Button } from '@/components'

type NotificationProps = {
  dot?: boolean
  className?: string
}

export const Notification: React.FC<NotificationProps> = ({
  dot = false,
  className = '',
}) => {
  return (
    <Button
      variant="transparent"
      tw={() => ({
        padding: TW.padding('p-0'),
      })}
      className={`relative ${className}`}
    >
      {dot && (
        <div className="absolute top-1 right-1 h-1.5 w-1.5 rounded-full bg-indigo-400 outline outline-indigo-100" />
      )}
      <Icon
        icon="heroicons-solid:bell"
        className={`h-full w-full text-slate-500 transition duration-150 ease-in-out hover:text-indigo-400`}
      />
    </Button>
  )
}

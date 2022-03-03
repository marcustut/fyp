import React from 'react'
import { Icon } from '@iconify/react'
import { TWidth, THeight } from 'tailwindcss-classnames'

import { Avatar, AvatarProps, Menu } from '@/components'
import { useAuth } from '@/lib/auth'
import { useDarkMode } from '@/hooks'

type AvatarDropdownProps = {
  avatarProps: AvatarProps & { width: TWidth; height: THeight }
}

export const AvatarDropdown: React.FC<AvatarDropdownProps> = ({
  avatarProps,
}) => {
  const { signOut } = useAuth()
  const { darkMode, setDarkMode } = useDarkMode()

  return (
    <Menu
      buttonRender={() => <Avatar {...avatarProps} />}
      buttonProps={{
        className: [avatarProps.width, avatarProps.height].join(' '),
      }}
      items={[
        {
          render: () => <>Switch to {darkMode ? 'Light Mode' : 'Dark Mode'}</>,
          icon: (defaultClasses) => (
            <Icon
              icon={
                darkMode ? 'heroicons-outline:sun' : 'heroicons-outline:moon'
              }
              className={defaultClasses}
            />
          ),
          onClick: () => setDarkMode(!darkMode),
        },
        {
          render: () => <>Log Out</>,
          icon: (defaultClasses) => (
            <Icon icon="heroicons-outline:logout" className={defaultClasses} />
          ),
          onClick: () => signOut(),
        },
      ]}
    />
  )
}

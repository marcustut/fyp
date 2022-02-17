import React, { Fragment } from 'react'
import { Menu as HeadlessMenu, Transition } from '@headlessui/react'
import { TW, classnames } from 'tailwindcss-classnames'

import { TailwindClasses } from '@/types/utils'
import { Button } from '@/components'

type MenuItemProps = {
  render: () => JSX.Element
  icon?: (defaultClasses: string) => JSX.Element
  onClick?: React.MouseEventHandler<HTMLButtonElement>
  tw?: TailwindClasses
}

const MenuItem: React.FC<MenuItemProps> = ({ render, icon, onClick, tw }) => {
  return (
    <HeadlessMenu.Item>
      {({ active }) => (
        <button
          className={classnames(
            // ...tw,
            TW.width('w-full'),
            TW.display('flex'),
            TW.padding('px-4', 'py-3'),
            TW.fontSize('text-sm'),
            TW.borderRadius('rounded-lg'),
            TW.alignItems('items-center'),
            TW.transitionProperty('transition'),
            TW.transitionDuration('duration-150'),
            TW.transitionTimingFunction('ease-in-out'),
            TW.textColor('hover:text-slate-50', {
              'text-slate-900': !active,
            }),
            TW.backgroundColor({
              'bg-transparent': !active,
              'bg-indigo-500': active,
            })
          )}
          onClick={onClick}
        >
          {icon &&
            icon(
              classnames(
                TW.margin('mr-2'),
                TW.height('h-5'),
                TW.width('w-5'),
                TW.textColor({
                  'text-indigo-500': !active,
                  'text-indigo-300': active,
                })
              )
            )}
          {render()}
        </button>
      )}
    </HeadlessMenu.Item>
  )
}

type MenuProps = {
  items: MenuItemProps[]
  buttonRender: () => JSX.Element
  buttonProps?: React.DetailedHTMLProps<
    React.ButtonHTMLAttributes<HTMLButtonElement>,
    HTMLButtonElement
  >
  tw?: TailwindClasses
}

export const Menu: React.FC<MenuProps> = ({
  items,
  buttonRender,
  buttonProps,
  tw,
}) => {
  return (
    <HeadlessMenu
      as="div"
      className={`relative ${tw ? classnames(...Object.values(tw)) : ''}`}
    >
      <HeadlessMenu.Button {...buttonProps}>
        {buttonRender()}
      </HeadlessMenu.Button>
      <Transition
        as={Fragment}
        enter="transition ease-out duration-100"
        enterFrom="transform opacity-0 scale-95"
        enterTo="transform opacity-100 scale-100"
        leave="transition ease-in duration-75"
        leaveFrom="transform opacity-100 scale-100"
        leaveTo="transform opacity-0 scale-95"
      >
        <HeadlessMenu.Items className="absolute right-0 w-56 rounded-md bg-indigo-50 text-slate-900">
          <div className="p-1">
            {items.map((props, idx) => (
              <MenuItem
                key={idx}
                {...props}
                tw={{ margin: TW.margin({ 'mt-1': idx !== 0 }) }}
              />
            ))}
          </div>
        </HeadlessMenu.Items>
      </Transition>
    </HeadlessMenu>
  )
}

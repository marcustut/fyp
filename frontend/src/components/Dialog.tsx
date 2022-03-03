import React, { Fragment } from 'react'
import { Transition, Dialog as HeadlessDialog } from '@headlessui/react'

type DialogProps = {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  render: (
    open: DialogProps['open'],
    setOpen: DialogProps['setOpen']
  ) => JSX.Element
  renderContainer?: (
    open: DialogProps['open'],
    setOpen: DialogProps['setOpen'],
    render: DialogProps['render']
  ) => JSX.Element
}

export const Dialog: React.FC<DialogProps> = ({
  open,
  setOpen,
  render,
  renderContainer,
}) => {
  return (
    <Transition appear show={open} as={Fragment}>
      <HeadlessDialog onClose={() => setOpen(false)}>
        <Transition.Child
          as={HeadlessDialog.Overlay}
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
          className="fixed inset-0 bg-black bg-opacity-30 backdrop-blur-[2px]"
        />

        {renderContainer ? (
          renderContainer(open, setOpen, render)
        ) : (
          <Transition.Child
            as="div"
            enter="ease-out duration-300"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
            className="fixed top-1/2 left-1/2 z-10 flex min-w-[300px] -translate-x-1/2 -translate-y-1/2 transform flex-col items-center rounded-lg bg-slate-50 px-6 py-4 text-black shadow-2xl dark:bg-slate-700 dark:text-slate-100 lg:min-w-[550px]"
          >
            {render(open, setOpen)}
          </Transition.Child>
        )}
      </HeadlessDialog>
    </Transition>
  )
}

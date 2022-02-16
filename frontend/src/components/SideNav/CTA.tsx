import React from 'react'
import { toast } from 'react-toastify'
import { TW } from 'tailwindcss-classnames'

import { Button } from '@/components'

export const CTA: React.FC = () => {
  return (
    <div className="mt-auto flex flex-col items-center justify-center rounded-3xl bg-gradient-to-t from-indigo-100 to-slate-100 p-8 text-center">
      <img src="/images/feedback.png" className="scale-125 transform" />
      <p className="mt-4 font-semibold">We are in Closed Beta</p>
      <p className="my-3 text-sm text-slate-400">
        Feel free to leave your feedback.
      </p>
      <Button
        tw={() => ({
          margin: TW.margin('mt-4'),
          padding: TW.padding('py-3', 'px-12'),
          borderRadius: TW.borderRadius('rounded-3xl'),
          backgroundColor: TW.backgroundColor('bg-indigo-600'),
        })}
        // TODO: Implement this
        onClick={() => toast.warning('not implemented')}
      >
        Feedback
      </Button>
    </div>
  )
}

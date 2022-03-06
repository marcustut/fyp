import React, { useMemo, useState } from 'react'
import { Icon } from '@iconify/react'

import { HalfCircleProgress } from '@/components/HalfCircleProgress'
import { niceBytes } from '@/utils/formatting'
import { useListSlideQuery, UserWithAuth } from '@/generated/graphql'

const maxSize = 102400

type SummaryProps = {
  user: UserWithAuth
  variant: 'detailed' | 'visual'
  className?: string
}

export const Summary: React.FC<SummaryProps> = ({
  user,
  variant,
  className = '',
}) => {
  const [slides] = useListSlideQuery({
    variables: {
      where: { hasUserWith: [{ id: user.user.id }], deleted: false },
    },
  })
  const usage = useMemo(() => {
    if (!slides.data) return []
    const x = slides.data.Slides.edges
      .map((e) => e.node)
      .map((s) => ({ type: 'markdown', size: s.size! }))

    return [
      {
        type: 'markdown',
        totalFiles: x.length,
        totalSize: x.reduce((a, b) => a + b.size, 0),
      },
    ]
  }, [slides])

  const storageUsed = useMemo(
    () => usage.reduce((a, b) => a + b.totalSize, 0),
    [usage]
  )

  const inner = () => {
    switch (variant) {
      case 'detailed': {
        return (
          <>
            {usage.map(({ type, totalFiles, totalSize }, idx) => (
              <div
                key={type}
                className={`flex w-full items-center rounded-3xl border border-slate-300 px-8 py-5 ${
                  idx !== 0 ? 'mt-4' : 'mt-4'
                }`}
              >
                <Icon
                  icon={`heroicons-solid:${
                    type === 'images'
                      ? 'photograph'
                      : type === 'videos'
                      ? 'film'
                      : 'document-text'
                  }`}
                  className="mr-8 h-7 w-7 text-orange-300"
                />
                <div className="flex flex-col">
                  <p className="capitalize">{type}</p>
                  <p className="text-sm text-slate-400">
                    {totalFiles} total files
                  </p>
                </div>
                <div className="ml-auto text-lg font-medium text-slate-500">
                  {niceBytes(totalSize)}
                </div>
              </div>
            ))}
          </>
        )
      }
      case 'visual': {
        return (
          <>
            <h3 className="text-2xl font-medium">Storage</h3>
            <HalfCircleProgress
              className="my-8"
              progress={(storageUsed / maxSize) * 100}
            />
            <p className="text-3xl font-medium">{niceBytes(storageUsed)}</p>
            <p className="mt-1 text-slate-400">
              of {niceBytes(maxSize)} capacity
            </p>
          </>
        )
      }
    }
  }

  return (
    <div className={`flex flex-col items-center justify-center ${className}`}>
      {inner()}
    </div>
  )
}

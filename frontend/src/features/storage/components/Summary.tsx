import React, { useMemo, useState } from 'react'
import { Icon } from '@iconify/react'

import { HalfCircleProgress } from '@/components/HalfCircleProgress'
import { niceBytes } from '@/utils/formatting'

type SummaryProps = {
  variant: 'detailed' | 'visual'
  className?: string
}

export const Summary: React.FC<SummaryProps> = ({
  variant,
  className = '',
}) => {
  const [usage, setUsage] = useState<
    {
      type: 'images' | 'videos' | 'markdown'
      totalFiles: number
      totalSize: number
    }[]
  >([
    { type: 'images', totalFiles: 123, totalSize: 2341234 },
    { type: 'videos', totalFiles: 3, totalSize: 363211234 },
    { type: 'markdown', totalFiles: 10, totalSize: 1231123 },
  ])
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
              progress={(storageUsed / 1000000000) * 100}
            />
            <p className="text-3xl font-medium">{niceBytes(storageUsed)}</p>
            <p className="mt-1 text-slate-400">of 1GB capacity</p>
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

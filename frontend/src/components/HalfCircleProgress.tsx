import React, { useEffect, useRef } from 'react'

type HalfCircleProgressProps = {
  progress: number
  className?: string
}

export const HalfCircleProgress: React.FC<HalfCircleProgressProps> = ({
  progress,
  className = '',
}) => {
  const circleRef = useRef<SVGCircleElement>(null)

  useEffect(() => {
    if (!circleRef.current) return
    const radius = circleRef.current.r.baseVal.value
    const circumference = radius * 2 * Math.PI
    circleRef.current.style.strokeDasharray = `${circumference} ${circumference}`
    circleRef.current.style.strokeDashoffset = `${
      circumference - (progress / 100 / 2) * circumference
    }`
  }, [circleRef, progress])

  return (
    <div
      className={`relative flex h-32 w-64 items-center justify-center rounded-t-full bg-slate-200 ${className}`}
    >
      <div className="absolute bottom-0 h-16 w-32 rounded-t-full bg-slate-200" />
      <svg className="absolute bottom-0 h-full w-full rotate-180 transform">
        <circle
          className="stroke-orange-300/20"
          strokeWidth={50}
          fill="transparent"
          r={90}
          cx="50%"
          cy={0}
        />
      </svg>
      <svg className="h-full w-full rotate-180 transform">
        <circle
          ref={circleRef}
          className="stroke-orange-300"
          strokeWidth={50}
          fill="transparent"
          r={90}
          cx="50%"
          cy={0}
        />
      </svg>
    </div>
  )
}

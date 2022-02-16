import React from 'react'

type AppContainerProps = {
  className?: string
}

export const AppContainer: React.FC<AppContainerProps> = ({
  className = '',
  children,
}) => {
  return (
    <div className={`w-3/5 rounded-3xl bg-indigo-100 p-8 ${className}`}>
      {children}
    </div>
  )
}

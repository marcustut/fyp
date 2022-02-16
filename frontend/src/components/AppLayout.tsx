import React from 'react'

type AppLayoutProps = {
  className?: string
}

export const AppLayout: React.FC<AppLayoutProps> = ({
  className = '',
  children,
}) => (
  <div
    className={`max-w-screen flex h-screen max-h-screen w-screen space-x-8 overflow-hidden p-8 dark:text-white ${className}`}
  >
    {children}
  </div>
)

import React from 'react'
import { classnames } from 'tailwindcss-classnames'

import { TailwindClasses } from '@/types/utils'

export type TableRowProps = React.DetailedHTMLProps<
  React.HTMLAttributes<HTMLTableRowElement>,
  HTMLTableRowElement
> & {
  tw?: TailwindClasses
  className?: string
}

export const TableRow: React.FC<TableRowProps> = ({
  children,
  tw,
  className = '',
  ...props
}) => {
  return (
    <tr
      className={`text-slate-400 dark:text-slate-300 ${className} ${
        tw ? classnames(...Object.values(tw)) : ''
      }`}
      {...props}
    >
      {children}
    </tr>
  )
}

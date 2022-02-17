import React from 'react'
import { classnames } from 'tailwindcss-classnames'

import { TailwindClasses } from '@/types/utils'

export type TableCellProps = React.DetailedHTMLProps<
  React.TdHTMLAttributes<HTMLTableDataCellElement>,
  HTMLTableDataCellElement
> & {
  tw?: TailwindClasses
  className?: string
}

export const TableCell: React.FC<TableCellProps> = ({
  children,
  tw,
  className = '',
  ...props
}) => {
  return (
    <td
      className={`${className} ${tw ? classnames(...Object.values(tw)) : ''}`}
      {...props}
    >
      {children}
    </td>
  )
}

import React from 'react'

type TableHeadCellProps = React.DetailedHTMLProps<
  React.ThHTMLAttributes<HTMLTableHeaderCellElement>,
  HTMLTableHeaderCellElement
> & {
  className?: string
}

export const TableHeadCell: React.FC<TableHeadCellProps> = ({
  children,
  className = '',
  ...props
}) => {
  return (
    <th className={`pb-2 text-left font-medium ${className}`} {...props}>
      {children}
    </th>
  )
}

import React from 'react'

type TableCellProps = React.DetailedHTMLProps<
  React.TdHTMLAttributes<HTMLTableDataCellElement>,
  HTMLTableDataCellElement
> & {
  className?: string
}

export const TableCell: React.FC<TableCellProps> = ({
  children,
  className = '',
  ...props
}) => {
  return (
    <td className={`py-4 ${className}`} {...props}>
      {children}
    </td>
  )
}

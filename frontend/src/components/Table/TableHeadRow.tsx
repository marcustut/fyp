import React from 'react'

type TableHeadRowProps = React.DetailedHTMLProps<
  React.HTMLAttributes<HTMLTableRowElement>,
  HTMLTableRowElement
> & {
  className?: string
}

export const TableHeadRow: React.FC<TableHeadRowProps> = ({
  children,
  className = '',
  ...props
}) => {
  return (
    <tr className={`capitalize text-slate-500 ${className}`} {...props}>
      {children}
    </tr>
  )
}

import React from 'react'

type TableTableProps = React.DetailedHTMLProps<
  React.TableHTMLAttributes<HTMLTableElement>,
  HTMLTableElement
> & {
  className?: string
}

export const TableTable: React.FC<TableTableProps> = ({
  children,
  className = '',
  ...props
}) => {
  return (
    <table className={className} {...props}>
      {children}
    </table>
  )
}

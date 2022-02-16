import React from 'react'

type TableRowProps = React.DetailedHTMLProps<
  React.HTMLAttributes<HTMLTableRowElement>,
  HTMLTableRowElement
> & {
  className?: string
}

export const TableRow: React.FC<TableRowProps> = ({
  children,
  className = '',
  ...props
}) => {
  // const defaultUtilities: WindiUtilities = {
  //   text: 'body-text text-word-active-light dark:text-word-active-dark',
  //   border: 'border-b border-grey dark:border-grey-border',
  // }

  return (
    <tr className={`text-slate-400 ${className}`} {...props}>
      {children}
    </tr>
  )
}

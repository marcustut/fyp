import React, { useMemo, useEffect } from 'react'
import {
  Column,
  Hooks,
  PluginHook,
  useGlobalFilter,
  useSortBy,
  useTable,
} from 'react-table'

import {
  GlobalSearch,
  TableTable,
  TableHead,
  TableHeadRow,
  TableHeadCell,
  TableBody,
  TableRow,
  TableCell,
} from '@/components'

type TableProps<T extends Record<string, unknown>> = {
  data: T[]
  columns: Column<T>[]
  hooks?: (hooks: Hooks<T>) => void
  updateGlobalFilter?: React.Dispatch<
    React.SetStateAction<(filterValue: any) => void>
  >
  searchable?: boolean
  sortable?: boolean
  className?: string
}

export const Table = <T extends Record<string, any>>({
  data,
  columns,
  hooks,
  updateGlobalFilter,
  searchable = true,
  sortable = true,
  className = '',
}: TableProps<T>) => {
  // make sure data and columns are memoized
  const memoizedData = useMemo(() => data, [data])
  const memoizedColumns = useMemo(() => columns, [columns])

  // construct the hooks array from given props
  const pluginHooks = useMemo(() => {
    const pluginHooks: PluginHook<T>[] = []
    searchable && pluginHooks.push(useGlobalFilter)
    hooks && pluginHooks.push(hooks)
    sortable && pluginHooks.push(useSortBy)
    return pluginHooks
  }, [hooks, searchable, sortable])

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    footerGroups,
    rows,
    prepareRow,
    setGlobalFilter,
  } = useTable<T>(
    {
      columns: memoizedColumns,
      data: memoizedData,
    },
    ...pluginHooks
  )

  useEffect(() => {
    if (!updateGlobalFilter) return
    if (setGlobalFilter) updateGlobalFilter(() => setGlobalFilter)
  }, [setGlobalFilter])

  return (
    <>
      {/* {searchable && <GlobalSearch setGlobalFilter={setGlobalFilter} />} */}
      <TableTable className={`${className}`} {...getTableProps()}>
        <TableHead>
          {headerGroups.map((headerGroup, i) => (
            <TableHeadRow {...headerGroup.getHeaderGroupProps()} key={i}>
              {headerGroup.headers.map((column, idx) => (
                <TableHeadCell
                  {...column.getHeaderProps(
                    sortable ? column.getSortByToggleProps() : undefined
                  )}
                  key={column.id}
                >
                  {column.render('Header')}
                  {sortable &&
                    (column.isSorted
                      ? column.isSortedDesc
                        ? ' ▼'
                        : ' ▲'
                      : '')}
                </TableHeadCell>
              ))}
            </TableHeadRow>
          ))}
        </TableHead>
        <TableBody {...getTableBodyProps()}>
          {rows.map((row) => {
            prepareRow(row)
            return (
              <TableRow {...row.getRowProps()} key={row.id}>
                {row.cells.map((cell) => (
                  <TableCell {...cell.getCellProps()} key={cell.column.id}>
                    {cell.render('Cell')}
                  </TableCell>
                ))}
              </TableRow>
            )
          })}
        </TableBody>
        <tfoot>
          {footerGroups.map((group, i) => (
            <tr {...group.getFooterGroupProps()} key={i}>
              {group.headers.map((column) => (
                <td {...column.getFooterProps()} key={column.id}>
                  {column.render('Footer')}
                </td>
              ))}
            </tr>
          ))}
        </tfoot>
      </TableTable>
    </>
  )
}

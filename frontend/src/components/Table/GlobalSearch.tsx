import React from 'react'
import { useAsyncDebounce } from 'react-table'

import { Input, InputProps } from '@/components'

export type GlobalSearchProps = {
  setGlobalFilter: (filterValue: string) => void
  icon?: InputProps['icon']
  placeholder?: string
  inputContainerClassName?: InputProps['className']
  inputClassName?: InputProps['inputClassName']
}

export const GlobalSearch: React.FC<GlobalSearchProps> = ({
  setGlobalFilter,
  icon,
  placeholder = 'Search',
  inputContainerClassName = '',
  inputClassName = '',
}) => {
  const onChange = useAsyncDebounce(
    (event: React.ChangeEvent<HTMLInputElement>) =>
      setGlobalFilter(event.target.value),
    300
  )

  return (
    <Input
      variant={'transparent'}
      icon={icon}
      placeholder={placeholder}
      inputClassName={inputClassName}
      className={inputContainerClassName}
      onChange={onChange}
    />
  )
}

import { Fragment } from 'react'
import { Listbox, Transition } from '@headlessui/react'
import { Icon } from '@iconify/react'

interface SelectionProps<T extends string | number> {
  value: T
  onChange: (value: T) => void
  labelText: string
  buttonText: string
  options: Record<string | number, string>
}

export function Selection<T extends string | number>({
  value,
  onChange,
  labelText,
  buttonText,
  options,
}: SelectionProps<T>) {
  return (
    <Listbox value={value} onChange={onChange}>
      <p className="mt-4 mb-1 self-start text-left text-sm font-medium">
        {labelText}
      </p>
      <Listbox.Button className="relative w-full cursor-default rounded-lg bg-indigo-200 py-2 pl-3 pr-10 text-left focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm">
        <span className="block truncate">{buttonText}</span>
        <span className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
          <Icon
            icon="heroicons-outline:selector"
            className="h-5 w-5 text-gray-700"
            aria-hidden="true"
          />
        </span>
      </Listbox.Button>
      <Transition
        as={Fragment}
        leave="transition ease-in duration-100"
        leaveFrom="opacity-100"
        leaveTo="opacity-0"
      >
        <Listbox.Options className="mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
          {Object.entries(options).map(([optValue, optText], idx) => (
            <Listbox.Option
              key={idx}
              className={({ active }) =>
                `relative cursor-default select-none py-2 pl-10 pr-4 ${
                  active ? 'bg-indigo-100 text-indigo-700' : 'text-gray-900'
                }`
              }
              value={optValue}
            >
              {({ selected }) => (
                <>
                  <span
                    className={`block truncate ${
                      selected ? 'font-medium' : 'font-normal'
                    }`}
                  >
                    {optText}
                  </span>
                  {selected ? (
                    <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-indigo-600">
                      <Icon
                        icon="heroicons-outline:check"
                        className="h-5 w-5"
                        aria-hidden="true"
                      />
                    </span>
                  ) : null}
                </>
              )}
            </Listbox.Option>
          ))}
        </Listbox.Options>
      </Transition>
    </Listbox>
  )
}

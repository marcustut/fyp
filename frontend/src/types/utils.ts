import { TW, TTailwindString } from 'tailwindcss-classnames'

export type MapOptional<Type> = {
  [Property in keyof Type]+?: Type[Property]
}

export type MapToNewType<Type, NewType> = {
  [Property in keyof Type]+?: NewType
}

export type TailwindClasses = MapOptional<
  MapToNewType<typeof TW, TTailwindString>
>

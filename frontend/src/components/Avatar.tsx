import React from 'react'
import { TWidth, THeight } from 'tailwindcss-classnames'

export type AvatarProps = {
  src?: string | null
  gender: 'male' | 'female'
  name: string
  width: TWidth
  height: THeight
  outline?: boolean
  rounded?: boolean
  className?: string
}

export const Avatar: React.FC<AvatarProps> = ({
  src,
  gender,
  name,
  outline = false,
  rounded = true,
  className = '',
}) => {
  return (
    <img
      src={
        src
          ? src
          : `https://avatars.dicebear.com/api/${gender}/${name}.svg?mood[]=happy`
      }
      alt={`${name}'s avatar`}
      className={`object-cover ${
        outline ? 'outline outline-offset-2 outline-indigo-400' : ''
      } ${rounded ? 'rounded-full' : ''} ${className}`}
    />
  )
}

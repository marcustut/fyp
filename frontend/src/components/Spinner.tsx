import { Icon } from '@iconify/react'

type SpinnerProps = {
  spin?: boolean
  icon?: string
  className?: string
}

export const Spinner: React.FC<SpinnerProps> = ({
  spin = true,
  icon = 'mdi:loading',
  className = '',
}) => (
  <Icon
    icon={icon}
    className={`mx-auto h-5 w-5 ${spin ? 'animate-spin' : ''} ${className}`}
  />
)

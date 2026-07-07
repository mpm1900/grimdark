import { STAT_ICONS } from '#/icons/stats'
import type { Stat } from '#/lib/game/core'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'
import type { IconType } from 'react-icons/lib'

const statVariants = cva('', {
  variants: {
    stat: {
      health: '',
      speed: '',
      melee: 'text-melee',
      ranged: 'text-ranged',
      special: 'text-special',
      'martial-defense': 'text-martial',
      'special-defense': 'text-special',
      accuracy: '',
      evasion: '',
      'critical-chance': '',
      'critical-damage': '',
      'effect-chance': '',
      range: '',
    },
  },
  defaultVariants: {
    stat: 'health',
  },
})

function StatName({
  stat,
  children,
  className,
  hideIcon,
  ...props
}: React.ComponentProps<'span'> & {
  stat: Stat
  hideIcon?: boolean
}) {
  const Icon = STAT_ICONS[stat]
  return (
    <span
      {...props}
      className={cn(
        statVariants({ stat }),
        'flex items-center gap-1',
        className
      )}
    >
      {Icon && !hideIcon && <Icon className="size-5" />}
      {children}
    </span>
  )
}

function StatIcon({
  stat,
  className,
  ...props
}: React.ComponentProps<IconType> & { stat: Stat }) {
  const Icon = STAT_ICONS[stat]
  return <Icon className={cn(statVariants({ stat }), className)} {...props} />
}

export { StatName, StatIcon, statVariants }

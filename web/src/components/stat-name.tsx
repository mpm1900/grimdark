import { STAT_ICONS } from '#/icons/stats'
import type { Stat } from '#/lib/game/core'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'

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
      {Icon && !hideIcon && <Icon className="text-foreground/40 size-5" />}
      {children}
    </span>
  )
}

export { StatName }

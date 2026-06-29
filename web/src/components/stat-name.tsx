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
  ...props
}: React.ComponentProps<'span'> & { stat: Stat }) {
  return (
    <span {...props} className={cn(statVariants({ stat }), className)}>
      {children}
    </span>
  )
}

export { StatName }

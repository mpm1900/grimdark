import type { Actor } from '#/lib/game/actor'
import type { Stat } from '#/lib/game/core'
import { cn } from '#/lib/utils'

function StatValue({
  actor,
  stat,
  className,
  ...props
}: React.ComponentProps<'span'> & { actor: Actor; stat: Stat }) {
  const value = actor.stats[stat]
  const unmodified = actor.unmodified_stats[stat]
  return (
    <span
      {...props}
      className={cn(
        {
          'text-green-400': value > unmodified,
          'text-red-400': value < unmodified,
        },
        className
      )}
    >
      {value}
    </span>
  )
}

export { StatValue }

import type { Actor } from '#/lib/game/actor'
import type { Stat } from '#/lib/game/core'
import { cn } from '#/lib/utils'

function StatValue({
  actor,
  stat,
  className,
  children,
  map,
  ...props
}: React.ComponentProps<'span'> & {
  actor: Actor
  stat: Stat
  map?: (v: number) => string
}) {
  const value = actor.stats[stat]
  const unmodified = actor.unmodified_stats[stat]
  return (
    <span
      {...props}
      className={cn(
        {
          'text-positive': value > unmodified,
          'text-negative': value < unmodified,
        },
        className
      )}
    >
      {children}
    </span>
  )
}

export { StatValue }

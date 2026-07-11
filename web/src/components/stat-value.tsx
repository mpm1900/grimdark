import type { Actor } from '#/lib/game/actor'
import type { Stat } from '#/lib/game/core'
import { DNumber } from './dnumber'

function StatValue({
  actor,
  stat,
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
    <DNumber value={value} r={unmodified} {...props}>
      {children}
    </DNumber>
  )
}

export { StatValue }

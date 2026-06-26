import type { Actor } from '#/lib/game/actor'
import {
  getBaseAffinityDamage,
  getBaseAffinityResistance,
  type Affinity,
} from '#/lib/game/core'
import { cn } from '#/lib/utils'

function AffinityResistanceValue({
  actor,
  affinity,
  className,
  ...props
}: React.ComponentProps<'span'> & { actor: Actor; affinity: Affinity }) {
  const value = actor.affinity_resistance[affinity]
  const unmodified = getBaseAffinityResistance(actor, affinity)

  return (
    <span
      className={cn({
        'text-green-400': value > unmodified,
        'text-red-400': value < unmodified,
      })}
      {...props}
    >
      {value}
    </span>
  )
}

function AffinityDamageValue({
  actor,
  affinity,
  className,
  ...props
}: React.ComponentProps<'span'> & { actor: Actor; affinity: Affinity }) {
  const value = actor.affinity_damage[affinity] ?? 0
  const unmodified = getBaseAffinityDamage(actor, affinity)
  if (value == 0) return null

  return (
    <span
      className={cn({
        'text-green-400': value > unmodified,
        'text-red-400': value < unmodified,
      })}
      {...props}
    >
      {value}
    </span>
  )
}

export { AffinityResistanceValue, AffinityDamageValue }

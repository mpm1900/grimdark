import type { Actor } from '#/lib/game/actor'
import {
  getBaseAffinityDamage,
  getBaseAffinityResistance,
  mapStage,
  type Affinity,
} from '#/lib/game/core'
import { cn } from '#/lib/utils'

function AffinityResistanceValue({
  actor,
  affinity,
  className,
  ...props
}: React.ComponentProps<'span'> & { actor: Actor; affinity: Affinity }) {
  const immunity = actor.affinity_immunities[affinity]
  const value = actor.affinity_resistance[affinity]
  const unmodified = getBaseAffinityResistance(actor, affinity)

  return (
    <span
      className={cn({
        'text-positive': value > unmodified,
        'text-negative': value < unmodified,
        'text-amber-400': immunity !== undefined,
      })}
      {...props}
    >
      {immunity !== undefined ? '∞' : value}
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
        'text-positive': value > unmodified,
        'text-negative': value < unmodified,
      })}
      {...props}
    >
      {value}
    </span>
  )
}

function AffinityResistanceMultiplier({
  actor,
  affinity,
  value,
  className,
  ...props
}: React.ComponentProps<'span'> & {
  actor: Actor
  affinity: Affinity
  value: number
}) {
  const immunity = actor.affinity_immunities[affinity]
  const mult = mapStage(value, 2, 1)
  return (
    <span
      {...props}
      className={cn(mult === 1 && !immunity && 'text-foreground/20', className)}
    >
      {immunity !== undefined ? (
        <span
          className={cn({
            'text-positive': immunity > 0,
            'text-amber-400': immunity <= 0,
          })}
        >
          x{immunity.toFixed(2)}
        </span>
      ) : (
        <span>x{mult.toFixed(2)}</span>
      )}
    </span>
  )
}

function AffinityDamageMultiplier({
  actor,
  affinity,
  className,
  ...props
}: React.ComponentProps<'span'> & {
  actor: Actor
  affinity: Affinity
}) {
  const mult = mapStage(actor.affinity_damage[affinity], 2, 1)
  return (
    <span
      {...props}
      className={cn({
        'text-foreground/20': mult === 1,
      })}
    >
      x{mult.toFixed(2)}
    </span>
  )
}

export {
  AffinityResistanceValue,
  AffinityDamageValue,
  AffinityResistanceMultiplier,
  AffinityDamageMultiplier,
}

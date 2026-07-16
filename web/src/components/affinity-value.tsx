import type { Actor } from '#/lib/game/actor'
import {
  getBaseAffinityDamage,
  getBaseAffinityResistance,
  mapStage,
  type Affinity,
} from '#/lib/game/core'
import { cn } from '#/lib/utils'
import { DNumber } from './dnumber'

function AffinityResistanceValue({
  actor,
  affinity,
  ...props
}: React.ComponentProps<'span'> & { actor: Actor; affinity: Affinity }) {
  const immunity = actor.affinity_immunities[affinity]
  const value = actor.affinity_resistance[affinity]
  const unmodified = getBaseAffinityResistance(actor, affinity)

  return (
    <DNumber
      value={value}
      r={unmodified}
      perfect={immunity !== undefined}
      {...props}
    >
      {immunity !== undefined ? '∞' : value}
    </DNumber>
  )
}

function AffinityDamageValue({
  actor,
  affinity,
  ...props
}: React.ComponentProps<'span'> & { actor: Actor; affinity: Affinity }) {
  const value = actor.affinity_damage[affinity] ?? 0
  const unmodified = getBaseAffinityDamage(actor, affinity)
  if (value == 0) return null

  return (
    <DNumber value={value} r={unmodified} {...props}>
      {value}
    </DNumber>
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
        <DNumber value={immunity} perfect={immunity <= 0}>
          x{immunity.toFixed(2)}
        </DNumber>
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
  const mult = mapStage(actor.affinity_damage[affinity] ?? 0, 2, 1)
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

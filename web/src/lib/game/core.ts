import type { Actor } from './actor'
import type { Context } from './context'

export type ID = string

export type Affinity =
  | 'arcane'
  | 'cryo'
  | 'fire'
  | 'kinetic'
  | 'lightning'
  | 'poison'
  | 'psychic'

export type Stat =
  | 'health'
  | 'speed'
  | 'melee'
  | 'ranged'
  | 'special'
  | 'martial-defense'
  | 'special-defense'
  | 'accuracy'
  | 'evasion'

export type Phase = 'init' | 'start' | 'main' | 'end' | 'cleanup'

export type Status = 'idle' | 'running' | 'waiting'

export type Bindable<T> = {
  ID: ID
  context: Context
  payload: T
}

export const STATS: Stat[] = [
  'health',
  'melee',
  'ranged',
  'special',
  'martial-defense',
  'special-defense',
  'speed',
  'accuracy',
  'evasion',
]

export const AFFINITIES: Affinity[] = [
  'arcane',
  'cryo',
  'fire',
  'kinetic',
  'lightning',
  'poison',
  'psychic',
]

export const AFFINITY_MATRIX: Record<
  Affinity,
  Partial<Record<Affinity, number>>
> = {
  arcane: {
    arcane: 2,
    fire: 2,
    lightning: 2,
    poison: -2,
    psychic: -2,
  },
  cryo: {
    cryo: -2,
    fire: -2,
    kinetic: -2,
    lightning: 2,
    poison: 2,
  },
  fire: {
    arcane: -2,
    cryo: 2,
    fire: -2,
    poison: 2,
    psychic: -2,
  },
  kinetic: {
    cryo: 2,
    lightning: -2,
    poison: 2,
    psychic: -2,
  },
  lightning: {
    arcane: -2,
    cryo: -2,
    kinetic: 2,
    lightning: -2,
    psychic: 2,
  },
  poison: {
    arcane: 2,
    cryo: -2,
    fire: 2,
    kinetic: -2,
    poison: -2,
  },
  psychic: {
    arcane: 2,
    fire: 2,
    lightning: 2,
    poison: -2,
    psychic: -2,
  },
}

function getBaseAffinityResistance(
  actor: Actor,
  target_affinity: Affinity
): number {
  return (
    -1 *
    actor.affinities.reduce((result, affinity) => {
      const stage = AFFINITY_MATRIX[target_affinity][affinity]
      if (stage == undefined) {
        return result
      }

      return result + stage
    }, 0)
  )
}

function getBaseAffinityDamage(
  actor: Actor,
  target_affinity: Affinity
): number {
  return actor.affinities.includes(target_affinity) ? 1 : 0
}

export { getBaseAffinityResistance, getBaseAffinityDamage }

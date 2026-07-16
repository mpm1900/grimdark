import type { Actor } from './actor'
import type { Context } from './context'

export type ID = string

function isIdNull(id: ID | null | undefined): boolean {
  if (!id) return true
  return id === '00000000-0000-0000-0000-000000000000'
}

export type Affinity =
  'arcane' | 'cryo' | 'fire' | 'kinetic' | 'lightning' | 'poison' | 'psychic'

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
  | 'actions'
  | 'critical-chance'
  | 'critical-damage'
  | 'damage-reflect'
  | 'effect-chance'

export type Stack = 'wounds'

export type Phase = 'init' | 'start' | 'main' | 'end' | 'cleanup'

export type Status = 'idle' | 'running' | 'waiting'

export type WeaponType = 'sword' | 'big-sword' | 'pistol'

export type Bindable<T> = {
  ID: ID
  context: Context
  payload: T
}

export const ALL_STATS: Stat[] = [
  'health',
  'melee',
  'ranged',
  'special',
  'speed',
  'martial-defense',
  'special-defense',
  'accuracy',
  'evasion',
  'critical-chance',
  'critical-damage',
  'damage-reflect',
]

export const MAIN_STATS: Stat[] = [
  'speed',
  'ranged',
  'melee',
  'special',
  'martial-defense',
  'special-defense',
]
export const CLASS_STATS: Stat[] = [
  'health',
  'speed',
  'ranged',
  'melee',
  'special',
  'martial-defense',
  'special-defense',
]

export const ACCURACY_STATS: Stat[] = ['accuracy', 'evasion']

export const CRITICAL_STATS: Stat[] = ['critical-chance', 'critical-damage']

export const STAT_LABELS: Record<Stat, string> = {
  health: 'Health',
  melee: 'Melee',
  ranged: 'Ranged',
  special: 'Special',
  'martial-defense': 'Martial Def.',
  'special-defense': 'Special Def.',
  speed: 'Speed',
  accuracy: 'Accuracy',
  evasion: 'Evasion',
  actions: 'Actions per Turn',
  'critical-chance': 'Critical Chance',
  'critical-damage': 'Critical Damage',
  'damage-reflect': 'Damage Reflect',
  'effect-chance': 'Effect Chance',
}
export const STAT_SLUGS: Record<Stat, string> = {
  health: 'Hp',
  melee: 'Mel',
  ranged: 'Rgd',
  special: 'Spc',
  'martial-defense': 'M.Def',
  'special-defense': 'S.Def',
  speed: 'Spe',
  accuracy: 'Acc',
  evasion: 'Eva',
  actions: 'AP',
  'critical-chance': 'Crit.%',
  'critical-damage': 'Crit.Dmg',
  'damage-reflect': 'Dmg.Rflct',
  'effect-chance': 'Eff.%+',
}

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
    kinetic: 2,
    lightning: 2,
    poison: -2,
    psychic: -2,
  },
}

function getBaseAffinityResistance(
  actor: Actor,
  target_affinity: Affinity
): number {
  const stage = actor.affinities.reduce((result, affinity) => {
    const stage = AFFINITY_MATRIX[target_affinity][affinity]
    if (stage == undefined) {
      return result
    }

    return result + stage
  }, 0)

  return stage === 0 ? 0 : -stage
}

function getBaseAffinityDamage(
  actor: Actor,
  target_affinity: Affinity
): number {
  return actor.affinities.includes(target_affinity) ? 1 : 0
}

function mapStage(stage: number, mod: number, input: number) {
  if (stage > 0) {
    return (input * (stage + mod)) / mod
  }
  if (stage < 0) {
    return (input * mod) / (mod - stage)
  }
  return input
}

export { isIdNull, getBaseAffinityResistance, getBaseAffinityDamage, mapStage }

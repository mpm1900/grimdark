import type { Context } from "./context"

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
  ID: ID,
  context: Context,
  payload: T
}

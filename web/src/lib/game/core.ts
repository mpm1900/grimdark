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
  | 'range'
  | 'special'
  | 'martial-defense'
  | 'special-defense'
  | 'accuracy'
  | 'evasion'

export type Phase = 'init' | 'start' | 'main' | 'end' | 'cleanup'

export type Status = 'idle' | 'running' | 'waiting'

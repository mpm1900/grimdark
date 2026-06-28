import type { Affinity, ID, Stat } from './core'

export type ActionConfig = {
  accuracy: number | null
  affinity: Affinity
  crit_chance: number
  crit_modifier: number
  description: string
  hits: number
  lifesteal: number
  name: string
  power: number
  priority: number
  recoil: number
  stat: Stat | null
}

export type Action = {
  ID: ID
  config: ActionConfig
  is_disabled: boolean
}

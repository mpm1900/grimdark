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
  range: number | null
  recoil: number
  stat: Stat | null
  target_count: number
}

export type Action = {
  ID: ID
  type: string
  subtype: string
  config: ActionConfig
  cooldown: number
  is_disabled: boolean
}

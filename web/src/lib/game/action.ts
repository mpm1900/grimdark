import type { Affinity, Entity, ID, Stat } from './core'

export type ActionConfig = {
  accuracy: number | null
  affinity: Affinity
  cooldown: number
  crit_stage: number
  crit_chance: number
  crit_modifier: number
  repeats: number
  lifesteal: number
  power: number
  priority: number
  range: number | null
  recoil: number
  stat: Stat | null
  target_count: number
  uses: number | null
}

export type Action = {
  ID: ID
  config: ActionConfig
  cooldown: number
  entity: Entity
  is_disabled: boolean
  tags: Array<string>
  uncancelable: boolean
  uses: number
}

export function getActionPower(action: Action): number {
  let power_rating = action.config.power ?? 0
  power_rating *= action.config.accuracy ?? 1
  power_rating *= 1 + action.config.crit_chance
  power_rating *= 1 + action.config.crit_modifier
  power_rating *= action.config.target_count
  if (action.config.target_count > 1) {
    power_rating *= 0.75
  }
  power_rating *= action.config.repeats + 1
  return power_rating
}

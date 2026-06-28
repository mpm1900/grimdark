import type { Action } from './action'
import type { Affinity, ID, Stat } from './core'
import type { Effect } from './effect'
import type { Weapon } from './weapon'

export type Actor = {
  ID: ID
  actions: Array<Action>
  active_modifiers: Array<ID>
  affinities: Array<Affinity>
  affinity_damage: Record<Affinity, number>
  affinity_resistance: Record<Affinity, number>
  augment: string
  effects: Array<Effect>
  is_active: boolean
  is_alive: boolean
  is_protected: boolean
  is_staggered: boolean
  is_stunned: boolean
  level: number
  name: string
  player_ID: ID
  position_ID: ID | null
  stages: Record<Stat, number>
  state: string
  stats: Record<Stat, number>
  status: string
  unmodified_stats: Record<Stat, number>
  weapon: Weapon | null
  wounds: number
}

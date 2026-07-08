import type { Action } from './action'
import type { Affinity, ID, Stat } from './core'
import type { Effect } from './effect'
import type { HeldItem, Weapon } from './weapon'

export type Actor = {
  ID: ID
  actions: Array<Action>
  active_modifiers: Array<ID>
  affinities: Array<Affinity>
  affinity_damage: Record<Affinity, number>
  affinity_resistance: Record<Affinity, number>
  affinity_immunities: Partial<Record<Affinity, number>>
  effects: Array<Effect>
  faction: string
  is_active: boolean
  is_bulwark: boolean
  is_alive: boolean
  is_hidden: boolean
  is_insulated: boolean
  is_player: boolean
  is_protected: boolean
  is_stunned: boolean
  item: HeldItem | null
  level: number
  name: string
  player_ID: ID
  position_ID: ID | null
  race: string
  sprite_url: string
  stages: Record<Stat, number>
  state: string
  stats: Record<Stat, number>
  status: string
  unmodified_stats: Record<Stat, number>
  weapon_l: Weapon | null
  weapon_r: Weapon | null
  wounds: number
}

import type { Affinity, ID, Stat } from './core'

export type Actor = {
  ID: ID
  name: string
  level: number
  affinities: Array<Affinity>
  affinity_damage: Record<Affinity, number>
  affinity_resistance: Record<Affinity, number>
  applied_modifiers: Array<ID>
  damage: number
  is_active: boolean
  is_alive: boolean
  is_protected: boolean
  is_staggered: boolean
  player_ID: ID
  position_ID: ID | null
  stats: Record<Stat, number>
  stages: Record<Stat, number>
  status: string
  unmodified_stats: Record<Stat, number>
}

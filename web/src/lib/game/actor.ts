import type { Affinity, ID, Stat } from './core'

export type Actor = {
  ID: ID
  affinities: Array<Affinity>
  affinity_damage: Record<Affinity, number>
  affinity_resistance: Record<Affinity, number>
  applied_effects: Record<ID, number>
  damage: number
  is_alive: boolean
  is_protected: boolean
  player_ID: ID
  position_ID: ID | null
  stats: Record<Stat, number>
  status: string
  unmodified_stats: Record<Stat, number>
}

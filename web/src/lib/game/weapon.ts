import type { Action } from './action'
import type { ID, Stat } from './core'
import type { Effect } from './effect'

export type Weapon = {
  ID: ID
  actions: Action[]
  aux_stats: Partial<Record<Stat, number>>
  description: string
  effects: Effect[]
  hands: number
  name: string
  weapon_type: string
}

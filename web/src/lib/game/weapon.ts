import type { Action } from './action'
import type { ID, Stat } from './core'
import type { Effect } from './effect'

export type Item = {
  ID: ID
  name: string
  description: string
  effects: Effect[]
}

export type Weapon = Item & {
  actions: Action[]
  aux_stats: Partial<Record<Stat, number>>
  hands: number
  weapon_type: string
}

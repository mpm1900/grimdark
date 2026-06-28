import type { Action } from './action'
import type { ID, Stat } from './core'
import type { Effect } from './effect'

export type Weapon = {
  ID: ID
  actions: Action[]
  aux_stats: Partial<Record<Stat, number>>
  effects: Effect[]
  name: string
}

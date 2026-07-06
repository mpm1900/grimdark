import type { Action } from './action'
import type { ID, Stat } from './core'
import type { Effect } from './effect'

export type Item = {
  ID: ID
  name: string
  description: string
}

export type Weapon = Item & {
  actions: Action[]
  aux_stats: Partial<Record<Stat, number>>
  effects: Effect[]
  hands: number
  weapon_type: string
}

export type HeldItem = Item & {
  effects: Effect[]
}

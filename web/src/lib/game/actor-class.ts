import type { Action } from './action'
import type { Affinity, ID, Stat } from './core'
import type { Effect } from './effect'
import type { Item, Weapon } from './weapon'

export type ActorClass = {
  ID: ID
  actions: Action[]
  affinities: Affinity[]
  effects: Effect[]
  faction: string
  name: string
  options: {
    weapons: Weapon[]
    items: Item[]
  }
  race: string
  sprite_url: string
  stats: Record<Stat, number>
}

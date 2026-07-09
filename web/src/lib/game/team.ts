import type { ID } from './core'

export type ActorConfig = {
  class: ID | null
  items: ID[]
  name: string
  weapon_l: ID | null
  weapon_r: ID | null
}

export type TeamConfig = {
  actors: ActorConfig[]
}

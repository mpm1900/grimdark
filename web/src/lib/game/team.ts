import type { User } from '../queries/auth'
import type { ID } from './core'

export type ActorConfig = {
  class: ID | null
  items: ID[]
  name: string
  weapon_l: ID | null
  weapon_r: ID | null
}

export type TeamConfig = {
  user: User
  actors: ActorConfig[]
}

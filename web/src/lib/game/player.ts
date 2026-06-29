import type { User } from '../queries/auth'
import type { ID } from './core'

export type Player = {
  ID: ID
  positions: Array<{
    ID: ID
    actor_ID: ID
  }>
  user: User
}

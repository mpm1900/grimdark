import type { User } from '../queries/auth'
import type { ID } from './core'

export type Player = {
  ID: ID
  positions: Record<ID, ID>
  user: User
}

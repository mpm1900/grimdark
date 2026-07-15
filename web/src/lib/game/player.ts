import type { User } from '../queries/auth'
import type { ID } from './core'

export type Player = {
  ID: ID
  user: User
  actor_count: number
  ready: boolean
}

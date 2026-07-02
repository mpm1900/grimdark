import type { User } from '../queries/auth'
import type { ID } from './core'

export type Position = {
  ID: ID
  actor_ID: ID
  player_ID: ID
  rank: number
}

export type Player = {
  ID: ID
  positions: Array<Position>
  user: User
}

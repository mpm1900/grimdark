import type { Lobby } from '../stores/clients'
import type { ID } from './core'

export type Instance = {
  ID: ID
  status: string
  lobby: Lobby
}

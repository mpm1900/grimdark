import { Store } from '@tanstack/store'
import type { ID } from '../game/core'
import type { User } from '../queries/auth'

type Client = {
  ID: string
  role: string
  user: User
}

type Lobby = {
  client: Client | null
  players: Array<Client>
  spectators: Array<Client>
  ready: Record<ID, boolean>
}

const lobbyStore = new Store<Lobby>({
  client: null,
  players: [],
  spectators: [],
  ready: {},
})

export { lobbyStore }
export type { Client, Lobby }

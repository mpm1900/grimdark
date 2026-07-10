import { Store } from '@tanstack/store'

type Client = {
  ID: string
  role: string
}

type Lobby = {
  client: Client | null
  players: Array<Client>
  spectators: Array<Client>
}

const lobbyStore = new Store<Lobby>({
  client: null,
  players: [],
  spectators: [],
})

export { lobbyStore }
export type { Client, Lobby }

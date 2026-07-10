import { Store } from '@tanstack/store'

type Client = {
  ID: string
}

type Lobby = {
  client: Client | null
  players: Array<Client>
}

const lobbyStore = new Store<Lobby>({
  client: null,
  players: [],
})

export { lobbyStore }
export type { Client, Lobby }

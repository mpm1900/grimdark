import { Store } from '@tanstack/store'

type Client = {
  ID: string
}

type Lobby = {
  client: Client | null
  clients: Array<Client>
}

const lobbyStore = new Store<Lobby>({
  client: null,
  clients: [],
})

export { lobbyStore }
export type { Client, Lobby }

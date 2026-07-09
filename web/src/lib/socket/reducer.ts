import { lobbyStore } from '../stores/clients'
import { gameStore } from '../stores/game'
import type { SocketResponse } from './request'

function socket_reducer(message: SocketResponse | null) {
  if (!message?.type) return
  switch (message.type) {
    case 'game': {
      if (message.game) {
        gameStore.setState(() => message.game!)
      }
      return
    }
    case 'lobby': {
      lobbyStore.setState((c) => ({
        ...c,
        clients: message.lobby?.clients!,
      }))
      return
    }
    case 'on-connect': {
      console.log('on-connect', message)
      lobbyStore.setState((c) => ({
        ...c,
        client: message.lobby?.client!,
      }))
      return
    }
  }
}

export { socket_reducer }

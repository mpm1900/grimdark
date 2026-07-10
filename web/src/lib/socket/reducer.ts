import { lobbyStore } from '../stores/clients'
import { gameStore } from '../stores/game'
import type { SocketResponse } from './request'

function socket_reducer(message: SocketResponse | null) {
  if (!message?.type) return
  switch (message.type) {
    case 'game': {
      if (message.game) {
        console.log(message.game)
        gameStore.setState(() => message.game!)
      }
      return
    }
    case 'lobby': {
      lobbyStore.setState((c) => ({
        ...c,
        players: message.lobby?.players!,
        spectators: message.lobby?.spectators!,
      }))
      return
    }
    case 'on-connect': {
      console.log('on connect', message)
      lobbyStore.setState((c) => ({
        ...c,
        client: message.lobby?.client!,
        players: message.lobby?.players!,
        spectators: message.lobby?.spectators!,
      }))
      return
    }
  }
}

export { socket_reducer }

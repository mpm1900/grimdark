import { clientsStore } from '../stores/clients'
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
    case 'clients': {
      clientsStore.setState((c) => ({
        ...c,
        clients: message.clients!,
      }))
      return
    }
    case 'join-success': {
      if (message.game) {
        gameStore.setState(() => message.game!)
      }
      clientsStore.setState((c) => ({
        ...c,
        me: message.clients![0],
      }))
      return
    }
  }
}

export { socket_reducer }

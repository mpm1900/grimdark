import type { Context } from '../game/context'
import type { Game } from '../game/game'
import type { Client } from '../stores/clients'

type SocketRequestType =
  | 'set-team'
  | 'ready-team'
  | 'cancel-team'
  | 'start-battle'
  | 'reset'
  | 'push-action'
  | 'remove-action'
  | 'run-game-actions'
  | 'resolve-prompt'
  | 'validate-context'
  | 'get-targets'

type SocketRequest = {
  type: SocketRequestType
  client_ID: string
  context: Context
}

type SocketResponse = {
  type: 'game' | 'clients' | 'join-success' | 'validate-context' | 'target-IDs'
  game: Game | null
  clients: Array<Client> | null
  context: Context | null
  valid: boolean | null
}

type SocketMessageSubscriber = (
  event: MessageEvent,
  message: SocketResponse | null
) => void

export type {
  SocketMessageSubscriber,
  SocketRequest,
  SocketRequestType,
  SocketResponse,
}

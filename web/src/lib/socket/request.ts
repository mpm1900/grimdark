import type { Context } from '../game/context'
import type { Game } from '../game/game'
import type { TeamConfig } from '../game/team'
import type { Client } from '../stores/clients'

type SocketRequestType =
  | 'load-team'
  | 'push-action'
  | 'cancel-action'
  | 'run-game-actions'
  | 'resolve-prompt'
  | 'validate-context'
  | 'get-targets'

type SocketRequest = {
  type: SocketRequestType
  client_ID: string
  context: Context
  team_config?: TeamConfig
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

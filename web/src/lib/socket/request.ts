import type { Context } from '../game/context'
import type { Game } from '../game/game'
import type { TeamConfig } from '../game/team'
import type { Lobby } from '../stores/clients'

type SocketRequestType =
  | 'post-connect'
  | 'ready'
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
  type:
    | 'on-connect'
    | 'post-connect'
    | 'game-start'
    | 'game'
    | 'lobby'
    | 'validate-context'
    | 'target-IDs'
  game: Game | null
  lobby: Lobby | null
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

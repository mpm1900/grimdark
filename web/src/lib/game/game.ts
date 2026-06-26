import type { Actor } from './actor'
import type { Context } from './context'
import type { Phase, Status } from './core'

export type Game = {
  active_context: Context | null
  actors: Actor[]
  logs: any[]
  phase: Phase
  players: any[]
  status: Status
  turn: number
}

import type { Actor } from './actor'
import type { Context } from './context'
import type { Bindable, Phase, Status } from './core'
import type { Modifier } from './effect'
import type { Log } from './log'
import type { Player } from './player'

export type Game = {
  active_context: Context | null
  actors: Actor[]
  logs: Bindable<Log>[]
  modifiers: Modifier[]
  phase: Phase
  players: Player[]
  status: Status
  turn: number
}

export function getAppliedModifiers(game: Game, actor: Actor) {
  return Array.from(
    new Set(
      game.modifiers.filter((modifier) => {
        return (
          actor.applied_modifiers[modifier.ID] != undefined &&
          !!modifier.payload.name
        )
      })
    )
  )
}

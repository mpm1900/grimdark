import type { Action } from './action'
import type { Actor } from './actor'
import type { Context } from './context'
import type { Bindable, ID, Phase, Status } from './core'
import type { Effect, Modifier } from './effect'
import type { Log } from './log'
import type { Player } from './player'
import type { Position } from './position'

export type Game = {
  active_context: Context | null
  actors: Actor[]
  commands: Bindable<Action>[]
  logs: Bindable<Log>[]
  modifiers: Modifier[]
  phase: Phase
  positions: Array<Position>
  player_ID: ID
  players: Player[]
  prompts: Bindable<Action>[]
  status: Status
  turn: number
}

export function getAppliedEffects(
  game: Game,
  actor: Actor
): (Effect & { count: number })[] {
  let effect_ids: Record<ID, number> = {}
  let effects: Effect[] = []

  actor.active_modifiers.map((modifier_id) => {
    const modifier = game.modifiers.find((m) => m.ID === modifier_id)
    if (modifier && modifier.payload.name) {
      const count = effect_ids[modifier.payload.ID]
      if (!count) {
        effect_ids[modifier.payload.ID] = 1
        effects = [...effects, modifier.payload]
      } else {
        effect_ids[modifier.payload.ID] = count + 1
      }
    }
  })

  return effects
    .map((e) => ({
      ...e,
      count: effect_ids[e.ID] ?? 0,
    }))
    .sort((a, b) => a.name.localeCompare(b.name))
}

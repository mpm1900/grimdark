import { lobbyStore } from '../stores/clients'
import { gameStore } from '../stores/game'
import type { Actor } from '../game/actor'
import type { Game } from '../game/game'
import type { SocketResponse } from './request'

function mergeActors(current: Actor[], updates: Actor[]): Actor[] {
  const updatesByID = new Map(updates.map((actor) => [actor.ID, actor]))
  const currentIDs = new Set(current.map((actor) => actor.ID))
  const merged = current.map((actor) => updatesByID.get(actor.ID) ?? actor)

  for (const actor of updates) {
    if (!currentIDs.has(actor.ID)) {
      merged.push(actor)
    }
  }

  return merged
}

function applyGamePatch(current: Game, patch: Game): Game {
  return {
    ...current,
    active_context: patch.active_context,
    actors: patch.actors
      ? mergeActors(current.actors, patch.actors)
      : current.actors,
    commands: patch.commands ?? current.commands,
    instance_ID: patch.instance_ID ?? current.instance_ID,
    logs: patch.logs ? [...current.logs, ...patch.logs] : current.logs,
    modifiers: patch.modifiers ?? current.modifiers,
    phase: patch.phase ?? current.phase,
    player_ID: patch.player_ID ?? current.player_ID,
    players: patch.players ?? current.players,
    positions: patch.positions ?? current.positions,
    prompts: patch.prompts ?? current.prompts,
    status: patch.status ?? current.status,
    turn: patch.turn ?? current.turn,
  }
}

function socket_reducer(message: SocketResponse | null) {
  if (!message?.type) return
  switch (message.type) {
    case 'game-start': {
      gameStore.setState((g) => ({ ...g, ready: true }))
      return
    }
    case 'post-connect': {
      if (message.game) {
        gameStore.setState(() => message.game!)
      }
      return
    }
    case 'game': {
      if (message.game) {
        gameStore.setState(() => message.game!)
      }
      return
    }
    case 'game-patch': {
      if (message.game) {
        gameStore.setState((g) => applyGamePatch(g, message.game!))
      }
      return
    }
    case 'lobby': {
      lobbyStore.setState((c) => ({
        ...c,
        players: message.lobby?.players!,
        spectators: message.lobby?.spectators!,
        ready: message.lobby?.ready!,
      }))
      return
    }
    case 'on-connect': {
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

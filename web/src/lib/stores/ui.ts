import { Store } from '@tanstack/store'
import type { Game } from '../game/game'
import { isIdNull } from '../game/core'

export type Ui = {
  active_actor: string | null
  hover_position: string | null
  range_positions: string[]
  source_actor: string | null
  target_positions: string[]
}

const INITIAL_UI: Ui = {
  active_actor: null,
  hover_position: null,
  range_positions: [],
  source_actor: null,
  target_positions: [],
}

const uiStore = new Store<Ui>(INITIAL_UI)

function setActiveActor(actor_id: string | null) {
  uiStore.setState((old) => ({
    ...old,
    active_actor: actor_id,
  }))
}

function setHoverPosition(actor_id: string | null) {
  uiStore.setState((old) => ({
    ...old,
    hover_position: actor_id,
  }))
}

function setRangePositions(positions: string[]) {
  uiStore.setState((old) => ({
    ...old,
    range_positions: positions,
  }))
}

function setSourceActor(actor_id: string | null | undefined) {
  uiStore.setState((old) => ({
    ...old,
    source_actor: actor_id ?? null,
  }))
}

function setTargetPositions(positions: string[]) {
  uiStore.setState((old) => ({
    ...old,
    target_positions: positions,
  }))
}

function setDefaultActiveActor(game: Game) {
  uiStore.setState((old) => {
    if (old.active_actor) {
      return old
    }

    const position = game.positions.find((p) => p.player_ID === game.player_ID && p.rank === 0)
    if (!position || isIdNull(position.actor_ID)) {
      return old
    }

    return {
      ...old,
      active_actor: position.actor_ID,
    }
  })
}

export {
  uiStore,
  setActiveActor,
  setDefaultActiveActor,
  setRangePositions,
  setSourceActor,
  setTargetPositions,
  setHoverPosition,
}

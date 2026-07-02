import { Store } from '@tanstack/store'

export type Ui = {
  active_actor: string | null
  hover_position: string | null
  source_actor: string | null
  target_positions: string[]
}

const INITIAL_UI: Ui = {
  active_actor: null,
  hover_position: null,
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

export {
  uiStore,
  setActiveActor,
  setSourceActor,
  setTargetPositions,
  setHoverPosition,
}

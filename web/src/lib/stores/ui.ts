import { Store } from '@tanstack/store'

type Ui = {
  active_actor: string | null
  hover_position: string | null
}

const INITIAL_UI: Ui = {
  active_actor: null,
  hover_position: null,
}

const uiStore = new Store<Ui>(INITIAL_UI)

function setActiveActor(actor_id: string) {
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

export { uiStore, setActiveActor, setHoverPosition }

import { Store } from '@tanstack/store'
import type { ActorConfig, TeamConfig } from '../game/team'

export type TeamState = TeamConfig & {
  active_actor: number
}
const INITIAL_TEAM: TeamState = {
  active_actor: 0,
  actors: [
    {
      class: null,
      items: [],
      name: '',
      weapon_l: null,
      weapon_r: null,
    },
    {
      class: null,
      items: [],
      name: '',
      weapon_l: null,
      weapon_r: null,
    },
    {
      class: null,
      items: [],
      name: '',
      weapon_l: null,
      weapon_r: null,
    },
    {
      class: null,
      items: [],
      name: '',
      weapon_l: null,
      weapon_r: null,
    },
  ],
}

const teamStore = new Store<TeamState>(INITIAL_TEAM)

function setActiveActor(active_actor: number) {
  teamStore.setState((old) => ({
    ...old,
    active_actor,
  }))
}

function updateActor(
  index: number,
  updater: (actor: ActorConfig) => ActorConfig
) {
  teamStore.setState((old) => ({
    ...old,
    actors: old.actors.map((a, i) => (i === index ? updater(a) : a)),
  }))
}

export { teamStore, setActiveActor, updateActor }

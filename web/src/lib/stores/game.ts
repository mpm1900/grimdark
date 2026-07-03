import { Store } from '@tanstack/store'
import type { Game } from '../game/game'
import { setDefaultActiveActor, setSourceActor, setTargetPositions } from './ui'

const INITIAL_GAME: Game = {
  active_context: null,
  actors: [],
  commands: [],
  logs: [],
  modifiers: [],
  phase: 'init',
  positions: [],
  player_ID: '',
  players: [],
  prompts: [],
  status: 'idle',
  turn: 0,
}

const gameStore = new Store<Game>(INITIAL_GAME)

gameStore.subscribe((game) => {
  console.log('game', game)
  setSourceActor(game.active_context?.source_ID)
  setTargetPositions(game.active_context?.position_IDs ?? [])
  setDefaultActiveActor(game)
})

export { gameStore }

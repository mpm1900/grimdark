import { Store } from '@tanstack/store'
import type { Game } from '../game/game'

const INITIAL_GAME: Game = {
  slice_ID: '',
  active_context: null,
  actors: [],
  logs: [],
  modifiers: [],
  phase: 'init',
  players: [],
  prompts: [],
  status: 'idle',
  turn: 0,
}

const gameStore = new Store<Game>(INITIAL_GAME)

gameStore.subscribe((game) => {
  console.log('game', game)
})

export { gameStore }

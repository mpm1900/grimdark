import { Store } from '@tanstack/store'
import type { Game } from '../game/game'

const INITIAL_GAME: Game = {
  active_context: null,
  actors: [],
  logs: [],
  phase: 'init',
  players: [],
  status: 'idle',
  turn: 0,
}

const gameStore = new Store<Game>(INITIAL_GAME)

gameStore.subscribe((game) => {
  console.log('game', game)
})

export { gameStore }

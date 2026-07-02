import { Store } from '@tanstack/store'
import type { Game } from '../game/game'
import { setSourceActor, setTargetPositions } from './ui'

const INITIAL_GAME: Game = {
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
  console.log(game.active_context?.source_ID)
  setSourceActor(game.active_context?.source_ID)
  setTargetPositions(game.active_context?.position_IDs ?? [])
})

export { gameStore }

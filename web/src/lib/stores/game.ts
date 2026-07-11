import { Store } from '@tanstack/store'
import type { Game } from '../game/game'
import { setDefaultActiveActor, setSourceActor, setTargetPositions } from './ui'

const INITIAL_GAME: Game = {
  instance_ID: null,
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
  ready: false,
  status: 'idle',
  turn: 0,
}

const gameStore = new Store<Game>(INITIAL_GAME)

let turn = 0
gameStore.subscribe((game) => {
  console.log('game', game)
  setSourceActor(game.active_context?.source_ID)
  setTargetPositions(game.active_context?.position_IDs ?? [])
  if (game.turn !== turn) {
    turn = game.turn
    setDefaultActiveActor(game)
  }
})

export { gameStore }

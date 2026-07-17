import type { Player } from '#/lib/game/player'
import { gameStore } from '#/lib/stores/game'
import { useSelector } from '@tanstack/react-store'

function PlayerTeam({ player }: { player: Player }) {
  const actors = useSelector(gameStore, (g) =>
    g.actors
      .filter((a) => a.player_ID === player.ID)
      .slice()
      .reverse()
  )
  return (
    <div className="flex gap-0">
      <div className="font-serif text-center">
        <div className="text-5xl font-black">
          {actors.filter((a) => a.is_alive).length}
        </div>
        <div>remaining</div>
      </div>
    </div>
  )
}

export { PlayerTeam }

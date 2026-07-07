import type { Player } from '#/lib/game/player'
import { gameStore } from '#/lib/stores/game'
import { useSelector } from '@tanstack/react-store'
import { GothicFrame } from './gothic-ui/frame'

function PlayerTeam({ player }: { player: Player }) {
  const actors = useSelector(gameStore, (g) =>
    g.actors
      .filter((a) => a.player_ID === player.ID)
      .slice()
      .reverse()
  )
  return (
    <div className="flex gap-0">
      {Array.from({ length: player.actor_count }).map((_, i) => {
        const actor = actors[i]
        return (
          <GothicFrame
            key={i}
            className="size-9 grid place-items-center font-cinzel-dec"
          >
            {actor?.name[0]}
          </GothicFrame>
        )
      })}
    </div>
  )
}

export { PlayerTeam }

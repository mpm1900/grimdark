import type { Player } from '#/lib/game/player'
import { gameStore } from '#/lib/stores/game'
import { useSelector } from '@tanstack/react-store'
import { GothicFrame } from './gothic-ui/frame'
import { cn } from '#/lib/utils'

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
            className={cn(
              'size-9 grid place-items-center text-foreground/60 font-cinzel-dec font-semibold',
              !!actor?.position_ID && 'bg-foreground/60 text-neutral-950',
              !!actor && !actor.is_alive && 'bg-red-500/20 text-foreground/60'
            )}
          >
            {!!actor && !actor.is_alive ? 'X' : (actor?.name[0] ?? '?')}
          </GothicFrame>
        )
      })}
    </div>
  )
}

export { PlayerTeam }

import type { Player } from '#/lib/game/player'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { getVariant, Platform, PlatformParent } from './platform'
import { setHoverPosition, uiStore } from '#/lib/stores/ui'
import { lobbyStore } from '#/lib/stores/clients'
import { PlayerPosition } from './player-position'
import { gameStore } from '#/lib/stores/game'
import { AnimatePresence, LayoutGroup } from 'motion/react'

function PlayerPositions({
  player,
  className,
  reverse,
  ...props
}: React.ComponentProps<'div'> & { player: Player; reverse?: boolean }) {
  const client = useSelector(lobbyStore, (s) => s.client!)
  const status = useSelector(gameStore, (g) => g.status)
  const positions = useSelector(gameStore, (g) =>
    g.positions
      .filter((p) => p.player_ID === player.ID)
      .sort((a, b) => a.rank - b.rank)
  )
  const ui = useSelector(uiStore, (ui) => ui)
  const hover_position = useSelector(uiStore, (s) => s.hover_position)
  return (
    <div
      {...props}
      className={cn(
        'relative z-0 flex flex-row-reverse items-end',
        reverse && 'flex-row',
        className
      )}
      style={props.style}
    >
      <PlatformParent
        reverse={reverse}
        className="flex-1 h-full absolute inset-0 z-0 "
      >
        {positions.map((position) => (
          <Platform
            key={position.ID}
            rank={position.rank}
            variant={getVariant(ui, client.ID, position, status)}
            className="flex-1"
          />
        ))}
      </PlatformParent>
      <div
        className={cn(
          'relative z-10 flex flex-1 flex-row-reverse items-end gap-1 h-full',
          reverse && 'flex-row'
        )}
      >
        <LayoutGroup id={`player-positions-${player.ID}`}>
          <AnimatePresence initial={false}>
            {positions.map((position) => (
              <PlayerPosition
                key={position.ID}
                hover_position={hover_position}
                position={position}
                reverse={reverse}
                onMouseEnter={() => setHoverPosition(position.ID)}
                onMouseLeave={() => setHoverPosition(null)}
              />
            ))}
          </AnimatePresence>
        </LayoutGroup>
      </div>
    </div>
  )
}

export { PlayerPositions }

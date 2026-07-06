import type { Player } from '#/lib/game/player'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { getVariant, Platform, PlatformParent } from './platform'
import { setHoverPosition, uiStore } from '#/lib/stores/ui'
import { clientsStore } from '#/lib/stores/clients'
import { PlayerPosition } from './player-position'
import { gameStore } from '#/lib/stores/game'

function PlayerPositions({
  player,
  className,
  reverse,
  ...props
}: React.ComponentProps<'div'> & { player: Player; reverse?: boolean }) {
  const client = useSelector(clientsStore, (s) => s.me!)
  const status = useSelector(gameStore, (g) => g.status)
  const positions = useSelector(gameStore, g => g.positions.filter(p => p.player_ID === player.ID))
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
      <PlatformParent reverse={reverse} className="flex-1">
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
          'relative z-10 flex flex-1 flex-row-reverse items-end gap-1',
          reverse && 'flex-row'
        )}
      >
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
      </div>
    </div>
  )
}

export { PlayerPositions }

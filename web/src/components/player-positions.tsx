import type { Player } from '#/lib/game/player'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { ActorFrameSlim } from './actor-frame'
import { getVariant, Platform, PlatformParent } from './platform'
import { type ComponentProps } from 'react'
import { setActiveActor, setHoverPosition, uiStore } from '#/lib/stores/ui'
import { clientsStore } from '#/lib/stores/clients'

function PlayerPosition({
  className,
  hover_position,
  position,
  reverse,
  ...props
}: ComponentProps<'div'> & {
  hover_position: string | null
  position: Player['positions'][number]
  reverse?: boolean
}) {
  const actor = useSelector(gameStore, (g) =>
    g.actors.find((a) => a.position_ID === position.ID)
  )
  const active_actor = useSelector(uiStore, (s) => s.active_actor)
  return (
    <div
      className={cn('relative flex-1 h-[428px] -mb-2', className)}
      {...props}
    >
      {actor && (
        <div
          className="relative z-10 px-8 py-2 flex justify-center"
          onClick={() => {
            setActiveActor(actor.ID)
          }}
        >
          <img
            src="/img/69_Asset_35.png"
            className={cn(
              'w-full max-w-30 relative z-10',
              !actor && 'opacity-0',
              reverse && '-scale-x-100',
              {
                'opacity-50': hover_position !== position.ID,
                'opacity-100': active_actor === position.actor_ID,
              }
            )}
          />
        </div>
      )}
      {actor && (
        <ActorFrameSlim
          actor={actor}
          className="relative z-40 px-2 cursor-pointer"
          onClick={() => {
            setActiveActor(actor.ID)
          }}
        />
      )}
    </div>
  )
}

function PlayerPositions({
  player,
  className,
  reverse,
  ...props
}: React.ComponentProps<'div'> & { player: Player; reverse?: boolean }) {
  const client = useSelector(clientsStore, (s) => s.me!)
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
        {player.positions.map((position) => (
          <Platform
            key={position.ID}
            variant={getVariant(ui, client.ID, position)}
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
        {player.positions.map((position) => (
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

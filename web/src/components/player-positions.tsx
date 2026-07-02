import type { Player } from '#/lib/game/player'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { ActorFrameSlim } from './actor-frame'
import { Popover, PopoverTrigger } from './ui/popover'
import { GothicPopoverContent } from './gothic-ui/popover'
import { ActorDetails } from './actor-details'
import { Platform, PlatformParent } from './platform'
import { type ComponentProps } from 'react'
import { setActiveActor, setHoverPosition, uiStore } from '#/lib/stores/ui'

function PlayerPosition({
  className,
  position,
  reverse,
  ...props
}: ComponentProps<'div'> & {
  position: Player['positions'][number]
  reverse?: boolean
}) {
  const actor = useSelector(gameStore, (g) =>
    g.actors.find((a) => a.position_ID === position.ID)
  )
  return (
    <div className={cn('relative flex-1', className)} {...props}>
      {actor && (
        <Popover>
          <PopoverTrigger className="w-full">
            <div className="relative z-10 px-8 flex justify-center">
              <img
                src="/img/69_Asset_35.png"
                className={cn(
                  'w-full max-w-30 [image-rendering:pixelated] relative z-10',
                  !actor && 'opacity-0',
                  reverse && '-scale-x-100'
                )}
              />
            </div>
          </PopoverTrigger>
          <GothicPopoverContent
            className="w-auto"
            side={reverse ? 'left' : 'right'}
            align="end"
            collisionPadding={16}
          >
            <ActorDetails actor={actor} />
          </GothicPopoverContent>
        </Popover>
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
            variant={
              hover_position === position.ID
                ? reverse
                  ? 'enemy-active'
                  : 'player-active'
                : reverse
                  ? 'enemy'
                  : 'player'
            }
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

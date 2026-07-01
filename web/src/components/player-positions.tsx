import type { Player } from '#/lib/game/player'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { ActorFrame } from './actor-frame'
import { Popover, PopoverTrigger } from './ui/popover'
import { GothicPopoverContent } from './gothic-ui/popover'
import { ActorDetails } from './actor-details'

function PlayerPosition({
  position,
  reverse,
}: {
  position: Player['positions'][number]
  reverse?: boolean
}) {
  const actor = useSelector(gameStore, (g) =>
    g.actors.find((a) => a.position_ID === position.ID)
  )
  return (
    <div className="flex-1">
      <img
        src="/gothic/CharSHRef.png"
        className={cn(
          '-scale-x-100 p-4 px-8 pb-0 -mb-4',
          reverse && 'scale-x-100',
          !actor && 'opacity-0'
        )}
      />
      {actor && (
        <Popover>
          <PopoverTrigger asChild>
            <ActorFrame actor={actor} />
          </PopoverTrigger>
          <GothicPopoverContent
            className="w-auto"
            collisionPadding={16}
            sideOffset={4}
          >
            <ActorDetails actor={actor} />
          </GothicPopoverContent>
        </Popover>
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
  return (
    <div
      className={cn(
        'flex flex-row-reverse items-end',
        reverse && 'flex-row',
        className
      )}
      {...props}
    >
      {player.positions.map((position) => (
        <PlayerPosition
          key={position.ID}
          position={position}
          reverse={reverse}
        />
      ))}
    </div>
  )
}

export { PlayerPositions }

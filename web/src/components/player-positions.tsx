import type { Player } from '#/lib/game/player'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { ActorFrame } from './actor-frame'
import { Popover, PopoverTrigger } from './ui/popover'
import { GothicPopoverContent } from './gothic-ui/popover'
import { ActorDetails } from './actor-details'

const positionStyle = {
  transformStyle: 'preserve-3d',
} satisfies React.CSSProperties

const platformStyle = {
  transform: 'translateX(-50%) rotateX(86deg)',
  transformOrigin: 'center center',
} satisfies React.CSSProperties

function getSidePerspectiveStyle(reverse?: boolean) {
  return {
    perspective: '300px',
    perspectiveOrigin: reverse ? '0% 50%' : '100% 50%',
    transformStyle: 'preserve-3d',
  } satisfies React.CSSProperties
}

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
    <div className="relative flex-1">
      {actor && (
        <Popover>
          <PopoverTrigger>
            <div className="relative z-10">
              <img
                src="/gothic/CharSHRef.png"
                className={cn('relative z-10', !actor && 'opacity-0')}
              />
            </div>
          </PopoverTrigger>
          <GothicPopoverContent
            className="w-auto"
            side={reverse ? 'left' : 'right'}
            align="end"
            collisionPadding={16}
            sideOffset={4}
          >
            <ActorDetails actor={actor} />
          </GothicPopoverContent>
        </Popover>
      )}
      {actor && <ActorFrame actor={actor} className="relative z-40" />}
    </div>
  )
}

function PlayerPositionPlatform({ reverse }: { reverse?: boolean }) {
  return (
    <div className="relative flex-1" style={positionStyle}>
      <div
        className={cn(
          'pointer-events-none absolute bottom-14 left-1/2 size-20 w-full border-6 ring-2 ring-black',
          {
            'bg-red-950/40 border-red-900': reverse,
            'bg-emerald-950/40 border-emerald-800': !reverse,
          }
        )}
        style={platformStyle}
      />
    </div>
  )
}

function PlayerPositions({
  player,
  className,
  reverse,
  ...props
}: React.ComponentProps<'div'> & { player: Player; reverse?: boolean }) {
  const sidePerspectiveStyle = getSidePerspectiveStyle(reverse)

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
      <div
        className={cn(
          'pointer-events-none absolute inset-0 z-0 gap-2 flex flex-row-reverse items-end',
          reverse && 'flex-row -translate-x-4',
          !reverse && 'translate-x-4'
        )}
        style={sidePerspectiveStyle}
      >
        {player.positions.map((position) => (
          <PlayerPositionPlatform key={position.ID} reverse={reverse} />
        ))}
      </div>
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
          />
        ))}
      </div>
    </div>
  )
}

export { PlayerPositions }

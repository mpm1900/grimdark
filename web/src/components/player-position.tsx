import { sendContextMessage } from '#/lib/stores/socket'
import { setActiveActor, uiStore } from '#/lib/stores/ui'
import { useSelector } from '@tanstack/react-store'
import { ActorFrameSlim } from './actor-frame'
import { GothicFramedButton } from './gothic-ui/button'
import type { Player } from '#/lib/game/player'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { isIdNull } from '#/lib/game/core'
import type { Actor } from '#/lib/game/actor'
import type { Position } from '#/lib/game/position'

function PlayerSprite({
  actor,
  className,
  hover_position,
  position_ID,
  ...props
}: React.ComponentProps<'div'> & {
  actor: Actor
  hover_position: string | null
  position_ID: string
}) {
  const client = useSelector(clientsStore, (s) => s.me)
  const status = useSelector(gameStore, (g) => g.status)
  const commands = useSelector(gameStore, (g) => g.commands)
  const active_actor = useSelector(uiStore, (s) => s.active_actor)
  const actor_command = commands.find((c) => c.context.source_ID === actor?.ID)
  return (
    <div
      className={cn(
        'relative z-10 px-8 py-2 flex justify-center',
        {
          'opacity-50': hover_position !== position_ID,
          'opacity-100': active_actor === actor.ID,
        },
        className
      )}
      onClick={() => {
        setActiveActor(actor.ID)
      }}
      {...props}
    >
      <img
        src="/img/69_Asset_35.png"
        className={cn(
          'w-full max-w-30 relative z-10 pointer-events-none select-none',
          !actor && 'opacity-0'
        )}
      />
      {!!actor_command && status === 'idle' && (
        <div className="absolute inset-0 grid place-items-center z-20">
          <GothicFramedButton
            onClick={() => {
              sendContextMessage({
                type: 'cancel-action',
                client_ID: client?.ID!,
                context: actor_command.context,
              })
            }}
          >
            Cancel Action
          </GothicFramedButton>
        </div>
      )}
    </div>
  )
}

function PlayerPosition({
  className,
  hover_position,
  position,
  reverse,
  ...props
}: React.ComponentProps<'div'> & {
  hover_position: string | null
  position: Position
  reverse?: boolean
}) {
  const actor = useSelector(gameStore, (g) =>
    isIdNull(position.actor_ID)
      ? undefined
      : g.actors.find((a) => a.ID === position.actor_ID)
  )
  return (
    <div
      id={position.ID}
      data-position-id={position.ID}
      data-actor-id={position.actor_ID}
      className={cn('relative flex-1 h-107 -mb-2', className)}
      {...props}
    >
      {actor && (
        <PlayerSprite
          actor={actor}
          hover_position={hover_position}
          position_ID={position.ID}
          className={cn({
            '-scale-x-100': reverse,
          })}
        />
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

export { PlayerPosition }

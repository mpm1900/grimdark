import { sendContextMessage } from '#/lib/stores/socket'
import { setActiveActor, uiStore } from '#/lib/stores/ui'
import { useSelector } from '@tanstack/react-store'
import { ActorFrameSlim } from './actor-frame'
import { GothicBigButton } from './gothic-ui/button'
import { lobbyStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { isIdNull } from '#/lib/game/core'
import type { Actor } from '#/lib/game/actor'
import type { Position } from '#/lib/game/position'
import { AnimatePresence, motion } from 'motion/react'
import { v4 } from 'uuid'

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
  const client = useSelector(lobbyStore, (s) => s.client)
  const status = useSelector(gameStore, (g) => g.status)
  const commands = useSelector(gameStore, (g) => g.commands)
  const ui = useSelector(uiStore, (s) => s)
  const actor_commands = commands.filter(
    (c) => c.context.source_ID === actor?.ID
  )
  const is_highlighted =
    status === 'idle'
      ? ui.active_actor === actor.ID || hover_position === position_ID
      : status === 'running'
        ? ui.source_actor === actor.ID ||
          ui.target_positions.includes(position_ID)
        : true
  return (
    <div
      className={cn(
        'relative z-10 py-2 flex items-end w-full min-w-0 justify-center h-full max-h-70',
        {
          'opacity-50': !is_highlighted,
          'opacity-100': is_highlighted,
        },
        className
      )}
      onClick={() => {
        setActiveActor(actor.ID)
      }}
      {...props}
    >
      <img
        src={actor.sprite_url}
        className={cn(
          'h-full w-full object-contain object-bottom max-w-60 relative z-10 pointer-events-none select-none',
          !actor && 'opacity-0'
        )}
      />
      {status === 'idle' && (
        <div className="absolute inset-0 justify-center flex flex-col gap-1 z-20">
          {actor_commands.map((actor_command) => (
            <GothicBigButton
              key={actor_command.ID}
              className="p-0 min-h-9"
              variant="red"
              onClick={() => {
                sendContextMessage({
                  request_ID: v4(),
                  type: 'cancel-action',
                  client_ID: client?.ID!,
                  context: actor_command.context,
                })
              }}
            >
              Cancel {actor_command.payload.config.name}
            </GothicBigButton>
          ))}
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
}: React.ComponentProps<typeof motion.div> & {
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
    <motion.div
      layout="position"
      initial={false}
      transition={{ type: 'spring', stiffness: 260, damping: 28 }}
      id={position.ID}
      data-position-id={position.ID}
      data-actor-id={position.actor_ID}
      className={cn(
        'relative flex-1 basis-0 min-w-0 -mb-2 group h-full flex flex-col justify-end',
        className
      )}
      {...props}
    >
      <AnimatePresence initial={false} mode="popLayout">
        {actor && (
          <motion.div
            key={actor.ID}
            layout
            layoutId={`actor-${actor.ID}`}
            initial={{ opacity: 0, y: 16, scale: 0.96 }}
            animate={{ opacity: 1, y: 0, scale: 1 }}
            exit={{ opacity: 0, y: 16, scale: 0.96 }}
            transition={{ type: 'spring', stiffness: 260, damping: 28 }}
            className="relative h-full w-full min-w-0 flex flex-col justify-end"
          >
            <PlayerSprite
              actor={actor}
              hover_position={hover_position}
              position_ID={position.ID}
              className={cn({
                '-scale-x-100': reverse,
              })}
            />
            <ActorFrameSlim
              actor={actor}
              className="relative z-40 px-2 cursor-pointer"
              onClick={() => {
                setActiveActor(actor.ID)
              }}
            />
          </motion.div>
        )}
      </AnimatePresence>
    </motion.div>
  )
}

export { PlayerPosition }

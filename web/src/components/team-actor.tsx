import { CLASS_STATS, type ID } from '#/lib/game/core'
import type { ActorConfig } from '#/lib/game/team'
import { actorsQuery } from '#/lib/queries/get-actors'
import { setActiveActor, setHoverActor, teamStore } from '#/lib/stores/team'
import { cn } from '#/lib/utils'
import { useQuery } from '@tanstack/react-query'
import { Gauge } from './gothic-ui/progress'
import { StatIcon } from './stat-name'
import { useSelector } from '@tanstack/react-store'
import { motion } from 'motion/react'

function ClassSprite({
  index,
  actor_class_id,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor_class_id: ID | null; index: number }) {
  const actors_query = useQuery(actorsQuery)
  const active_index = useSelector(teamStore, (s) => s.active_actor)
  const actor_class = actors_query.data?.find((a) => a.ID === actor_class_id)
  return (
    <div
      className={cn(
        'relative z-10 flex h-full flex-1 basis-0 min-w-0 items-end justify-center',
        className
      )}
      {...props}
    >
      <div className="relative flex h-80 w-full min-w-0 items-end justify-center pb-8">
        <img
          src={actor_class?.sprite_url ?? '/gothic/CharSHRef.png'}
          className={cn(
            'pointer-events-none relative z-10 h-full w-full max-w-72 select-none object-contain object-bottom',
            {
              'opacity-50': active_index !== index,
            }
          )}
        />
      </div>
    </div>
  )
}

function TeamActor({
  config,
  index,
  className,
  onClick,
  ...props
}: React.ComponentProps<typeof motion.div> & {
  config: ActorConfig
  index: number
}) {
  const actors_query = useQuery(actorsQuery)
  const active_index = useSelector(teamStore, (s) => s.active_actor)
  const actor_class = actors_query.data?.find((a) => a.ID === config.class)
  return (
    <motion.div
      layout="position"
      transition={{ type: 'spring', stiffness: 260, damping: 28 }}
      className={cn('relative flex flex-1 basis-0 min-w-0 flex-col')}
      onClick={(e) => {
        setActiveActor(index)
        onClick?.(e)
      }}
      onMouseEnter={() => {
        setHoverActor(index)
      }}
      onMouseLeave={() => {
        setHoverActor(null)
      }}
      {...props}
    >
      <div className="relative h-full">
        <ClassSprite actor_class_id={config?.class} index={index} />
        <div className="absolute bottom-0 inset-x-0 text-center h-9 leading-9 mx-1">
          <div>
            <span
              className={cn(
                'text-foreground/60 truncate font-cinzel-dec font-semibold [text-shadow:1px_2px_0_var(--color-black)]',
                active_index === index && 'text-foreground'
              )}
            >
              {config.name || actor_class?.name}
            </span>
          </div>
        </div>
      </div>
      <div className="mx-2 mt-4 h-43 space-y-1">
        {CLASS_STATS.map((stat) => (
          <div key={stat} className="flex items-center gap-2">
            <StatIcon stat={stat} className="size-5" />
            <Gauge
              value={((actor_class?.stats[stat] ?? 0) * 100) / 255}
              className="h-5"
            >
              {actor_class?.stats[stat]}
            </Gauge>
          </div>
        ))}
      </div>
    </motion.div>
  )
}

export { TeamActor }

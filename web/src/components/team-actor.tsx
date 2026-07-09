import { CLASS_STATS, type ID } from '#/lib/game/core'
import type { ActorConfig } from '#/lib/game/team'
import { actorsQuery } from '#/lib/queries/get-actors'
import { setActiveActor } from '#/lib/stores/team'
import { cn } from '#/lib/utils'
import { keys } from '#/utils/maps'
import { useQuery } from '@tanstack/react-query'
import { Gauge } from './gothic-ui/progress'

function ClassSprite({
  actor_class_id,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor_class_id: ID | null }) {
  const actors_query = useQuery(actorsQuery)
  const actor_class = actors_query.data?.find((a) => a.ID === actor_class_id)
  return (
    <div
      className={cn(
        'relative z-10 flex h-full flex-1 basis-0 min-w-0 items-end justify-center',
        className
      )}
      {...props}
    >
      <div className="relative flex h-full w-full min-w-0 max-h-80 items-end justify-center pb-8">
        <img
          src={actor_class?.sprite_url ?? '/gothic/CharSHRef.png'}
          className={cn(
            'pointer-events-none relative z-10 h-full w-full max-w-72 select-none object-contain object-bottom'
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
}: React.ComponentProps<'div'> & { config: ActorConfig; index: number }) {
  const actors_query = useQuery(actorsQuery)
  const actor_class = actors_query.data?.find((a) => a.ID === config.class)
  return (
    <div
      className={cn('relative flex flex-1 basis-0 min-w-0 flex-col')}
      onClick={(e) => {
        setActiveActor(index)
        onClick?.(e)
      }}
      {...props}
    >
      <div className="relative h-full">
        <ClassSprite actor_class_id={config?.class} />
        <div className="absolute bottom-0 inset-x-0 text-center h-9 leading-9 mx-1">
          <div>
            {actor_class?.name ? (
              <span className="text-foreground font-cinzel-dec font-semibold [text-shadow:1px_2px_0_var(--color-black)]">
                {actor_class.name}
              </span>
            ) : (
              'Select a Class'
            )}
          </div>
        </div>
      </div>
      {actor_class && (
        <div className="mx-2 mt-4 h-40 space-y-1">
          {CLASS_STATS.map((stat) => (
            <Gauge key={stat} value={(actor_class.stats[stat] * 100) / 255} />
          ))}
        </div>
      )}
      {!actor_class && <div className="mx-2 mt-4 h-40">details here</div>}
    </div>
  )
}

export { TeamActor }

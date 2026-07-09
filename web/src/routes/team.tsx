import { ActorCombobox } from '#/components/actor-combobox'
import { AppHeader } from '#/components/app-header'
import { GothicCard } from '#/components/gothic-ui/card'
import { Platform, PlatformParent } from '#/components/platform'
import { Input } from '#/components/ui/input'
import { Marker, MarkerContent } from '#/components/ui/marker'
import type { ID } from '#/lib/game/core'
import { actorsQuery } from '#/lib/queries/get-actors'
import { setActiveActor, teamStore, updateActor } from '#/lib/stores/team'
import { cn } from '#/lib/utils'
import { useQuery } from '@tanstack/react-query'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'

export const Route = createFileRoute('/team')({
  component: RouteComponent,
})

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
      <div className="relative flex h-full w-full min-w-0 max-h-80 items-end justify-center pb-16">
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

function RouteComponent() {
  const actors_query = useQuery(actorsQuery)
  const team = useSelector(teamStore, (s) => s)
  const active_actor = team.actors[team.active_actor]
  const active_class = actors_query.data?.find(
    (c) => c.ID === active_actor.class
  )
  return (
    <ClientOnly>
      <AppHeader />
      <div className="relative flex flex-col h-full">
        <div className="absolute inset-0 bottom-1/2 bg-neutral-900 z-0" />
        <div className="relative flex h-full">
          <div className="flex-1 px-10 flex flex-col justify-center min-w-0">
            <div className="relative z-0 flex items-end h-2/3 min-h-88 max-h-192 min-w-0">
              <PlatformParent className="flex-1 h-full absolute inset-0 z-0 perspective-origin-top">
                {team.actors.map((_, i) => (
                  <Platform
                    key={i}
                    rank={i}
                    variant={
                      i === team.active_actor ? 'player-active' : 'player'
                    }
                    className="flex-1"
                  >
                    {''}
                  </Platform>
                ))}
              </PlatformParent>
              <div className="flex-1 h-full min-w-0 flex flex-row-reverse">
                {team.actors.map((a, i) => (
                  <ClassSprite
                    key={i}
                    actor_class_id={a.class}
                    onClick={() => {
                      setActiveActor(i)
                    }}
                  />
                ))}
              </div>
            </div>
          </div>
          <div className="w-1/3 p-8 h-full grid place-items-center">
            <GothicCard className="p-4 w-full gap-2">
              <Input
                value={active_actor.name ?? active_class?.name}
                placeholder="Name"
              />
              <Marker variant="separator">
                <MarkerContent>Class</MarkerContent>
              </Marker>
              <ActorCombobox
                value={active_actor.class}
                onValueChange={(c) => {
                  updateActor(team.active_actor, (old) => ({
                    ...old,
                    class: c,
                  }))
                }}
              />
            </GothicCard>
          </div>
        </div>
      </div>
    </ClientOnly>
  )
}

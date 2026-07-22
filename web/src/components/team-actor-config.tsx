import { teamStore, updateActor } from '#/lib/stores/team'
import { useSelector } from '@tanstack/react-store'
import { ActorCombobox } from './actor-combobox'
import { GothicCard } from './gothic-ui/card'
import { Marker, MarkerContent } from './ui/marker'
import { useQuery } from '@tanstack/react-query'
import { actorsQuery } from '#/lib/queries/get-actors'
import { WeaponCombobox } from './weapon-combobox'
import { GothicFrame } from './gothic-ui/frame'
import type { ActorClass } from '#/lib/game/actor-class'
import type { ActorConfig } from '#/lib/game/team'
import { entries } from '#/utils/maps'
import { v4 } from 'uuid'

function WeaponsConfig({
  actor_class,
  value,
  onValueChange,
}: {
  actor_class: ActorClass
  value: ActorConfig
  onValueChange: (value: ActorConfig) => void
}) {
  const total_weight = Object.values(value.weapons).reduce((prev, curr) => {
    const found = actor_class.options.weapons.find((w) => w.ID == curr)
    return prev + (found?.weight ?? 0)
  }, 0)
  const remaining_weight = actor_class.strength - total_weight

  return (
    <div className="flex flex-col gap-1">
      <Marker variant="separator">
        <MarkerContent>Weapons</MarkerContent>
      </Marker>
      <div className="grid grid-cols-1">
        {entries(value.weapons).map(([slot, weapon]) => (
          <WeaponCombobox
            disabled={remaining_weight <= 0 && !weapon}
            options={actor_class.options.weapons}
            value={weapon}
            remaining_weight={remaining_weight}
            onValueChange={(w) => {
              onValueChange({
                ...value,
                weapons: {
                  ...value.weapons,
                  [slot]: w,
                },
              })
            }}
          />
        ))}
      </div>
    </div>
  )
}

function TeamActorConfig() {
  const team = useSelector(teamStore, (s) => s)
  const active_actor = team.config.actors[team.active_actor]
  const actors_query = useQuery(actorsQuery)
  const active_class = actors_query.data?.find(
    (c) => c.ID === active_actor.class
  )
  return (
    <div className="flex-1 grid place-items-start">
      <GothicCard className="p-2 w-full gap-4">
        <div className="flex flex-col gap-1">
          <ActorCombobox
            value={active_actor.class}
            onValueChange={(class_ID) => {
              const ac = actors_query.data?.find((c) => c.ID === class_ID)
              if (!ac) return
              updateActor(team.active_actor, (old) => ({
                ...old,
                class: class_ID,
                weapons: Object.fromEntries(
                  Array.from({ length: ac.arms }).map((_) => [v4(), null])
                ),
              }))
            }}
          />
        </div>
        {active_class && (
          <WeaponsConfig
            actor_class={active_class}
            value={active_actor}
            onValueChange={(config) => {
              updateActor(team.active_actor, (old) => ({
                ...old,
                ...config,
              }))
            }}
          />
        )}
        {active_class && (
          <div className="flex flex-col gap-1">
            <Marker variant="separator">
              <MarkerContent>Items</MarkerContent>
            </Marker>
            <div className="flex justify-center [&>div]:size-16">
              <GothicFrame></GothicFrame>
              <GothicFrame></GothicFrame>
              <GothicFrame></GothicFrame>
            </div>
          </div>
        )}
      </GothicCard>
    </div>
  )
}

export { TeamActorConfig }

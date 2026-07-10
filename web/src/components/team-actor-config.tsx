import { teamStore, updateActor } from '#/lib/stores/team'
import { useSelector } from '@tanstack/react-store'
import { ActorCombobox } from './actor-combobox'
import { GothicCard } from './gothic-ui/card'
import { Marker, MarkerContent } from './ui/marker'
import { useQuery } from '@tanstack/react-query'
import { actorsQuery } from '#/lib/queries/get-actors'
import { WeaponCombobox } from './weapon-combobox'
import { GothicFrame } from './gothic-ui/frame'

function TeamActorConfig() {
  const team = useSelector(teamStore, (s) => s)
  const active_actor = team.actors[team.active_actor]
  const actors_query = useQuery(actorsQuery)
  const active_class = actors_query.data?.find(
    (c) => c.ID === active_actor.class
  )
  return (
    <div className="w-1/3 p-8 h-full grid place-items-start">
      <GothicCard className="p-2 w-full gap-4">
        <div className="flex flex-col gap-1">
          <ActorCombobox
            value={active_actor.class}
            onValueChange={(c) => {
              updateActor(team.active_actor, (old) => ({
                ...old,
                class: c,
                weapon_l: null,
                weapon_r: null,
              }))
            }}
          />
        </div>
        {active_class && (
          <div className="flex flex-col gap-1">
            <Marker variant="separator">
              <MarkerContent>Weapons</MarkerContent>
            </Marker>
            <div className="grid grid-cols-1">
              <WeaponCombobox
                disabled={
                  active_class.options.weapons.find(
                    (w) => w.ID === active_actor.weapon_l
                  )?.hands == 2
                }
                options={active_class.options.weapons}
                value={active_actor.weapon_r}
                other={active_actor.weapon_l}
                onValueChange={(w) => {
                  updateActor(team.active_actor, (old) => ({
                    ...old,
                    weapon_r: w,
                  }))
                }}
              />
              <WeaponCombobox
                disabled={
                  active_class.options.weapons.find(
                    (w) => w.ID === active_actor.weapon_r
                  )?.hands == 2
                }
                options={active_class.options.weapons}
                value={active_actor.weapon_l}
                other={active_actor.weapon_r}
                onValueChange={(w) => {
                  updateActor(team.active_actor, (old) => ({
                    ...old,
                    weapon_l: w,
                  }))
                }}
              />
            </div>
          </div>
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

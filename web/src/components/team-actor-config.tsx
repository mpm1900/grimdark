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

function WeaponsConfig({
  actor_class,
  value,
  onValueChange,
}: {
  actor_class: ActorClass
  value: ActorConfig
  onValueChange: (value: ActorConfig) => void
}) {
  const found_l = actor_class.options.weapons.find(
    (w) => w.ID === value.weapon_l
  )
  const found_r = actor_class.options.weapons.find(
    (w) => w.ID === value.weapon_r
  )
  const remaining_hands =
    actor_class.hands - (found_l?.hands ?? 0) - (found_r?.hands ?? 0)
  console.log('rem', remaining_hands, actor_class)
  return (
    <div className="flex flex-col gap-1">
      <Marker variant="separator">
        <MarkerContent>Weapons</MarkerContent>
      </Marker>
      <div className="grid grid-cols-1">
        <WeaponCombobox
          disabled={remaining_hands <= 0 && !value.weapon_r}
          options={actor_class.options.weapons}
          value={value.weapon_r}
          other={value.weapon_l}
          remaining_hands={remaining_hands}
          onValueChange={(w) => {
            onValueChange({
              ...value,
              weapon_r: w,
            })
          }}
        />
        <WeaponCombobox
          disabled={remaining_hands <= 0 && !value.weapon_l}
          options={actor_class.options.weapons}
          value={value.weapon_l}
          other={value.weapon_r}
          remaining_hands={remaining_hands}
          onValueChange={(w) => {
            onValueChange({
              ...value,
              weapon_l: w,
            })
          }}
        />
      </div>
    </div>
  )
}

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

import { teamStore, updateActor } from '#/lib/stores/team'
import { useSelector } from '@tanstack/react-store'
import { ActorCombobox } from './actor-combobox'
import { GothicCard } from './gothic-ui/card'
import { Input } from './ui/input'
import { Marker, MarkerContent } from './ui/marker'
import { useQuery } from '@tanstack/react-query'
import { actorsQuery } from '#/lib/queries/get-actors'

function TeamActorConfig() {
  const team = useSelector(teamStore, (s) => s)
  const active_actor = team.actors[team.active_actor]
  const actors_query = useQuery(actorsQuery)
  const active_class = actors_query.data?.find(
    (c) => c.ID === active_actor.class
  )
  return (
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
        <Marker variant="separator">
          <MarkerContent>Weapon</MarkerContent>
        </Marker>
      </GothicCard>
    </div>
  )
}

export { TeamActorConfig }

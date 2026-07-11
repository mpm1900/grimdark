import { AppHeader } from '#/components/app-header'
import { GothicBigButton } from '#/components/gothic-ui/button'
import { InstanceCombobox } from '#/components/instance-combobox'
import { TeamActor } from '#/components/team-actor'
import { TeamActorConfig } from '#/components/team-actor-config'
import { TeamPlatforms } from '#/components/team-platforms'
import type { ID } from '#/lib/game/core'
import { useConnect } from '#/lib/mutations/connect'
import { useUser } from '#/lib/queries/auth'
import { teamStore } from '#/lib/stores/team'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
import { LayoutGroup } from 'motion/react'

export const Route = createFileRoute('/_authed/')({
  component: RouteComponent,
})

function RouteComponent() {
  const user = useUser()
  const team = useSelector(teamStore, (s) => s)
  const navigate = Route.useNavigate()
  const mutation = useConnect()
  const actor_order = [...team.actors.map((_, i) => i)]

  function battle(instance_ID?: ID) {
    if (!user.data) return

    mutation.mutate(
      {
        ...team,
        instance_ID,
      },
      {
        onSuccess: (message) => {
          if (!message.game?.instance_ID) return

          navigate({
            to: '/lobby/$gameID',
            params: {
              gameID: message.game?.instance_ID,
            },
          })
        },
      }
    )
  }

  return (
    <ClientOnly>
      <AppHeader />
      <div className="relative flex flex-col justify-center gap-6 h-full">
        <div className="flex gap-6 justify-around items-center z-10">
          <div className="text-center font-cinzel text-3xl capitalize font-semibold">
            ready your team!
          </div>
          <div className="flex gap-2">
            <GothicBigButton
              variant="red"
              className="text-xl"
              disabled={!!team.actors.find((a) => !a.class)}
              onClick={() => battle()}
            >
              Battle!
            </GothicBigButton>
            <InstanceCombobox
              value={null}
              onValueChange={(i) => {
                battle(i)
              }}
              render={
                <GothicBigButton
                  variant="red"
                  className="text-xl"
                  disabled={!!team.actors.find((a) => !a.class)}
                >
                  Join
                </GothicBigButton>
              }
            />
          </div>
        </div>
        <div className="relative flex">
          <div className="flex-1 pl-10 flex flex-col justify-around pb-8 min-w-0">
            <div className="relative z-0 min-h-88">
              <TeamPlatforms />
              <LayoutGroup>
                <div className="h-full min-w-0 flex flex-row-reverse gap-2">
                  {actor_order.map((i) => (
                    <TeamActor key={i} config={team.actors[i]} index={i} />
                  ))}
                </div>
              </LayoutGroup>
            </div>
          </div>
          <TeamActorConfig />
        </div>
      </div>
    </ClientOnly>
  )
}

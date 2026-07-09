import { AppHeader } from '#/components/app-header'
import { GothicFramedButton } from '#/components/gothic-ui/button'
import { TeamActor } from '#/components/team-actor'
import { TeamActorConfig } from '#/components/team-actor-config'
import { TeamPlatforms } from '#/components/team-platforms'
import { useConnect } from '#/lib/mutations/connect'
import { useUser } from '#/lib/queries/auth'
import { teamStore } from '#/lib/stores/team'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'

export const Route = createFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  const user = useUser()
  const team = useSelector(teamStore, (s) => s)
  const navigate = Route.useNavigate()
  const mutation = useConnect()

  function battle() {
    if (!user.data) return

    mutation.mutate(
      {
        ...team,
        user: user.data,
      },
      {
        onSuccess: (message) => {
          if (!message.game?.instance_ID) return

          navigate({
            to: '/battle/$gameID',
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
      <div className="relative flex flex-col h-full">
        <div className="absolute inset-0 bottom-1/2 bg-neutral-900 z-0" />
        <div className="relative flex h-full">
          <div className="flex-1 pl-10 flex flex-col justify-around pb-8 min-w-0">
            <div className="flex gap-6 justify-around">
              <div className="font-cinzel text-3xl capitalize font-semibold">
                ready your team
              </div>
              <GothicFramedButton onClick={() => battle()}>
                Battle!
              </GothicFramedButton>
            </div>
            <div className="relative z-0 min-h-88">
              <TeamPlatforms />
              <div className="h-full min-w-0 flex flex-row-reverse gap-2">
                {team.actors.map((a, i) => (
                  <TeamActor key={i} config={a} index={i} />
                ))}
              </div>
            </div>
          </div>
          <TeamActorConfig />
        </div>
      </div>
    </ClientOnly>
  )
}

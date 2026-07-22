import { AppHeader } from '#/components/app-header'
import {
  GothicBigButton,
  GothicFramedButton,
} from '#/components/gothic-ui/button'
import { InstanceCombobox } from '#/components/instance-combobox'
import { TeamActor } from '#/components/team-actor'
import { TeamActorConfig } from '#/components/team-actor-config'
import { TeamCombobox } from '#/components/team-combobox'
import { TeamPlatforms } from '#/components/team-platforms'
import { Input } from '#/components/ui/input'
import type { ID } from '#/lib/game/core'
import { useConnect } from '#/lib/mutations/connect'
import { useSaveTeam } from '#/lib/mutations/save-team'
import { useUser } from '#/lib/queries/auth'
import { teamStore } from '#/lib/stores/team'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
import { DownloadCloud, Save } from 'lucide-react'
import { LayoutGroup } from 'motion/react'
import { useEffect } from 'react'
import { toast } from 'sonner'

export const Route = createFileRoute('/_authed/')({
  component: RouteComponent,
})

function RouteComponent() {
  const user = useUser()
  const team = useSelector(teamStore, (s) => s)
  const navigate = Route.useNavigate()
  const connect = useConnect()
  const save = useSaveTeam()
  const actor_order = [...team.config.actors.map((_, i) => i)]

  function battle(instance_ID?: ID) {
    if (!user.data) return
    localStorage.setItem('saved-actors', JSON.stringify(team.config.actors))

    connect.mutate(
      {
        ...team.config,
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

  useEffect(() => {
    const json = localStorage.getItem('saved-actors')
    const saved = json && JSON.parse(json)
    if (saved) {
      teamStore.setState((s) => ({
        ...s,
        config: {
          ...s.config,
          actors: saved,
        },
      }))
    }
  }, [])

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
              disabled={!!team.config.actors.find((a) => !a.class)}
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
                  disabled={!!team.config.actors.find((a) => !a.class)}
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
                    <TeamActor
                      key={i}
                      config={team.config.actors[i]}
                      index={i}
                    />
                  ))}
                </div>
              </LayoutGroup>
            </div>
          </div>
          <div className="h-full flex flex-col w-1/3 p-8 gap-2">
            <div className="flex items-center gap-0">
              <Input placeholder="Name your team" value={team.config.name} />
              <div className="flex -space-x-px">
                <GothicFramedButton
                  className="size-10 p-1.5"
                  onClick={() => {
                    save.mutate(team, {
                      onSuccess: (saved) => {
                        toast('Team saved successfully')
                        teamStore.setState((s) => ({
                          ...s,
                          ...saved,
                        }))
                      },
                    })
                  }}
                >
                  <Save />
                </GothicFramedButton>
                <TeamCombobox
                  onSelect={(v) => {
                    teamStore.setState((s) => ({
                      ...s,
                      ...v,
                    }))
                  }}
                  render={
                    <GothicFramedButton className="size-10 p-1.5">
                      <DownloadCloud />
                    </GothicFramedButton>
                  }
                />
              </div>
            </div>
            <TeamActorConfig />
          </div>
        </div>
      </div>
    </ClientOnly>
  )
}

import { ActionButton } from '#/components/action-button'
import { ActionContextDialog } from '#/components/action-context-dialog'
import { ActorFrame } from '#/components/actor-frame'
import { AppHeader } from '#/components/app-header'
import { GothicCard } from '#/components/gothic-ui/card'
import { GothicFrame } from '#/components/gothic-ui/frame'
import { PlayerPositions } from '#/components/player-positions'
import { PromptController } from '#/components/prompt-controller'
import { DialogTrigger } from '#/components/ui/dialog'
import { RenderLog } from '#/lib/game/log'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { uiStore } from '#/lib/stores/ui'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'

export const Route = createFileRoute('/')({ component: Home })

function Home() {
  const game = useSelector(gameStore, (g) => g)
  const client = useSelector(clientsStore, (s) => s.me)
  const active_actor_id = useSelector(uiStore, (s) => s.active_actor)
  const active_actor = game.actors.find((a) => a.ID === active_actor_id)
  const client_player = game.players.find((p) => p.user.id === client?.ID)

  return (
    <ClientOnly>
      <PromptController />
      <div className="h-dvh overflow-hidden bg-center bg-cover">
        <AppHeader />

        <div className="relative flex h-full min-h-0 flex-col overflow-hidden pt-12 z-20">
          <div className="min-h-0 flex-1" />
          <div className="flex shrink-0 px-3 z-10 -mb-1">
            {client_player && (
              <PlayerPositions
                className="flex-1 gap-1"
                player={client_player}
              />
            )}
            <div className="w-1/12" />
            {client_player && (
              <PlayerPositions
                className="flex-1 gap-1"
                player={client_player}
                reverse
              />
            )}
          </div>
          <div className="bg-[url(/gothic/DecoratedLineHorizontal.png)] h-7 bg-center bg-contain bg-repeat-x"></div>
          <div className="flex items-start -z-10 bg-neutral-950">
            <GothicFrame className="relative flex flex-1 flex-col h-full p-0">
              {active_actor && (
                <ActorFrame actor={active_actor} className="-mt-2.5 -ml-1.5" />
              )}
            </GothicFrame>
            <GothicFrame className="grid grid-cols-2 max-w-1/3 gap-px p-0">
              {active_actor?.actions.map((action) => (
                <ActionContextDialog
                  key={action.ID}
                  actor={active_actor}
                  action={action}
                  enabled={!action.is_disabled}
                >
                  <DialogTrigger asChild>
                    <ActionButton action={action} actor={active_actor} />
                  </DialogTrigger>
                </ActionContextDialog>
              ))}
            </GothicFrame>
            <GothicCard className="h-51 min-w-0 max-w-1/4 flex-1 flex bg-neutral-950 p-0">
              <div className="flex-1 overflow-auto font-serif text-foreground/40">
                <ul className="px-2">
                  {game.logs.map((log) => (
                    <li key={log.ID}>{RenderLog(log)}</li>
                  ))}
                </ul>
              </div>
            </GothicCard>
          </div>
        </div>
      </div>
    </ClientOnly>
  )
}

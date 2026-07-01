import { AppHeader } from '#/components/app-header'
import { GothicCard } from '#/components/gothic-ui/card'
import { GothicFrame } from '#/components/gothic-ui/frame'
import { PlayerPositions } from '#/components/player-positions'
import { PromptController } from '#/components/prompt-controller'
import { RenderLog } from '#/lib/game/log'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'

export const Route = createFileRoute('/')({ component: Home })

function Home() {
  const game = useSelector(gameStore, (g) => g)
  const client = useSelector(clientsStore, (s) => s.me)
  const client_player = game.players.find((p) => p.user.id === client?.ID)
  return (
    <ClientOnly>
      <PromptController />
      <div className="h-dvh overflow-hidden">
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
          <div className="bg-[url(/gothic/DecoratedLineHorizontal.png)] h-7 bg-center bg-contain bg-repeat-x -mb-1"></div>
          <div className="flex justify-between items-end -z-10">
            <div></div>
            <GothicCard className="h-50 w-[512px] flex bg-neutral-950 p-0">
              <GothicFrame className="-mx-2 -mt-1.5 py-1.5 leading-0 font-cinzel text-foreground/20 font-semibold bg-neutral-950">
                Battle Log
              </GothicFrame>
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

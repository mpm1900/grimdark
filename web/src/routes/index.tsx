import { AppHeader } from '#/components/app-header'
import { GothicChatPanel } from '#/components/gothic-ui/frame'
import { PlayerPositions } from '#/components/player-positions'
import { PromptController } from '#/components/prompt-controller'
import { ScrollArea } from '#/components/ui/scroll-area'
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

        <div className="relative flex h-full min-h-0 flex-col overflow-hidden pt-12">
          <div className="min-h-0 flex-1" />
          <div className="flex shrink-0 px-6 pb-6">
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
          <div className="flex justify-between items-end">
            <div className="h-[188px]">
              <img
                src="/gothic/SkillBarRightColb_reverse.png"
                className="h-full"
              />
            </div>
            <GothicChatPanel className="h-[188px] w-[512px]">
              <div className="h-8 leading-8 pt-1 mb-1 pl-6 font-cinzel text-foreground/60 font-semibold">
                Battle Log
              </div>
              <ScrollArea className="h-[148px] mr-2 pb-2 font-serif text-foreground/40">
                <ul className="px-4">
                  {game.logs.map((log) => (
                    <li key={log.ID}>{RenderLog(log)}</li>
                  ))}
                </ul>
              </ScrollArea>
            </GothicChatPanel>
          </div>
        </div>
      </div>
    </ClientOnly>
  )
}

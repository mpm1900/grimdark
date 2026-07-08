import { AppHeader } from '#/components/app-header'
import { BattlePanel } from '#/components/panels/battle'
import { PlayerPositions } from '#/components/player-positions'
import { PlayerTeam } from '#/components/player-team'
import { PromptController } from '#/components/prompt-controller'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
export const Route = createFileRoute('/')({ component: Home })

function Home() {
  const game = useSelector(gameStore, (g) => g)
  const client = useSelector(clientsStore, (s) => s.me)
  const client_player = game.players.find((p) => p.user.id === client?.ID)
  const other_player = game.players.find((p) => p.user.id !== client?.ID)

  return (
    <ClientOnly>
      <PromptController />
      <div className="h-dvh overflow-hidden bg-center bg-cover">
        <AppHeader />

        <div className="relative flex h-full min-h-0 flex-col overflow-hidden pt-12 z-20">
          <div className="absolute inset-0 bottom-1/2 bg-neutral-900 z-0" />
          <div className="z-0 flex items-start justify-between p-2">
            {client_player && <PlayerTeam player={client_player} />}
            {other_player && <PlayerTeam player={other_player} />}
          </div>
          <div className="flex flex-1 h-full px-3 z-10 mb-56">
            {client_player && (
              <PlayerPositions
                className="flex-1 gap-1"
                player={client_player}
              />
            )}
            <div className="w-1/12" />
            {other_player && (
              <PlayerPositions
                className="flex-1 gap-1"
                player={other_player}
                reverse
              />
            )}
          </div>
          <BattlePanel />
        </div>
      </div>
    </ClientOnly>
  )
}

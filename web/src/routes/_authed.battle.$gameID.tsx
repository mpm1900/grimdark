import { AppHeader } from '#/components/app-header'
import { BattlePanel } from '#/components/panels/battle'
import { PlayerPositions } from '#/components/player-positions'
import { PlayerTeam } from '#/components/player-team'
import { PromptController } from '#/components/prompt-controller'
import { TargetingArrows } from '#/components/targeting-arrows'
import { TurnContext } from '#/components/turn-context'
import { useReconnect } from '#/hooks/use-reconnect'
import { getInstance } from '#/lib/queries/get-instance'
import { disconnect } from '#/lib/socket/connect'
import { lobbyStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { ClientOnly, createFileRoute, redirect } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
import z from 'zod'

export const Route = createFileRoute('/_authed/battle/$gameID')({
  component: RouteComponent,
  preload: false,
  params: z.object({
    gameID: z.uuid(),
  }),
  loader: async ({ params }) => {
    const instance = await getInstance({ data: params.gameID })
    if (!instance) {
      throw redirect({ to: '/' })
    }
    if (instance.status === 'running') return
    throw redirect({ to: '/lobby/$gameID', params })
  },
  onLeave: () => {
    const next_path = window.location.pathname
    if (next_path.startsWith('/lobby/')) return
    disconnect(1000, 'Route: onLeave')
  },
})

function RouteComponent() {
  const params = Route.useParams()
  const game = useSelector(gameStore, (g) => g)
  const client = useSelector(lobbyStore, (s) => s.client)
  const client_player = game.players.find((p) => p.user.id === client?.ID)
  const other_player = game.players.find((p) => p.user.id !== client?.ID)
  useReconnect(params.gameID)

  return (
    <ClientOnly>
      <PromptController />
      <div className="h-dvh relative overflow-hidden bg-center bg-cover">
        <AppHeader />
        <TurnContext />

        <div className="relative flex h-full min-h-0 flex-col overflow-hidden pt-12 z-20">
          <div className="z-0 flex items-start justify-between p-2">
            {client_player && <PlayerTeam player={client_player} />}
            {other_player && <PlayerTeam player={other_player} />}
          </div>
          <div className="relative flex flex-1 h-full px-3 z-10 mb-56">
            <TargetingArrows />
            {client_player && (
              <PlayerPositions
                className="flex-1 gap-1"
                player={client_player}
              />
            )}
            <div className="w-1/24" />
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

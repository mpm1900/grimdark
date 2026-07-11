import { AppHeader } from '#/components/app-header'
import { GothicFramedButton } from '#/components/gothic-ui/button'
import { GothicCard } from '#/components/gothic-ui/card'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '#/components/ui/table'
import { useReconnect } from '#/hooks/use-reconnect'
import { NULL_CONTEXT } from '#/lib/game/context'
import { getInstance } from '#/lib/queries/get-instance'
import { disconnect } from '#/lib/socket/connect'
import { lobbyStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { sendContextMessage } from '#/lib/stores/socket'
import { ClientOnly, createFileRoute, redirect } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
import { useEffect } from 'react'
import z from 'zod'

export const Route = createFileRoute('/_authed/lobby/$gameID')({
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
    if (instance.status === 'init') return
    throw redirect({ to: '/battle/$gameID', params })
  },
  onLeave: () => {
    const next_path = window.location.pathname
    if (next_path.startsWith('/battle/')) return
    disconnect(1000, 'Route: onLeave')
  },
})

function RouteComponent() {
  const params = Route.useParams()
  const lobby = useSelector(lobbyStore)
  const ready = useSelector(gameStore, (g) => g.ready)
  const navigate = Route.useNavigate()
  useReconnect(params.gameID)
  useEffect(() => {
    if (ready) {
      navigate({
        to: '/battle/$gameID',
        params,
      })
    }
  }, [ready])
  return (
    <ClientOnly>
      <AppHeader />
      <div className="absolute inset-0 bottom-1/2 bg-neutral-900 z-0" />
      <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
        <GothicCard className="relative z-10">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead colSpan={2}>Players</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {lobby.players.map((player) => (
                <TableRow key={player.ID}>
                  <TableCell>{player.user.email}</TableCell>
                  <TableCell>
                    {player.ID === lobby.client?.ID ? (
                      <GothicFramedButton
                        onClick={() => {
                          sendContextMessage({
                            type: 'ready',
                            client_ID: lobby.client?.ID!,
                            context: NULL_CONTEXT,
                          })
                        }}
                      >
                        Ready: {String(!!lobby.ready[player.ID])}
                      </GothicFramedButton>
                    ) : (
                      <div>{String(!!lobby.ready[player.ID])}</div>
                    )}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </GothicCard>
      </div>
    </ClientOnly>
  )
}

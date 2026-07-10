import { AppHeader } from '#/components/app-header'
import { getInstance } from '#/lib/queries/get-instance'
import { disconnect } from '#/lib/socket/connect'
import {
  ClientOnly,
  createFileRoute,
  Link,
  redirect,
} from '@tanstack/react-router'
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
  },
  onLeave: () => {
    const next_path = window.location.pathname
    if (next_path.startsWith('/battle/')) return
    disconnect(1000, 'Route: onLeave')
  },
})

function RouteComponent() {
  const params = Route.useParams()
  return (
    <ClientOnly>
      <AppHeader />
      <div className="relative flex flex-col justify-center gap-6 h-full">
        <div className="absolute inset-0 bottom-1/2 bg-neutral-900 z-0" />
        <div className="relative z-10">
          <Link to="/battle/$gameID" params={params}>
            To battle
          </Link>
        </div>
      </div>
    </ClientOnly>
  )
}

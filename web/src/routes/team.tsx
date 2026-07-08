import { AppHeader } from '#/components/app-header'
import { GothicCard } from '#/components/gothic-ui/card'
import { Platform, PlatformParent } from '#/components/platform'
import { actorsQuery } from '#/lib/queries/get-actors'
import { useQuery } from '@tanstack/react-query'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'

const TEAM_SIZE = 4

export const Route = createFileRoute('/team')({
  component: RouteComponent,
})

function RouteComponent() {
  const actors_query = useQuery(actorsQuery)
  console.log(actors_query.data)
  return (
    <ClientOnly>
      <AppHeader />
      <div className="relative flex flex-col h-full">
        <div className="relative flex h-full">
          <div className="flex-1 px-10 flex flex-col justify-center">
            <div className="relative z-0 flex items-end h-80">
              <PlatformParent className="flex-1 h-full absolute inset-0 z-0 perspective-origin-top">
                {Array.from({ length: TEAM_SIZE }).map((_, i) => (
                  <Platform
                    key={i}
                    rank={i}
                    variant="player"
                    className="flex-1"
                  />
                ))}
              </PlatformParent>
            </div>
          </div>
          <div className="w-1/3 h-full grid place-items-center">
            <GothicCard>test</GothicCard>
          </div>
        </div>
      </div>
    </ClientOnly>
  )
}

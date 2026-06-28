import { ActorDetails } from '#/components/actor-details'
import { AppHeader } from '#/components/app-header'
import { PromptController } from '#/components/prompt-controller'
import { Card, CardContent } from '#/components/ui/card'
import { RenderLog } from '#/lib/game/log'
import { gameStore } from '#/lib/stores/game'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'

export const Route = createFileRoute('/')({ component: Home })

function Home() {
  const game = useSelector(gameStore, (g) => g)
  const actors = useSelector(gameStore, (g) => g.actors)
  return (
    <ClientOnly>
      <PromptController />
      <div>
        <AppHeader />
        <div className="pt-12 flex">
          <div className="p-4 flex flex-col gap-3">
            {actors.map((actor) => (
              <Card key={actor.ID}>
                <CardContent>
                  <ActorDetails actor={actor} />
                </CardContent>
              </Card>
            ))}
          </div>
          <ul className="p-4">
            {game.logs.map((log) => (
              <li key={log.ID}>{RenderLog(log)}</li>
            ))}
          </ul>
        </div>
      </div>
    </ClientOnly>
  )
}

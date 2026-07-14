import { NULL_CONTEXT } from '#/lib/game/context'
import { useUser } from '#/lib/queries/auth'
import { lobbyStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { sendContextMessage } from '#/lib/stores/socket'
import { Link } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
import { GiWingedSword } from 'react-icons/gi'
import { Button } from './ui/button'
import { GothicFramedButton } from './gothic-ui/button'
import { Loader, User } from 'lucide-react'
import { AuthMenu } from './auth-menu'

function AppHeader() {
  const { data: user } = useUser()

  const client = useSelector(lobbyStore, (c) => c.client)
  const game_status = useSelector(gameStore, (g) => g.status)
  const ready = useSelector(gameStore, (g) => g.ready)
  const turn = useSelector(gameStore, (g) => g.turn)

  return (
    <header className="fixed flex justify-between p-1 z-30 w-full">
      <div className="flex items-center gap-2">
        <Button variant="ghost" size="icon" asChild>
          <Link to="/">
            <GiWingedSword />
          </Link>
        </Button>

        {client && (ready || turn > 0) && (
          <div className="flex gap-2">
            <GothicFramedButton
              variant="red"
              disabled={game_status === 'running'}
              onClick={() => {
                sendContextMessage({
                  type: 'turn-ready',
                  client_ID: client.ID,
                  context: NULL_CONTEXT,
                })
              }}
            >
              Run
            </GothicFramedButton>
          </div>
        )}

        <div>
          <span className="slice-loader" />
        </div>

        <div className="flex items-center">
          {(game_status === 'running' || game_status === 'waiting') && (
            <Loader className="animate-spin size-4" />
          )}
        </div>
      </div>

      <div className="flex items-center gap-4">
        <div className="font-mono text-sm flex items-center">
          {user && (
            <AuthMenu asChild>
              <Button variant="ghost" size="icon">
                <User />
              </Button>
            </AuthMenu>
          )}
        </div>
      </div>
    </header>
  )
}

export { AppHeader }

import { NULL_CONTEXT } from '#/lib/game/context'
import { useLogout } from '#/lib/mutations/logout'
import { useUser } from '#/lib/queries/auth'
import { lobbyStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { sendContextMessage } from '#/lib/stores/socket'
import { Link } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
import { GiWhirlpoolShuriken, GiWingedSword } from 'react-icons/gi'
import { Button } from './ui/button'
import { GothicFramedButton } from './gothic-ui/button'

function AppHeader() {
  const { data: user } = useUser()
  const logout = useLogout()
  const client = useSelector(lobbyStore, (c) => c.client)
  const game_status = useSelector(gameStore, (g) => g.status)
  const ready = useSelector(gameStore, (g) => g.ready)
  const turn = useSelector(gameStore, (g) => g.turn)

  return (
    <header className="fixed flex justify-between p-1 z-30 w-full">
      <div className="flex items-center gap-2">
        <Link to="/" className="pl-2">
          <GiWingedSword />
        </Link>

        {client && (ready || turn > 0) && (
          <div className="flex gap-2">
            <GothicFramedButton
              variant="red"
              disabled={game_status === 'running'}
              onClick={() => {
                sendContextMessage({
                  type: 'run-game-actions',
                  client_ID: client.ID,
                  context: NULL_CONTEXT,
                })
              }}
            >
              Run
            </GothicFramedButton>
          </div>
        )}

        <div className="flex items-center">
          {game_status === 'running' && (
            <GiWhirlpoolShuriken className="animate-spin" />
          )}
          {game_status === 'waiting' && (
            <GiWhirlpoolShuriken className="animate-spin" />
          )}
        </div>
      </div>

      <div className="flex items-center gap-4 px-2">
        <div className="font-mono text-sm flex items-center">
          {user && (
            <div className="flex items-center gap-2">
              <Button
                variant="ghost"
                onClick={() => logout.mutate()}
                title="Logout"
              >
                <span className="hidden lg:inline">{user.email}</span>
              </Button>
            </div>
          )}
        </div>
      </div>
    </header>
  )
}

export { AppHeader }

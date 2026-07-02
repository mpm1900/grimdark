import { NULL_CONTEXT } from '#/lib/game/context'
import { useLogout } from '#/lib/mutations/logout'
import { useUser } from '#/lib/queries/auth'
import { connect } from '#/lib/socket/connect'
import { clientsStore } from '#/lib/stores/clients'
import { gameStore } from '#/lib/stores/game'
import { sendContextMessage, socketStore } from '#/lib/stores/socket'
import { Link } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
import { Globe, Loader2, LogOut, TriangleAlert, Unplug } from 'lucide-react'
import { GiWhirlpoolShuriken, GiWingedSword } from 'react-icons/gi'
import { InstanceCombobox } from './instance-combobox'
import { Button } from './ui/button'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'

function AppHeader() {
  const { data: user } = useUser()
  const logout = useLogout()
  const instanceID = useSelector(socketStore, (s) => s.instanceID)
  const status = useSelector(socketStore, (s) => s.status)
  const client = useSelector(clientsStore, (c) => c.me)
  const game_status = useSelector(gameStore, (g) => g.status)
  const turn = useSelector(gameStore, (g) => g.turn)
  return (
    <header className="fixed flex justify-between p-1 z-30 w-full">
      <div className="flex items-center gap-2">
        <Link to="/" className="pl-2">
          <GiWingedSword />
        </Link>

        {user && (
          <InstanceCombobox
            icon={
              <>
                {status === 'idle' && <Unplug />}
                {(status === 'connecting' || status === 'reconnecting') && (
                  <Loader2 className="animate-spin" />
                )}
                {status === 'open' && <Globe />}
                {(status === 'closed' || status === 'error') && (
                  <TriangleAlert className="text-destructive" />
                )}
              </>
            }
            value={instanceID}
            onValueChange={(instanceID) => connect(instanceID)}
          />
        )}

        <span>Turn {turn}</span>

        {client && (
          <div className="flex gap-2">
            <Button
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
            </Button>
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
              <Tooltip>
                <TooltipTrigger>
                  <span className="hidden lg:inline">{user.email}</span>
                </TooltipTrigger>
                <TooltipContent>{user.id}</TooltipContent>
              </Tooltip>
              <Button
                variant="ghost"
                size="icon"
                className="size-8"
                onClick={() => logout.mutate()}
                title="Logout"
              >
                <LogOut className="size-4" />
              </Button>
            </div>
          )}
        </div>
      </div>
    </header>
  )
}

export { AppHeader }

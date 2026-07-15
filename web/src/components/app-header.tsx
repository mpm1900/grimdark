import { useUser } from '#/lib/queries/auth'
import { gameStore } from '#/lib/stores/game'
import { Link } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'
import { GiWingedSword } from 'react-icons/gi'
import { Button } from './ui/button'
import { Loader, User } from 'lucide-react'
import { AuthMenu } from './auth-menu'

function AppHeader() {
  const { data: user } = useUser()

  const game_status = useSelector(gameStore, (g) => g.status)

  return (
    <header className="fixed flex justify-between p-1 z-30 w-full">
      <div className="flex items-center gap-2">
        <Button variant="ghost" size="icon" asChild>
          <Link to="/">
            <GiWingedSword />
          </Link>
        </Button>

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

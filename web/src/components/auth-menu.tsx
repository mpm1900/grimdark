import { LogOut, User } from 'lucide-react'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuShortcut,
  DropdownMenuTrigger,
} from './ui/dropdown-menu'
import { useLogout } from '#/lib/mutations/logout'

function AuthMenu({
  ...props
}: React.ComponentProps<typeof DropdownMenuTrigger>) {
  const logout = useLogout()
  return (
    <DropdownMenu>
      <DropdownMenuTrigger {...props} />
      <DropdownMenuContent>
        <DropdownMenuGroup>
          <DropdownMenuItem>
            My Account
            <DropdownMenuShortcut>
              <User />
            </DropdownMenuShortcut>
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={() => {
              logout.mutate()
            }}
          >
            Log Out
            <DropdownMenuShortcut>
              <LogOut />
            </DropdownMenuShortcut>
          </DropdownMenuItem>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

export { AuthMenu }

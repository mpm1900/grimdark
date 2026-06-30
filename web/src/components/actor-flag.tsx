import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { GothicBadge } from './gothic-ui/badge'
import { Badge } from './ui/badge'

function ActorFlag({
  actor,
  flag,
  variant,
  ...props
}: React.ComponentProps<typeof Badge> & { actor: Actor; flag: keyof Actor }) {
  const value = !!actor[flag]
  return (
    <GothicBadge
      variant={value ? 'light' : 'disabled'}
      className={cn({
        'opacity-50': !value,
      })}
      {...props}
    />
  )
}

export { ActorFlag }

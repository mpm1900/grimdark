import type { ActorStatus } from '#/lib/game/core'
import { cn } from '#/lib/utils'
import { GiFlame } from 'react-icons/gi'
import type { IconType } from 'react-icons/lib'

export const STATUS_ICON: Partial<Record<ActorStatus, IconType>> = {
  burned: ({ className, ...props }) => (
    <GiFlame className={cn(className, 'fill-orange-300')} {...props} />
  ),
}

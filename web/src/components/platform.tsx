import type { ComponentProps } from 'react'
import { cn } from '#/lib/utils'
import { cva, type VariantProps } from 'class-variance-authority'
import type { Ui } from '#/lib/stores/ui'
import type { Player } from '#/lib/game/player'
import type { Status } from '#/lib/game/core'
import { ChevronRight } from 'lucide-react'
import { FaArrowRight, FaChevronRight } from 'react-icons/fa6'
import type { Position } from '#/lib/game/position'

function PlatformParent({
  children,
  className,
  reverse,
  ...props
}: ComponentProps<'div'> & { reverse?: boolean }) {
  return (
    <div
      className={cn(
        'pointer-events-none absolute inset-0 z-0 gap-2 flex flex-row-reverse items-end',
        'perspective-near transform-3d',
        {
          'perspective-origin-[0_50%] flex-row': reverse,
          'perspective-origin-[100%_50%]': !reverse,
        },
        className
      )}
      {...props}
    >
      {children}
    </div>
  )
}

const platformVariants = cva(
  'transform-3d rotate-x-86! -translate-x-1/2 transition-all pointer-events-none absolute bottom-10 left-1/2 size-20 w-full border-6 ring-2 ring-black',
  {
    variants: {
      variant: {
        player: 'bg-emerald-950/20 border-emerald-800/40 text-emerald-800/40',
        'player-active':
          'bg-emerald-900/60 border-emerald-700 text-emerald-700',
        'player-hover': 'bg-emerald-900/40 border-emerald-900 text-emerald-900',
        enemy: 'bg-orange-950/20 border-orange-900/40 text-orange-900/40',
        'enemy-active': 'bg-orange-900/60 border-orange-700 text-orange-700',
        'enemy-hover': 'bg-orange-900/40 border-orange-900 text-orange-900',
        source: 'bg-foreground/20 border-foreground/60 text-foreground/60',
        targeted: 'bg-red-900/40 border-red-800 text-red-800',
      },
    },
  }
)

function getVariant(
  store: Ui,
  client_ID: string,
  position: Position,
  status: Status
): VariantProps<typeof platformVariants>['variant'] {
  if (store.source_actor === position.actor_ID) {
    return 'source'
  }
  if (store.target_positions.includes(position.ID)) {
    return 'targeted'
  }

  if (position.player_ID === client_ID) {
    if (store.range_positions.includes(position.ID)) {
      return 'player-hover'
    }
    if (store.active_actor == position.actor_ID && status === 'idle') {
      return 'player-active'
    }
    if (store.hover_position === position.ID) {
      return 'player-hover'
    }

    return 'player'
  }

  if (position.player_ID !== client_ID) {
    if (store.range_positions.includes(position.ID)) {
      return 'enemy-hover'
    }
    if (store.active_actor == position.actor_ID && status === 'idle') {
      return 'enemy-active'
    }
    if (store.hover_position === position.ID) {
      return 'enemy-hover'
    }

    return 'enemy'
  }
}

function Platform({
  className,
  variant,
  ...props
}: ComponentProps<'div'> & VariantProps<typeof platformVariants>) {
  return (
    <div className="relative flex-1 transform-3d" {...props}>
      <div
        className={platformVariants({
          variant,
          className:
            'flex items-end relative overflow-hidden [&>svg]:opacity-50',
        })}
      >
        {/*
        <FaChevronRight className="absolute top-0 bottom-0 -translate-x-1/2 size-full" />
        <FaChevronRight className="absolute top-0 bottom-0 -translate-x-1/3 size-full" />
        <FaChevronRight className="absolute top-0 bottom-0 -translate-x-1/6 size-full" />
        <FaChevronRight className="absolute top-0 bottom-0 translate-x-1/3 size-full" />
        <FaChevronRight className="absolute top-0 bottom-0 translate-x-1/2 size-full" />
        <FaChevronRight className="absolute top-0 bottom-0 translate-x-1/6 size-full" />
        <FaChevronRight className="absolute top-0 bottom-0 size-full" />
        */}
      </div>
    </div>
  )
}

export { getVariant, Platform, PlatformParent }

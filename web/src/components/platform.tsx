import type { ComponentProps } from 'react'
import { cn } from '#/lib/utils'
import { cva, type VariantProps } from 'class-variance-authority'

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
  'transform-3d rotate-x-86! -translate-x-1/2 pointer-events-none absolute bottom-10 left-1/2 size-20 w-full border-6 ring-2 ring-black',
  {
    variants: {
      variant: {
        player: 'bg-emerald-950/20 border-emerald-800/40',
        'player-active': 'bg-emerald-900/60 border-emerald-700',
        enemy: 'bg-red-950/20 border-red-900/40',
        'enemy-active': 'bg-red-900/60 border-red-700',
      },
    },
  }
)

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
        })}
      />
    </div>
  )
}

export { Platform, PlatformParent }

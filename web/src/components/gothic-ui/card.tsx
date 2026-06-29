import * as React from 'react'

import { cn } from '#/lib/utils'

function GothicCard({ className, ...props }: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="card"
      className={cn(
        'flex flex-col gap-4 py-4 text-card-foreground',
        'border-[10px] border-solid border-transparent bg-neutral-950',
        'bg-clip-border',
        '[image-rendering:pixelated]',
        "[border-image-source:url('/gothic/SkillFrameVert_Dark.png')]",
        '[border-image-slice:22_fill]',
        '[border-image-repeat:stretch]',
        className
      )}
      {...props}
    />
  )
}

function GothicCardContent({
  className,
  ...props
}: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="card-content"
      className={cn('px-4', className)}
      {...props}
    />
  )
}

export { GothicCard, GothicCardContent }

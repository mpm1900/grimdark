import * as React from 'react'

import { cn } from '#/lib/utils'

function GothicCard({
  className,
  children,
  ...props
}: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="card"
      className={cn(
        'flex flex-col relative',
        'border-[10px] border-solid border-transparent bg-gradient-to-b from-neutral-900 to-[oklch(10%_0_0)]',
        'bg-clip-border',
        "[border-image-source:url('/gothic/SkillFrameVert_Dark.png')]",
        '[border-image-slice:22_fill]',
        '[border-image-repeat:stretch]',
        className
      )}
      {...props}
    >
      {children}
    </div>
  )
}

function GothicCardContent({
  className,
  ...props
}: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="card-content"
      className={cn('px-3', className)}
      {...props}
    />
  )
}

export { GothicCard, GothicCardContent }

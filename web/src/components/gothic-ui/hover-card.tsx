import * as React from 'react'
import { HoverCard as HoverCardPrimitive } from 'radix-ui'

import { cn } from '#/lib/utils.ts'

function GothicHoverCardContent({
  className,
  children,
  align = 'center',
  sideOffset = 4,
  ...props
}: React.ComponentProps<typeof HoverCardPrimitive.Content>) {
  return (
    <HoverCardPrimitive.Portal data-slot="hover-card-portal">
      <HoverCardPrimitive.Content
        data-slot="hover-card-content"
        align={align}
        sideOffset={sideOffset}
        className={cn(
          'z-50 w-64 origin-(--radix-hover-card-content-transform-origin) text-popover-foreground shadow-md outline-hidden data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95',
          '[&>span]:z-0',
          'data-[side=top]:[&>span]:mb-1',
          'data-[side=bottom]:[&>span]:mt-1',
          'data-[side=left]:[&>span]:mr-1',
          'data-[side=right]:[&>span]:ml-1',
          className
        )}
        {...props}
      >
        <div
          className={cn(
            'flex flex-col isolate',
            'border-[10px] border-solid border-transparent bg-gradient-to-b from-neutral-900 to-[oklch(10%_0_0)]',
            'bg-clip-border',
            "[border-image-source:url('/gothic/SkillFrameVert_Dark.png')]",
            '[border-image-slice:22_fill]',
            '[border-image-repeat:stretch]'
          )}
        >
          {children}
        </div>
        <HoverCardPrimitive.Arrow
          height={20}
          width={47}
          className="block fill-transparent bg-[url('/gothic/PopoverArrow.png')] bg-contain bg-center bg-no-repeat"
        />
      </HoverCardPrimitive.Content>
    </HoverCardPrimitive.Portal>
  )
}

export { GothicHoverCardContent }

import * as React from 'react'
import { Popover as PopoverPrimitive } from 'radix-ui'

import { cn } from '#/lib/utils.ts'

function GothicPopoverContent({
  className,
  children,
  align = 'center',
  sideOffset = 4,
  ...props
}: React.ComponentProps<typeof PopoverPrimitive.Content>) {
  return (
    <PopoverPrimitive.Portal>
      <PopoverPrimitive.Content
        data-slot="popover-content"
        align={align}
        sideOffset={sideOffset}
        className={cn(
          'relative isolate z-50 w-72 origin-(--radix-popover-content-transform-origin)',
          'text-popover-foreground shadow-md outline-hidden data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95',
          '[&>span]:z-0 [&>span]:translate-y-[-6px]',
          className
        )}
        {...props}
      >
        <div
          className={cn(
            'flex flex-col',
            'border-[10px] border-solid border-transparent bg-gradient-to-b from-neutral-900 to-[oklch(10%_0_0)]',
            'bg-clip-border',
            "[border-image-source:url('/gothic/SkillFrameVert_Dark.png')]",
            '[border-image-slice:22_fill]',
            '[border-image-repeat:stretch]'
          )}
        >
          {children}
        </div>
        <PopoverPrimitive.Arrow
          height={24}
          width={57}
          className="block fill-transparent bg-[url('/gothic/PopoverArrow.png')] bg-contain bg-center bg-no-repeat bottom-3"
        />
      </PopoverPrimitive.Content>
    </PopoverPrimitive.Portal>
  )
}

export { GothicPopoverContent }

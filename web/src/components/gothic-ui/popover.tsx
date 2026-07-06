import * as React from 'react'
import { Popover as PopoverPrimitive } from 'radix-ui'

import { cn } from '#/lib/utils.ts'
import { GothicCloseButton } from './button'

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
          '[&>span]:z-0 [&>span]:translate-y-[-5px]',
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
          <PopoverPrimitive.Close asChild>
            <GothicCloseButton className="absolute z-10 right-0.5 top-0.5 shadow-[0px_0px_6px_1px_rgba(0,_0,_0,_0.5)] z-20">
              <span className="sr-only">Close</span>
            </GothicCloseButton>
          </PopoverPrimitive.Close>
        </div>
        <PopoverPrimitive.Arrow
          height={20}
          width={47}
          className="block fill-transparent bg-[url('/gothic/PopoverArrow.png')] bg-contain bg-center bg-no-repeat bottom-3"
        />
      </PopoverPrimitive.Content>
    </PopoverPrimitive.Portal>
  )
}

function GothicMessage({
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
          ' z-50 relative isolate origin-(--radix-popover-content-transform-origin)',
          'text-popover-foreground outline-hidden data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95',
          '[&>span]:z-0 [&>span]:translate-y-[-5px]',

          className
        )}
        {...props}
      >
        <div className="absolute inset-2 shadow-[0px_8px_24px_5px_rgba(0,_0,_0,_0.4)]" />
        <div
          className={cn(
            'bg-[url(/gothic/ContextFrame.png)] bg-contain bg-no-repeat bg-center',
            'text-foreground/80 font-serif',
            'relative isolate w-96 h-12 leading-12'
          )}
        >
          {children}
        </div>
      </PopoverPrimitive.Content>
    </PopoverPrimitive.Portal>
  )
}

export { GothicPopoverContent, GothicMessage }

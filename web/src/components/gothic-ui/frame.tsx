import * as React from 'react'

import { cn } from '#/lib/utils'

const GothicFrame = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<'div'>
>(({ className, ...props }, ref) => {
  return (
    <div
      ref={ref}
      data-slot="gothic-frame"
      className={cn(
        'flex flex-col p-1.5 text-foreground',
        'border-0 border-solid border-transparent',
        '[border-image-width:10px]',
        'bg-clip-border',
        'bg-neutral-900',
        "[border-image-source:url('/gothic/SkillFrameNormal.png')]",
        '[border-image-slice:22_fill]',
        '[border-image-repeat:stretch]',
        className
      )}
      {...props}
    />
  )
})
GothicFrame.displayName = 'GothicFrame'

const GothicShadowFrame = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<'div'>
>(({ className, ...props }, ref) => {
  return (
    <div
      ref={ref}
      data-slot="gothic-frame"
      className={cn(
        'relative isolate flex flex-col p-2 text-foreground',
        'before:pointer-events-none before:absolute before:inset-0 before:-z-10',
        'before:border-x-[12px] before:border-t-[12px] before:border-b-[28px] before:border-solid before:border-transparent',
        "before:[border-image-source:url('/gothic/BuffShadow.png')]",
        'before:[border-image-slice:12_12_64_12_fill]',
        'before:[border-image-repeat:round]',
        className
      )}
      {...props}
    />
  )
})
GothicShadowFrame.displayName = 'GothicShadowFrame'

const GothicHighlightFrame = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<'div'>
>(({ className, ...props }, ref) => {
  return (
    <div
      ref={ref}
      data-slot="gothic-highlight-frame"
      className={cn(
        'relative isolate flex flex-col p-4 px-6 text-foreground',
        'before:pointer-events-none before:absolute before:inset-0 before:-z-10',
        'before:border-x-[24px] before:border-t-[88px] before:border-b-[88px] before:border-solid before:border-transparent',
        "before:[border-image-source:url('/gothic/MainMenuBar.png')]",
        'before:[border-image-slice:360_122_360_122_fill]',
        'before:[border-image-repeat:stretch]',
        className
      )}
      {...props}
    />
  )
})
GothicHighlightFrame.displayName = 'GothicHighlightFrame'

const GothicChatPanel = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<'div'>
>(({ className, ...props }, ref) => {
  return (
    <div
      ref={ref}
      data-slot="gothic-chat-panel"
      className={cn(
        'relative isolate text-foreground p-1',
        'before:pointer-events-none before:absolute before:inset-0 before:-z-10',
        'before:border-t-[30px] before:border-x-[6px] before:border-b-[10px] before:border-solid before:border-transparent',
        "before:[border-image-source:url('/gothic/SkillBarMiddleSlot.png')]",
        'before:[border-image-slice:25_10_10_10_fill]',
        'before:[border-image-repeat:stretch]',
        className
      )}
      {...props}
    />
  )
})
GothicChatPanel.displayName = 'GothicChatPanel'

export { GothicFrame, GothicShadowFrame, GothicHighlightFrame, GothicChatPanel }

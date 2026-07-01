import { cn } from '#/lib/utils'

function GothicFrame({ className, ...props }: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="gothic-frame"
      className={cn(
        'flex flex-col gap-4 px-1 text-foreground',
        'border-[10px] border-solid border-transparent',
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
}

function GothicShadowFrame({ className, ...props }: React.ComponentProps<'div'>) {
  return (
    <div
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
}


function GothicHighlightFrame({ className, ...props }: React.ComponentProps<'div'>) {
  return (
    <div
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
}

export { GothicFrame, GothicShadowFrame, GothicHighlightFrame }

import { cn } from '#/lib/utils'

function GothicFrame({ className, ...props }: React.ComponentProps<'div'>) {
  return (
    <div
      data-slot="gothic-frame"
      className={cn(
        'flex flex-col gap-4 px-1 text-foreground',
        'border-[10px] border-solid border-transparent bg-neutral-950',
        'bg-clip-border',
        '[image-rendering:pixelated]',
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

export { GothicFrame }

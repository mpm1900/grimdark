import { cn } from '#/lib/utils'

function PanelHeader({ className, ...props }: React.ComponentProps<'div'>) {
  return (
    <div
      className={cn(
        'absolute -top-1.5 -translate-y-2/3 bg-[url(/gothic/TitleFrameMain.png)] bg-contain bg-center bg-no-repeat h-14 w-full z-20',
        'text-center pt-5 font-cinzel font-semibold text-foreground/70',
        className
      )}
      {...props}
    />
  )
}

export { PanelHeader }

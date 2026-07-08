import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'

function ActorLore({
  actor,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor: Actor }) {
  return (
    <div
      className={cn(
        'relative overflow-hidden flex items-center justify-center h-full w-full',
        'border-4 [border-image-source:url(/gothic/WindowPaperLong.png)] [border-image-slice:24_fill] [border-image-repeat:stretch]',
        className
      )}
      {...props}
    >
      <p>lorem ispusm</p>
    </div>
  )
}

export { ActorLore }

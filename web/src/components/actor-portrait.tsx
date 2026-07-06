import type { Actor } from '#/lib/game/actor'
import type { Position } from '#/lib/game/position'
import { cn } from '#/lib/utils'
import { TinyBadge } from './gothic-ui/badge'

function ActorPortrait({
  actor,
  className,
  position,
  ...props
}: React.ComponentProps<'div'> & {
  actor: Actor
  position: Position | undefined
}) {
  return (
    <div className={cn('relative', className)} {...props}>
      <img
        //src="/img/portrait1_.png"
        className="absolute size-18 top-1 left-1 bg-neutral-950"
      />
      <div className="relative bg-[url('/gothic/CharacterTopFrame_Cframe.png')] z-10 size-20 bg-cover">
        <TinyBadge
          variant="default"
          className="absolute top-1.5 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
        >
          Lv {actor.level}
        </TinyBadge>
        {position && (
          <div className="absolute -bottom-2 -left-1 bg-[url('/gothic/MiniIconUIFrame_48.png')] bg-cover z-10 size-6 grid place-items-center text-foreground/70 font-bold font-cinzel">
            {position.rank + 1}
          </div>
        )}
        {actor.status !== 'none' && (
          <TinyBadge
            variant={actor.status as any}
            className="absolute -bottom-0.5 left-2 right-2 text-center capitalize font-cinzel"
          >
            {actor.status}
          </TinyBadge>
        )}
      </div>
    </div>
  )
}
export { ActorPortrait }

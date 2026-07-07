import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { TinyBadge } from './gothic-ui/badge'

function ActorAvatar({ actor }: { actor: Actor }) {
  return (
    <div className="relative h-48 w-52">
      <div className="absolute -bottom-8 -right-5 bg-[url(/gothic/AvatarCircleFrame.png)] bg-contain bg-center size-52">
        <div
          className={cn(
            'rounded-xs absolute -top-6 -left-12 right-5 h-18 font-cinzel-dec text-foreground/80 font-semibold',
            'bg-[url(/gothic/TitleFrameMainRaven_Gray.png)] bg-cover bg-left origin-center pl-22 pb-4 pt-6 -space-y-1.5'
          )}
        >
          <div>{actor.name}</div>
          <div className="text-xs text-foreground/40">{actor.level}</div>
          <div className="absolute -right-5 top-15 flex flex-col gap-px pr-2 text-center capitalize font-cinzel">
            {actor.status !== 'none' && (
              <TinyBadge
                variant={actor.status as any}
                className="pr-3 text-center capitalize font-cinzel"
              >
                {actor.status}
              </TinyBadge>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}

export { ActorAvatar }

import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { TinyBadge } from './gothic-ui/badge'
import { gameStore } from '#/lib/stores/game'
import { getAppliedEffects } from '#/lib/game/game'
import { EffectTooltip } from './effect-tooltip'
import { AffinityIcon } from './affinity-name'

function ActorAvatar({ actor }: { actor: Actor }) {
  const game = useSelector(gameStore, (g) => g)
  const applied_effects = getAppliedEffects(game, actor).sort(
    (a, b) => b.name.length - a.name.length
  )
  const position = game.positions.find((p) => p.actor_ID === actor.ID)
  return (
    <div className="relative h-48 w-52">
      <div className="absolute size-42 left-6 mt-13 rounded-full overflow-hidden bg-linear-to-b from-emerald-950/0 to-emerald-950">
        <div
          className={cn(
            'absolute w-56 h-56 -left-6 top-0 bg-no-repeat bg-cover bg-top'
          )}
          style={{ backgroundImage: `url(${actor.sprite_url})` }}
        />
      </div>
      <div className="absolute -bottom-9 -right-1 bg-[url(/gothic/AvatarCircleFrame.png)] bg-contain bg-center size-52">
        <div
          className={cn(
            'rounded-xs absolute z-10 -top-7 -left-12 right-1 h-18 font-cinzel-dec text-foreground/80 font-semibold',
            'bg-[url(/gothic/TitleFrameMainRaven_Gray.png)] bg-cover bg-left origin-center pl-22 pb-4 pt-6 -space-y-1.5'
          )}
        >
          <div className="truncate">{actor.name}</div>
          <div className="text-xs text-foreground/40">{actor.level}</div>
          <div className="absolute -right-5 top-15 flex flex-col items-end gap-px pr-2 text-center capitalize font-cinzel">
            {applied_effects.map((effect) => (
              <EffectTooltip key={effect.ID} effect={effect} asChild>
                <TinyBadge className="pr-3 text-center capitalize font-cinzel">
                  {effect.name}
                  {effect.count > 1 && `(${effect.count})`}
                </TinyBadge>
              </EffectTooltip>
            ))}
          </div>
          {actor.affinities[0] && (
            <div className="absolute -bottom-4 left-8.5 bg-[url('/gothic/MiniIconUIFrame_48.png')] bg-cover size-10 grid place-items-center text-3xl text-foreground/70 font-bold font-cinzel">
              <AffinityIcon affinity={actor.affinities[0]} className="size-5" />
            </div>
          )}
          {actor.affinities[1] && (
            <div className="absolute -bottom-14 left-5 bg-[url('/gothic/MiniIconUIFrame_48.png')] bg-cover size-10 grid place-items-center text-3xl text-foreground/70 font-bold font-cinzel">
              <AffinityIcon affinity={actor.affinities[1]} className="size-5" />
            </div>
          )}
        </div>
        {position && (
          <div className="absolute bottom-7 -right-1 bg-[url('/gothic/MiniIconUIFrame_48.png')] bg-cover size-12 grid place-items-center text-3xl text-foreground/70 font-bold font-cinzel">
            {position.rank + 1}
          </div>
        )}
      </div>
    </div>
  )
}

export { ActorAvatar }

import type { Actor } from '#/lib/game/actor'
import {
  ACCURACY_STATS,
  AFFINITIES,
  MAIN_STATS,
  mapStage,
  STAT_SLUGS,
  type Stat,
} from '#/lib/game/core'
import { AffinityName } from './affinity-name'
import { TinyBadge } from './gothic-ui/badge'
import { HealthBar } from './health-bar'

function StatMultBadge({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  let mod = 2
  const mult = mapStage(stage, mod, 1)
  if (mult == 1) {
    return null
  }
  return (
    <TinyBadge
      className="uppercase font-cinzel shadow-[1px_1px_0_var(--color-black)]"
      variant={mult > 1 ? 'positive' : 'negative'}
    >
      {STAT_SLUGS[stat]} {mult.toFixed(2)}
    </TinyBadge>
  )
}

function ActorFrame({ actor }: { actor: Actor }) {
  return (
    <div className="flex -ml-1 -mt-1">
      <div className="bg-[url('/gothic/CharacterTopFrame_Cframe.png')] z-10 size-20 bg-cover"></div>
      <div className="flex-1 flex flex-col -ml-1 mt-2.5 pr-1 bg-gradient-to-b from-foreground/20 to-black/0">
        <div className="h-1 bg-white/30 border -mr-1.5 border-black mb-1" />
        <div className="flex flex-row justify-between items-center pr-1 pb-1">
          <span className="text-sm leading-0 font-bold font-cinzel-dec [text-shadow:2px_1px_0_var(--color-black)] text-foreground px-2 -mb-1">
            {actor.name}
          </span>
          <div className="flex gap-1">
            {AFFINITIES.filter((a) => actor.affinities.includes(a)).map((a) => (
              <AffinityName key={a} affinity={a} />
            ))}
          </div>
        </div>
        <div className="relative">
          <HealthBar
            actor={actor}
            type="value"
            className="rounded-l-none -ml-px"
          />
          <div className="absolute -bottom-2 left-1 flex gap-0.5">
            {[...MAIN_STATS, ...ACCURACY_STATS].map((stat) => (
              <StatMultBadge actor={actor} stat={stat} />
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}

export { ActorFrame }

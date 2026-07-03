import type { Actor } from '#/lib/game/actor'
import {
  ACCURACY_STATS,
  AFFINITIES,
  MAIN_STATS,
  mapStage,
  STAT_SLUGS,
  type Stat,
} from '#/lib/game/core'
import { gameStore } from '#/lib/stores/game'
import { useSelector } from '@tanstack/react-store'
import { AffinityName } from './affinity-name'
import { TinyBadge } from './gothic-ui/badge'
import { HealthBar } from './health-bar'
import { cn } from '#/lib/utils'
import { Popover, PopoverTrigger } from './ui/popover'
import { GothicPopoverContent } from './gothic-ui/popover'
import { ActorStatsPanel } from './panels/actor-stats'
import { ActorPortrait } from './actor-portrait'
import { clientsStore } from '#/lib/stores/clients'

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

function ActorFrame({
  actor,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor: Actor }) {
  const player = useSelector(gameStore, (g) =>
    g.players.find((p) => p.ID === actor.player_ID)
  )
  const position = player?.positions.find((p) => p.ID === actor.position_ID)
  return (
    <div className={cn('relative flex', className)} {...props}>
      <ActorPortrait actor={actor} position={position} />
      <div className="relative flex-1 flex flex-col -ml-1 mt-[7px] pr-1 bg-[url('/gothic/DialogFlag_stone_shadow.png')] bg-[length:100%_100%] bg-center bg-no-repeat">
        <div className="h-1 bg-gradient-to-b from-white/40 to-neutral-800/60 border -mr-1 border-black mb-1" />
        <div className="pointer-events-none absolute inset-x-0 top-0 -z-10 h-24 overflow-hidden">
          <div className="absolute left-1/2 h-full w-[calc(100%+5rem)] -translate-x-1/2 opacity-70" />
        </div>
        <div className="flex flex-row justify-between items-center pr-1 pb-1">
          <span className="text-sm leading-0 font-bold font-cinzel-dec [text-shadow:2px_1px_0_var(--color-black)] text-foreground px-2 -mb-1">
            {actor.name}
          </span>
          <div className="flex gap-0">
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
          <div className="absolute -bottom-1.5 left-1 flex gap-0.5">
            {[...MAIN_STATS, ...ACCURACY_STATS].map((stat) => (
              <StatMultBadge key={stat} actor={actor} stat={stat} />
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}

function ActorFrameSlim({
  actor,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor: Actor }) {
  const client = useSelector(clientsStore, (s) => s.me)
  return (
    <div className={cn('relative mt-4', className)} {...props}>
      <div className="flex flex-row justify-between items-end mb-1">
        <Popover>
          <PopoverTrigger
            className="cursor-pointer -mr-1"
            onClick={(e) => {
              e.stopPropagation()
            }}
          >
            <span className="text-sm font-bold font-cinzel-dec [text-shadow:2px_1px_0_var(--color-black)] text-foreground -mb-1 hover:underline">
              {actor.name}
            </span>
          </PopoverTrigger>
          <GothicPopoverContent
            className="w-auto"
            side={'top'}
            align="center"
            collisionPadding={16}
          >
            <ActorStatsPanel actor={actor} />
          </GothicPopoverContent>
        </Popover>

        <div className="flex gap-0"></div>
      </div>
      <div className="relative">
        <HealthBar
          actor={actor}
          type="value"
          className="rounded-l-none -ml-px"
          hide_numbers={client?.ID !== actor.player_ID}
        />
        <div className="absolute -bottom-1.5 left-0 flex gap-0.5">
          {[...MAIN_STATS, ...ACCURACY_STATS].map((stat) => (
            <StatMultBadge key={stat} actor={actor} stat={stat} />
          ))}
        </div>
      </div>
    </div>
  )
}

export { ActorFrame, ActorFrameSlim }

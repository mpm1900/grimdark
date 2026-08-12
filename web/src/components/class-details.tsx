import type { ActorClass } from '#/lib/game/actor-class'
import { cn } from '#/lib/utils'
import { FaCircle } from 'react-icons/fa'
import { ActionTooltip } from './action-tooltip'
import { AffinityIcon } from './affinity-name'
import { StatIcon, statVariants } from './stat-name'
import { ItemDescription } from './ui/item'
import { EffectTooltip } from './effect-tooltip'
import { CLASS_STATS } from '#/lib/game/core'

function ClassDetails({
  actor_class,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor_class?: ActorClass }) {
  if (!actor_class) return null
  return (
    <div className={cn(className, 'font-serif p-1')} {...props}>
      <div className="font-cinzel-dec font-semibold mb-2">
        {actor_class.name}
      </div>
      <div className="flex gap-1">
        {CLASS_STATS.map((stat) => (
          <div key={stat} className="flex-1 flex flex-col items-center gap-0">
            <StatIcon stat={stat} className="size-5" />
            <div>{actor_class.stats[stat]}</div>
          </div>
        ))}
      </div>
      <ItemDescription className="text-foreground/80">
        <span className="text-foreground/40 block font-cinzel font-semibold">
          Actions
        </span>
        <span className="space-x-2 flex flex-wrap">
          {actor_class.actions.map((a, i) => (
            <ActionTooltip
              key={`${a.ID}-${a.entity.name}-${i}`}
              action={a}
              className={statVariants({
                stat: a.config.stat,
                className:
                  'cursor-default hover:underline capitalize flex items-center gap-0.5',
              })}
            >
              <AffinityIcon affinity={a.config.affinity}>
                <FaCircle className="size-2" />
              </AffinityIcon>
              {a.entity.name}
              {a.tags.includes('conditional') && '*'}
            </ActionTooltip>
          ))}
        </span>
        {actor_class.actions.length == 0 && <span>None</span>}
      </ItemDescription>
      <ItemDescription className="text-foreground/80">
        <span className="text-foreground/40 block font-cinzel font-semibold">
          Effects
        </span>
        {actor_class.effects.length > 0 && (
          <span className="space-x-2 flex flex-wrap">
            {actor_class.effects.map((e, i) => (
              <EffectTooltip
                key={`${e.ID}-${e.entity.name}-${i}`}
                effect={e}
                className="cursor-default hover:underline capitalize"
              >
                {e.entity.name}
              </EffectTooltip>
            ))}
          </span>
        )}
        {actor_class.effects.length == 0 && <span>None</span>}
      </ItemDescription>
    </div>
  )
}

export { ClassDetails }

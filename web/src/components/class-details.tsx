import type { ActorClass } from '#/lib/game/actor-class'
import { cn } from '#/lib/utils'
import { FaCircle } from 'react-icons/fa'
import { ActionTooltip } from './action-tooltip'
import { AffinityIcon } from './affinity-name'
import { statVariants } from './stat-name'
import { ItemDescription } from './ui/item'
import { EffectTooltip } from './effect-tooltip'

function ClassDetails({
  actor_class,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor_class?: ActorClass }) {
  if (!actor_class) return null
  return (
    <div className={cn(className)} {...props}>
      <div>{actor_class.name}</div>
      <ItemDescription className="text-foreground/80">
        <span className="text-foreground/40 block font-cinzel font-semibold">
          Actions
        </span>
        <span className="space-x-2 flex flex-wrap">
          {actor_class.actions.map((a) => (
            <ActionTooltip
              key={a.ID}
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
              {a.config.name}
              {a.tags.includes('conditional') && '*'}
            </ActionTooltip>
          ))}
        </span>
      </ItemDescription>
      <ItemDescription className="text-foreground/80">
        <span className="text-foreground/40 block font-cinzel font-semibold">
          Effects
        </span>
        {actor_class.effects.length > 0 && (
          <span className="space-x-2 flex flex-wrap">
            {actor_class.effects.map((e) => (
              <EffectTooltip
                key={e.ID}
                effect={e}
                className="cursor-default hover:underline capitalize"
              >
                {e.name}
              </EffectTooltip>
            ))}
          </span>
        )}
      </ItemDescription>
    </div>
  )
}

export { ClassDetails }

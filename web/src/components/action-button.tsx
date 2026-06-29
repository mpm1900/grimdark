import type { Action } from '#/lib/game/action'
import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { ArrowRight } from 'lucide-react'
import { AffinityName, affinityVariants } from './affinity-name'
import { Button } from './ui/button'
import {
  Item,
  ItemActions,
  ItemContent,
  ItemDescription,
  ItemMedia,
  ItemTitle,
} from './ui/item'
import { StatName } from './stat-name'

function ActionButton({
  action,
  actor,
  ...props
}: React.ComponentProps<typeof Button> & { action: Action; actor: Actor }) {
  return (
    <Item className="p-2" asChild>
      <Button
        {...props}
        variant="outline"
        className="h-auto"
        disabled={action.is_disabled || !actor.is_active}
      >
        {action.config.affinity && (
          <ItemMedia>
            <AffinityName affinity={action.config.affinity} />
          </ItemMedia>
        )}
        <ItemContent>
          <ItemTitle>{action.config.name}</ItemTitle>
          <ItemDescription className="text-left">
            {action.config.description || (
              <span className="flex gap-2">
                {action.config.stat && (
                  <StatName stat={action.config.stat} className="capitalize">
                    {action.config.stat}
                  </StatName>
                )}
                {action.config.accuracy && (
                  <span
                    className={cn({
                      'text-green-400': actor.stats['accuracy'] > 100,
                      'text-red-400': actor.stats['accuracy'] < 100,
                    })}
                  >
                    {Math.min(action.config.accuracy * 100, 100)}%
                  </span>
                )}
                {action.config.crit_chance && (
                  <span className={cn('flex items-center')}>
                    <span
                      className={cn({
                        'text-green-400': actor.stats['critical-chance'] > 100,
                        'text-red-400': actor.stats['critical-chance'] < 100,
                      })}
                    >
                      {(action.config.crit_chance * 100).toFixed(0)}%
                    </span>
                    <ArrowRight />
                    <span
                      className={cn({
                        'text-green-400': actor.stats['critical-damage'] > 100,
                        'text-red-400': actor.stats['critical-damage'] < 100,
                      })}
                    >
                      x{action.config.crit_modifier.toFixed(2)}
                    </span>
                  </span>
                )}
              </span>
            )}
          </ItemDescription>
        </ItemContent>
        {!!action.config.power && (
          <ItemActions className="flex flex-col h-full items-start">
            <span
              className={cn(
                'text-xl font-black',
                affinityVariants({
                  affinity: action.config.affinity,
                })
              )}
            >
              {action.config.power}
            </span>
          </ItemActions>
        )}
      </Button>
    </Item>
  )
}

export { ActionButton }

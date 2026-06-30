import type { Action } from '#/lib/game/action'
import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { MdKeyboardDoubleArrowRight } from 'react-icons/md'
import { AffinityName, affinityVariants } from './affinity-name'
import { Button } from './ui/button'
import {
  ItemActions,
  ItemContent,
  ItemDescription,
  ItemMedia,
  ItemTitle,
} from './ui/item'
import { StatName } from './stat-name'
import { GothicFramedButton } from './gothic-ui/button'
import { STAT_LABELS } from '#/lib/game/core'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'

function ActionButton({
  action,
  actor,
  ...props
}: React.ComponentProps<typeof Button> & { action: Action; actor: Actor }) {
  const status = useSelector(gameStore, (g) => g.status)
  return (
    <GothicFramedButton
      {...props}
      variant="basic"
      className="h-auto"
      disabled={action.is_disabled || !actor.is_active || status === 'running'}
    >
      {action.config.affinity && (
        <ItemMedia>
          <AffinityName affinity={action.config.affinity} />
        </ItemMedia>
      )}
      <ItemContent className="gap-0 py-0.5">
        <ItemTitle
          className={cn('text-white', action.is_disabled && 'text-white/60')}
        >
          {action.config.name}
        </ItemTitle>
        <ItemDescription className="text-left font-serif text-foreground/70">
          {!action.config.power && action.config.description}
          {!!action.config.power && (
            <span className="flex gap-2 font-cinzel text-white/50">
              {action.config.stat && (
                <StatName stat={action.config.stat} className="font-serif">
                  {STAT_LABELS[action.config.stat]}
                </StatName>
              )}
              <span className="flex gap-2 font-semibold">
                {action.config.accuracy && (
                  <span
                    className={cn({
                      'text-positive': actor.stats['accuracy'] > 100,
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
                        'text-positive': actor.stats['critical-chance'] > 100,
                        'text-red-400': actor.stats['critical-chance'] < 100,
                      })}
                    >
                      {(action.config.crit_chance * 100).toFixed(0)}%
                    </span>
                    <MdKeyboardDoubleArrowRight />
                    <span
                      className={cn({
                        'text-positive': actor.stats['critical-damage'] > 100,
                        'text-red-400': actor.stats['critical-damage'] < 100,
                      })}
                    >
                      x{action.config.crit_modifier.toFixed(2)}
                    </span>
                  </span>
                )}
              </span>
            </span>
          )}
        </ItemDescription>
      </ItemContent>
      {!!action.config.power && (
        <ItemActions className="flex flex-col h-full items-start">
          <span
            className={cn(
              'text-xl font-black font-cinzel-dec',
              affinityVariants({
                affinity: action.config.affinity,
              })
            )}
          >
            {action.config.power}
          </span>
        </ItemActions>
      )}
    </GothicFramedButton>
  )
}

export { ActionButton }

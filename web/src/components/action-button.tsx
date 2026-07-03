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
import { StatIcon } from './stat-name'
import { GothicFramedButton } from './gothic-ui/button'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { useQuery } from '@tanstack/react-query'
import { NULL_CONTEXT } from '#/lib/game/context'
import { setRangePositions } from '#/lib/stores/ui'

function ActionButton({
  action,
  actor,
  disabled,
  ...props
}: React.ComponentProps<typeof Button> & { action: Action; actor: Actor }) {
  const turn = useSelector(gameStore, (g) => g.turn)
  const status = useSelector(gameStore, (g) => g.status)
  const targets_options = getTargetsQuery(
    actor?.ID,
    actor?.player_ID,
    action.ID,
    [turn]
  )
  targets_options.enabled = !disabled
  const targets_query = useQuery(targets_options)
  const targets_context = targets_query.data ?? NULL_CONTEXT

  return (
    <GothicFramedButton
      {...props}
      variant="basic"
      className="relative h-auto w-full min-w-0 justify-start overflow-hidden"
      disabled={
        disabled ||
        action.is_disabled ||
        !actor.is_active ||
        status === 'running'
      }
      onMouseEnter={() => {
        setRangePositions(targets_context.position_IDs)
      }}
      onMouseLeave={() => {
        setRangePositions([])
      }}
    >
      {action.config.stat && (
        <StatIcon
          stat={action.config.stat}
          className="size-15 absolute opacity-20 -right-6 bottom-0"
        />
      )}
      {action.config.affinity && (
        <ItemMedia>
          <AffinityName affinity={action.config.affinity} />
        </ItemMedia>
      )}
      <ItemContent className="gap-0 py-0.5 min-w-0 overflow-hidden">
        <ItemTitle
          className={cn(
            'text-white gap-1',
            action.is_disabled && 'text-white/60'
          )}
        >
          {action.config.name}
        </ItemTitle>
        <ItemDescription className="block min-w-0 max-w-full overflow-hidden text-ellipsis whitespace-nowrap text-left font-serif text-foreground/70">
          {!action.config.power && (
            <span className="block truncate">
              {action.is_disabled && (
                <span className="text-red-300/40">[Disabled] </span>
              )}
              {action.config.description}
            </span>
          )}
          {!!action.config.power && (
            <span className="font-cinzel text-white/50">
              {action.is_disabled && (
                <span className="text-red-300/40">[Disabled] </span>
              )}
              <span className="mr-2 font-semibold">
                <span className="font-serif font-normal text-foreground/70">
                  {action.config.description}
                </span>
                {action.config.crit_chance && (
                  <span className="inline-flex items-baseline align-baseline">
                    <span
                      className={cn({
                        'text-positive': actor.stats['critical-chance'] > 100,
                        'text-negative': actor.stats['critical-chance'] < 100,
                      })}
                    >
                      {(action.config.crit_chance * 100).toFixed(0)}%
                    </span>
                    <MdKeyboardDoubleArrowRight className="self-center" />
                    <span
                      className={cn({
                        'text-positive': actor.stats['critical-damage'] > 100,
                        'text-negative': actor.stats['critical-damage'] < 100,
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
        <ItemActions className="relative flex flex-col gap-0 h-full items-end -mt-1 font-cinzel">
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
          {action.config.accuracy && (
            <span
              className={cn('text-foreground/80 text-xs font-semibold', {
                'text-positive/80': actor.stats['accuracy'] > 100,
                'text-negative/80': actor.stats['accuracy'] < 100,
              })}
            >
              {Math.min(action.config.accuracy * 100, 100)}%
            </span>
          )}
        </ItemActions>
      )}
    </GothicFramedButton>
  )
}

export { ActionButton }

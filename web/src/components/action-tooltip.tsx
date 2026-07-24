import type { Action } from '#/lib/game/action'
import { cn } from '#/lib/utils'
import { AffinityIcon } from './affinity-name'
import { GothicHoverCardContent } from './gothic-ui/hover-card'
import { statVariants } from './stat-name'
import { HoverCard, HoverCardContent, HoverCardTrigger } from './ui/hover-card'

function ActionTooltip({
  action,
  card_content = {},
  ...props
}: React.ComponentProps<typeof HoverCardTrigger> & {
  action: Action
  card_content?: Partial<React.ComponentProps<typeof HoverCardContent>>
}) {
  return (
    <HoverCard>
      <HoverCardTrigger className="cursor-default" {...props} />
      <GothicHoverCardContent
        {...card_content}
        className={cn('font-serif', card_content?.className)}
      >
        <div className="absolute top-2 right-2 overflow-hidden size-12">
          <AffinityIcon
            affinity={action.config.affinity}
            className="size-12 opacity-50"
          />
        </div>
        <div
          className={statVariants({
            stat: action.config.stat,
            className: 'font-semibold text-lg text-foreground/80 px-1',
          })}
        >
          {action.config.name}
        </div>
        <div className="flex flex-col px-1 gap-0 pr-14 text-sm leading-4">
          {!!action.config.power && (
            <div>
              <span className="text-foreground/50 text-xs">Power:</span>{' '}
              <span className="font-bold">{action.config.power}</span>
            </div>
          )}
          {action.config.accuracy && (
            <div>
              <span className="text-foreground/50 text-xs">Accuracy:</span>{' '}
              <span className="font-bold">
                {action.config.accuracy * 100 + '%'}
              </span>
            </div>
          )}
        </div>
        <div className="text-xs p-1">
          {action.config.description}
          {action.config.range && (
            <span>
              {' '}
              Range of {action.config.range} tile
              {action.config.range != 1 && 's'}.
            </span>
          )}
          {action.config.cooldown > 0 && (
            <span>{` ${action.config.cooldown} turn cooldown.`}</span>
          )}
          {!!action.config.uses && (
            <span>
              {' '}
              {action.config.uses - action.uses} of {action.config.uses} uses
              remaining.
            </span>
          )}
        </div>
      </GothicHoverCardContent>
    </HoverCard>
  )
}

export { ActionTooltip }

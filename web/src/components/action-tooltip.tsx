import type { Action } from '#/lib/game/action'
import { cn } from '#/lib/utils'
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
        <div
          className={statVariants({
            stat: action.config.stat,
            className: 'font-semibold text-foreground/80 px-1',
          })}
        >
          {action.config.name}
        </div>
        <div className="grid grid-cols-2 text-center">
          <span>{action.config.power || '-'}</span>
          <span>
            {action.config.accuracy ? action.config.accuracy * 100 + '%' : '-'}
          </span>
        </div>
        <div className="text-xs p-1">
          {action.config.description}
          {action.config.range && (
            <span> Range of {action.config.range} tile.</span>
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

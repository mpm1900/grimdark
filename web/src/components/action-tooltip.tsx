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
            className: 'font-semibold text-foreground/80',
          })}
        >
          {action.config.name}
        </div>
        <div className="grid grid-cols-3 text-center">
          <span>
            {action.config.accuracy ? action.config.accuracy * 100 + '%' : '-'}
          </span>
          <span>{action.config.power || '-'}</span>
          <span>{action.config.range ?? '-'}</span>
        </div>
        <div className="text-xs">
          {action.config.description}
          {action.config.range && (
            <span> Range of {action.config.range} tile.</span>
          )}
          {action.config.cooldown > 0 && (
            <span>{` ${action.config.cooldown} turn cooldown.`}</span>
          )}
        </div>
      </GothicHoverCardContent>
    </HoverCard>
  )
}

export { ActionTooltip }

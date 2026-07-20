import type { Effect } from '#/lib/game/effect'
import { cn } from '#/lib/utils'
import { GothicHoverCardContent } from './gothic-ui/hover-card'
import { HoverCard, HoverCardContent, HoverCardTrigger } from './ui/hover-card'

function EffectTooltip({
  effect,
  card_content = {},
  ...props
}: React.ComponentProps<typeof HoverCardTrigger> & {
  effect: Effect
  card_content?: Partial<React.ComponentProps<typeof HoverCardContent>>
}) {
  return (
    <HoverCard>
      <HoverCardTrigger className="cursor-default" {...props} />
      <GothicHoverCardContent
        {...card_content}
        className={cn('font-serif', card_content?.className)}
      >
        <div className="font-semibold text-foreground/80 capitalize px-1">
          {effect.name}
        </div>
        <div className="text-xs p-1">
          {effect.description}
          {!!effect.duration && (
            <span className="text-xs p-1">
              {' '}
              {effect.duration} turn duration.
            </span>
          )}
        </div>
      </GothicHoverCardContent>
    </HoverCard>
  )
}

export { EffectTooltip }

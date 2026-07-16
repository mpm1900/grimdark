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
  console.log(effect)
  return (
    <HoverCard>
      <HoverCardTrigger className="cursor-default" {...props} />
      <GothicHoverCardContent
        {...card_content}
        className={cn('font-serif', card_content?.className)}
      >
        <div className="font-semibold text-foreground/80">{effect.name}</div>
        <div className="text-xs">{effect.description}</div>
      </GothicHoverCardContent>
    </HoverCard>
  )
}

export { EffectTooltip }

import type { Effect } from '#/lib/game/effect'
import { cn } from '#/lib/utils'
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
      <HoverCardContent
        {...card_content}
        className={cn('font-serif', card_content?.className)}
      >
        <div>{effect.name}</div>
        <div>{effect.description}</div>
      </HoverCardContent>
    </HoverCard>
  )
}

export { EffectTooltip }

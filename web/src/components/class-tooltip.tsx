import type { ActorClass } from '#/lib/game/actor-class'
import { cn } from '#/lib/utils'
import { ClassDetails } from './class-details'
import { GothicHoverCardContent } from './gothic-ui/hover-card'
import { HoverCard, HoverCardTrigger } from './ui/hover-card'

function ClassTooltip({
  actor_class,
  className,
  ...props
}: React.ComponentProps<typeof HoverCardTrigger> & {
  actor_class?: ActorClass
}) {
  return (
    <HoverCard>
      <HoverCardTrigger className={cn(className)} {...props} />
      <GothicHoverCardContent sideOffset={0} side="left">
        <ClassDetails actor_class={actor_class} />
      </GothicHoverCardContent>
    </HoverCard>
  )
}

export { ClassTooltip }

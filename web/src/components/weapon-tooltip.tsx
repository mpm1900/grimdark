import type { Weapon } from '#/lib/game/weapon'
import { cn } from '#/lib/utils'
import { GothicHoverCardContent } from './gothic-ui/hover-card'
import { HoverCard, HoverCardTrigger } from './ui/hover-card'
import { WeaponDetails } from './weapon-details'

function WeaponTooltip({
  weapon,
  content_props = {},
  className,
  hover_card = {},
  ...props
}: React.ComponentProps<typeof HoverCardTrigger> & {
  weapon?: Weapon
  content_props?: Partial<React.ComponentProps<typeof GothicHoverCardContent>>
  hover_card?: Partial<React.ComponentProps<typeof HoverCard>>
}) {
  return (
    <HoverCard {...hover_card}>
      <HoverCardTrigger
        className={cn(weapon && 'hover:underline cursor-default', className)}
        {...props}
      />
      <GothicHoverCardContent
        sideOffset={0}
        side="left"
        {...content_props}
        className={cn('w-80', content_props.className)}
      >
        {weapon && <WeaponDetails weapon={weapon} />}
      </GothicHoverCardContent>
    </HoverCard>
  )
}

export { WeaponTooltip }

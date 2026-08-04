import type { Item } from '#/lib/game/weapon'
import { cn } from '#/lib/utils'
import { GothicHoverCardContent } from './gothic-ui/hover-card'
import { ItemDetails } from './item-details'
import { HoverCard, HoverCardTrigger } from './ui/hover-card'

function ItemTooltip({
  item,
  content_props = {},
  className,
  hover_card = {},
  ...props
}: React.ComponentProps<typeof HoverCardTrigger> & {
  item?: Item
  content_props?: Partial<React.ComponentProps<typeof GothicHoverCardContent>>
  hover_card?: Partial<React.ComponentProps<typeof HoverCard>>
}) {
  return (
    <HoverCard {...hover_card}>
      <HoverCardTrigger
        className={cn(item && 'hover:underline cursor-default', className)}
        {...props}
      />
      <GothicHoverCardContent
        sideOffset={0}
        side="left"
        {...content_props}
        className={cn('w-80', content_props.className)}
      >
        {item && <ItemDetails item={item} />}
      </GothicHoverCardContent>
    </HoverCard>
  )
}

export { ItemTooltip }

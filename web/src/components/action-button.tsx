import type { Action } from '#/lib/game/action'
import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { AffinityName, affinityVariants } from './affinity-name'
import { Button } from './ui/button'
import {
  Item,
  ItemActions,
  ItemContent,
  ItemDescription,
  ItemMedia,
  ItemTitle,
} from './ui/item'

function ActionButton({
  action,
  actor,
  ...props
}: React.ComponentProps<typeof Button> & { action: Action; actor: Actor }) {
  return (
    <Item className="p-2" asChild>
      <Button
        {...props}
        variant="outline"
        className="h-auto"
        disabled={action.is_disabled || !actor.is_active}
      >
        {action.config.affinity && (
          <ItemMedia>
            <AffinityName affinity={action.config.affinity} />
          </ItemMedia>
        )}
        <ItemContent>
          <ItemTitle>{action.config.name}</ItemTitle>
          <ItemDescription className="text-left">
            {action.config.description || (
              <span className="flex gap-2">
                {action.config.accuracy && (
                  <span>Accuracy {action.config.accuracy * 100}%</span>
                )}
              </span>
            )}
          </ItemDescription>
        </ItemContent>
        {!!action.config.power && (
          <ItemActions>
            <span
              className={cn(
                'text-xl font-black',
                affinityVariants({
                  affinity: action.config.affinity,
                })
              )}
            >
              {action.config.power}
            </span>
          </ItemActions>
        )}
      </Button>
    </Item>
  )
}

export { ActionButton }

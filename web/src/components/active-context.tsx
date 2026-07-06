import type { Context } from '#/lib/game/context'
import { useSelector } from '@tanstack/react-store'
import { GothicMessage } from './gothic-ui/popover'
import { Popover, PopoverAnchor } from './ui/popover'
import { gameStore } from '#/lib/stores/game'

function ActiveContext({
  active_context,
  ...props
}: React.ComponentProps<typeof PopoverAnchor> & {
  active_context: Context | null
}) {
  const actors = useSelector(gameStore, (g) => g.actors)
  const modifiers = useSelector(gameStore, (g) => g.modifiers)
  const source = actors.find((a) => a.ID === active_context?.source_ID)
  const parent = actors.find((a) => a.ID === active_context?.parent_ID)
  const action = source?.actions.find((a) => a.ID === active_context?.action_ID)
  const modifier = modifiers.find(
    (m) => m.payload.ID === active_context?.effect_ID
  )
  console.log(active_context)
  return (
    <Popover open={!!active_context && !!source}>
      <PopoverAnchor {...props} />
      <GothicMessage side="bottom">
        {!!action && (
          <div>
            {source?.name} used {action.config.name}
          </div>
        )}
        {!!modifier && (
          <div>
            {parent?.name}'s {modifier.payload.name} trigger
          </div>
        )}
      </GothicMessage>
    </Popover>
  )
}

export { ActiveContext }

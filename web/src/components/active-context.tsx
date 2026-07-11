import type { Context } from '#/lib/game/context'
import { useSelector } from '@tanstack/react-store'
import { GothicMessage } from './gothic-ui/popover'
import { Popover, PopoverAnchor } from './ui/popover'
import { gameStore } from '#/lib/stores/game'
import { Loader } from 'lucide-react'

function ActiveContext({
  active_context,
  ...props
}: React.ComponentProps<typeof PopoverAnchor> & {
  active_context: Context | null
}) {
  const actors = useSelector(gameStore, (g) => g.actors)
  const modifiers = useSelector(gameStore, (g) => g.modifiers)
  const status = useSelector(gameStore, (g) => g.status)
  const prompts = useSelector(gameStore, (g) => g.prompts.length)
  const source = actors.find((a) => a.ID === active_context?.source_ID)
  const parent = actors.find((a) => a.ID === active_context?.parent_ID)
  const action = source?.actions.find((a) => a.ID === active_context?.action_ID)
  const modifier = modifiers.find(
    (m) => m.payload.ID === active_context?.effect_ID
  )
  const context_open = !!active_context && (!!source || !!modifier)
  const waiting = status === 'waiting' && prompts === 0

  return (
    <Popover key={source?.ID ?? modifier?.ID} open={context_open || waiting}>
      <PopoverAnchor {...props} />
      <GothicMessage
        side="bottom"
        className="text-center grid place-items-center"
      >
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
        {waiting && (
          <div className="flex gap-2 items-center justify-center">
            <div>Waiting on another player</div>
            <Loader className="size-4 animate-spin" />
          </div>
        )}
      </GothicMessage>
    </Popover>
  )
}

export { ActiveContext }

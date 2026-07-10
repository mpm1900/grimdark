import { GothicBigButton } from './gothic-ui/button'
import { Marker, MarkerContent } from './ui/marker'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { getHealthRatio, type Actor } from '#/lib/game/actor'
import type { Action } from '#/lib/game/action'
import { useQueryClient } from '@tanstack/react-query'
import { getTargetsFromContext, NULL_CONTEXT } from '#/lib/game/context'
import { useContext } from '#/hooks/use-context'
import { validateContextQuery } from '#/lib/queries/validate-context'
import { setHoverPosition } from '#/lib/stores/ui'
import { lobbyStore } from '#/lib/stores/clients'
import { Gauge } from './gothic-ui/progress'
import { cn } from '#/lib/utils'

function TargetButton({
  is_done,
  is_selected,
  is_valid_target,
  target,
  disabled,
  onClick,
}: Omit<React.ComponentProps<typeof GothicBigButton>, 'onClick'> & {
  is_done: boolean
  is_selected: boolean
  is_valid_target: boolean
  target: Actor | undefined
  onClick: (actor: Actor, is_selected: boolean) => void
}) {
  const is_disabled =
    !is_valid_target || disabled || (!is_selected ? is_done : false)
  if (!target) return <div />
  return (
    <GothicBigButton
      key={target.ID}
      variant={is_selected ? 'red' : 'basic'}
      disabled={is_disabled}
      className={cn('w-full min-w-0 flex-col gap-0 items-center p-0 text-xs')}
      onMouseEnter={() => setHoverPosition(target.position_ID)}
      onMouseLeave={() => setHoverPosition(null)}
      onClick={() => onClick(target, !is_selected)}
    >
      <div className="w-full min-w-0 truncate text-left">{target.name}</div>
      <Gauge
        value={getHealthRatio(target)}
        className="ring ring-black"
        indicator={{ className: 'bg-red-700/20' }}
      />
    </GothicBigButton>
  )
}

function TargetsButtonGrid({
  actor,
  action,
  context,
  disabled,
  ...props
}: React.ComponentProps<'div'> & {
  actor: Actor | null
  action: Action
  context: ReturnType<typeof useContext>
  disabled?: boolean
}) {
  const client = useSelector(lobbyStore, (s) => s.client)
  const players = useSelector(gameStore, (g) => g.players)
  const actors = useSelector(gameStore, (g) => g.actors)
  const positions = useSelector(gameStore, (g) => g.positions)
  const targets_context = context.targets_context ?? NULL_CONTEXT
  const targets = getTargetsFromContext(actors, targets_context)
  const selected = getTargetsFromContext(actors, context.value)
  const query_client = useQueryClient()
  const validate_options = validateContextQuery(context.value)
  const cached_validate = query_client.getQueryData<boolean>(
    validate_options.queryKey
  )
  const validate_result = cached_validate

  if (!client) return null
  return (
    <div {...props}>
      {targets.length === 0 && (
        <Marker variant="separator" className="px-6">
          <MarkerContent>
            {validate_result
              ? "This action doesn't need targets"
              : 'No targets available'}
          </MarkerContent>
        </Marker>
      )}
      {targets.length > 0 && (
        <Marker variant="separator" className="px-16">
          <MarkerContent>
            Select Targets ({action.config.target_count})
          </MarkerContent>
        </Marker>
      )}
      {targets_context.position_IDs.length > 0 ? (
        <div className="grid grid-cols-3 mt-3">
          {players
            .flatMap((p) => {
              const new_pos = positions.filter((pos) => pos.player_ID === p.ID)
              if (p.ID === client?.ID) {
                new_pos.reverse()
              }
              return new_pos
            })
            .map((pos, i) => {
              const target = actors.find((t) => t.position_ID === pos.ID)
              return (
                <TargetButton
                  key={i}
                  is_done={selected.length === action.config.target_count}
                  is_selected={!!target && context.hasTarget(target)}
                  is_valid_target={!!targets.find((t) => t.ID === target?.ID)}
                  target={target}
                  onClick={(t, selected) =>
                    selected ? context.addTarget(t) : context.removeTarget(t)
                  }
                />
              )
            })}
        </div>
      ) : (
        <div className="grid grid-cols-3 mt-3">
          {targets.map((target) => (
            <TargetButton
              key={target.ID}
              is_done={selected.length === action.config.target_count}
              is_selected={!!target && context.hasTarget(target)}
              is_valid_target={!!targets.find((t) => t.ID === target?.ID)}
              target={target}
              onClick={(t, selected) =>
                selected ? context.addTarget(t) : context.removeTarget(t)
              }
            />
          ))}
        </div>
      )}
    </div>
  )
}

export { TargetsButtonGrid }

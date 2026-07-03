import { Loader } from 'lucide-react'
import { GothicBigButton } from './gothic-ui/button'
import { Marker, MarkerContent } from './ui/marker'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import type { Actor } from '#/lib/game/actor'
import type { Action } from '#/lib/game/action'
import { useQuery } from '@tanstack/react-query'
import { getTargetsFromContext, NULL_CONTEXT } from '#/lib/game/context'
import { useContext } from '#/hooks/use-context'
import { validateContextQuery } from '#/lib/queries/validate-context'
import { setHoverPosition } from '#/lib/stores/ui'

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
  if (!target) return <div />
  return (
    <GothicBigButton
      key={target.ID}
      variant={is_selected ? 'red' : 'basic'}
      disabled={
        !is_valid_target || disabled || (!is_selected ? is_done : false)
      }
      onMouseEnter={() => setHoverPosition(target.position_ID)}
      onMouseLeave={() => setHoverPosition(null)}
      onClick={() => onClick(target, !is_selected)}
    >
      {target.name}
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
  const players = useSelector(gameStore, (g) => g.players)
  const actors = useSelector(gameStore, (g) => g.actors)
  const positions = useSelector(gameStore, g => g.positions)
  const turn = useSelector(gameStore, (g) => g.turn)
  const targets_options = getTargetsQuery(
    actor?.ID,
    actor?.player_ID,
    action.ID,
    [turn]
  )
  targets_options.enabled = !disabled
  const targets_query = useQuery(targets_options)
  const targets_context = targets_query.data ?? NULL_CONTEXT
  const targets = getTargetsFromContext(actors, targets_context)
  const selected = getTargetsFromContext(actors, context.value)
  const validate_options = validateContextQuery(context.value)
  validate_options.enabled = !disabled
  const validate_query = useQuery(validate_options)
  const is_loading = targets_query.isFetching || validate_query.isFetching

  return (
    <div {...props}>
      {targets.length === 0 && !is_loading && (
        <Marker variant="separator" className="px-6">
          <MarkerContent>
            {validate_query.data
              ? "This action doesn't have targets."
              : 'No targets available.'}
          </MarkerContent>
        </Marker>
      )}
      {targets.length > 0 && (
        <Marker variant="separator" className="px-16">
          <MarkerContent>Select Targets</MarkerContent>
        </Marker>
      )}
      {targets_context.position_IDs.length > 0 ? (
        <div className="grid grid-cols-3 mt-3">
          {players
            .flatMap((p) => positions.filter(pos => pos.player_ID === p.ID))
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
      {targets_query.isFetching && (
        <div className="grid place-items-center inset-0">
          <Loader className="animate-spin" />
        </div>
      )}
    </div>
  )
}

export { TargetsButtonGrid }

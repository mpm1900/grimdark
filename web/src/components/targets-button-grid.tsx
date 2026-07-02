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
            .flatMap((p) => p.positions.slice().reverse())
            .map((pos, i) => {
              const target = actors.find((t) => t.position_ID === pos.ID)
              const is_target = !!targets.find((t) => t.ID === target?.ID)
              if (!target) return <div key={i} />
              return (
                <GothicBigButton
                  key={target.ID}
                  variant={context.hasTarget(target) ? 'red' : 'basic'}
                  disabled={
                    !is_target ||
                    (disabled || !context.hasTarget(target)
                      ? selected.length === action.config.target_count
                      : false)
                  }
                  onClick={() =>
                    !context.hasTarget(target)
                      ? context.addTarget(target)
                      : context.removeTarget(target)
                  }
                  onMouseEnter={() => setHoverPosition(target.position_ID)}
                  onMouseLeave={() => setHoverPosition(null)}
                >
                  {target.name}
                </GothicBigButton>
              )
            })}
        </div>
      ) : (
        <div className="grid grid-cols-3 mt-3">
          {targets.map((target) => (
            <GothicBigButton
              key={target.ID}
              variant={context.hasTarget(target) ? 'red' : 'basic'}
              disabled={
                disabled || !context.hasTarget(target)
                  ? selected.length === action.config.target_count
                  : false
              }
              onClick={() =>
                !context.hasTarget(target)
                  ? context.addTarget(target)
                  : context.removeTarget(target)
              }
              onMouseEnter={() => setHoverPosition(target.position_ID)}
              onMouseLeave={() => setHoverPosition(null)}
            >
              {target.name}
            </GothicBigButton>
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

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
import { clientsStore } from '#/lib/stores/clients'
import { cva, type VariantProps } from 'class-variance-authority'

const targetButtonVariants = cva('', {
  variants: {
    variant: {
      disabled: 'text-foreground/40',
      source:
        'text-foreground/60 group-hover:text-foreground/70 group-active:text-foreground/60',
      ally: 'text-ally/70 group-hover:text-ally/80 group-active:text-ally/60',
      enemy:
        'text-enemy/70 group-hover:text-enemy/80 group-active:text-enemy/60',
      default: '',
    },
  },
  defaultVariants: {
    variant: 'default',
  },
})

function getVariant(
  is_selected: boolean,
  is_player: boolean,
  is_source: boolean,
  is_disabled: boolean
): VariantProps<typeof targetButtonVariants>['variant'] {
  if (is_selected) return 'default'
  if (is_source) return 'source'
  if (is_disabled) return 'disabled'

  if (is_player) return 'ally'
  if (!is_player) return 'enemy'
  return 'default'
}

function TargetButton({
  client_ID,
  is_done,
  is_selected,
  is_valid_target,
  is_source,
  rank,
  target,
  disabled,
  onClick,
}: Omit<React.ComponentProps<typeof GothicBigButton>, 'onClick'> & {
  client_ID: string
  is_done: boolean
  is_selected: boolean
  is_valid_target: boolean
  is_source: boolean
  rank: number | null
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
      className="flex-col gap-0 p-0 group"
      disabled={is_disabled}
      onMouseEnter={() => setHoverPosition(target.position_ID)}
      onMouseLeave={() => setHoverPosition(null)}
      onClick={() => onClick(target, !is_selected)}
    >
      <div
        className={targetButtonVariants({
          variant: getVariant(
            is_selected,
            client_ID === target.player_ID,
            is_source,
            !!is_disabled
          ),
        })}
      >
        {rank !== null && <span>{rank + 1}. </span>}
        {target.name}
      </div>
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
  const client = useSelector(clientsStore, (s) => s.me)
  const players = useSelector(gameStore, (g) => g.players)
  const actors = useSelector(gameStore, (g) => g.actors)
  const positions = useSelector(gameStore, (g) => g.positions)
  const turn = useSelector(gameStore, (g) => g.turn)
  const targets_options = getTargetsQuery(
    context.value.source_ID ?? actor?.ID,
    context.value.player_ID ?? actor?.player_ID,
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

  if (!client) return null
  return (
    <div {...props}>
      {targets.length === 0 && !is_loading && (
        <Marker variant="separator" className="px-6">
          <MarkerContent>
            {validate_query.data
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
                  client_ID={client.ID}
                  is_done={selected.length === action.config.target_count}
                  is_selected={!!target && context.hasTarget(target)}
                  is_valid_target={!!targets.find((t) => t.ID === target?.ID)}
                  is_source={target?.ID === targets_context.source_ID}
                  rank={pos.rank}
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
              client_ID={client.ID}
              is_done={selected.length === action.config.target_count}
              is_selected={!!target && context.hasTarget(target)}
              is_valid_target={!!targets.find((t) => t.ID === target?.ID)}
              is_source={target?.ID === targets_context.source_ID}
              rank={null}
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

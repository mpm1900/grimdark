import type { Action } from '#/lib/game/action'
import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { MdKeyboardDoubleArrowRight } from 'react-icons/md'
import { AffinityIcon, affinityVariants } from './affinity-name'
import { Button } from './ui/button'
import {
  ItemActions,
  ItemContent,
  ItemDescription,
  ItemMedia,
  ItemTitle,
} from './ui/item'
import { StatIcon } from './stat-name'
import { GothicFramedButton } from './gothic-ui/button'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { useQuery } from '@tanstack/react-query'
import { NULL_CONTEXT } from '#/lib/game/context'
import { setRangePositions } from '#/lib/stores/ui'
import {
  TbArrowBigLeft,
  TbArrowBigRight,
  TbArrowBigRightLines,
  TbSwitchHorizontal,
  TbSwitchVertical,
} from 'react-icons/tb'
import { DNumber } from './dnumber'
import { sendContextMessage } from '#/lib/stores/socket'
import { lobbyStore } from '#/lib/stores/clients'
import { v4 } from 'uuid'

function ActionButton({
  action,
  actor,
  disabled,
  onClick,
  ...props
}: React.ComponentProps<typeof Button> & { action: Action; actor: Actor }) {
  const client = useSelector(lobbyStore, (s) => s.client)
  const turn = useSelector(gameStore, (g) => g.turn)
  const status = useSelector(gameStore, (g) => g.status)
  const actor_command = useSelector(gameStore, (g) =>
    g.commands.find(
      (c) =>
        c.context.source_ID === actor?.ID && c.context.action_ID === action.ID
    )
  )
  const targets_options = getTargetsQuery(
    actor?.ID,
    actor?.player_ID,
    action.ID,
    [turn]
  )
  targets_options.enabled = !disabled
  const targets_query = useQuery(targets_options)
  const targets_context = targets_query.data ?? NULL_CONTEXT
  const is_disabled =
    disabled || action.is_disabled || !actor.is_active || status === 'running'

  return (
    <GothicFramedButton
      {...props}
      onClick={is_disabled ? undefined : onClick}
      variant="basic"
      className={cn(
        'relative h-auto w-full min-w-0 justify-start overflow-hidden border-6',
        is_disabled && 'disabled'
      )}
      aria-disabled={is_disabled}
      onMouseEnter={() => {
        setRangePositions(targets_context.position_IDs)
      }}
      onMouseLeave={() => {
        setRangePositions([])
      }}
    >
      {action.config.stat && (
        <StatIcon
          stat={action.config.stat}
          className="size-15 absolute opacity-20 -right-6 bottom-0 text-white"
        />
      )}
      {action.config.affinity && (
        <ItemMedia className="mr-2">
          <div className="absolute top-0 overflow-hidden size-12 left-0.5">
            <AffinityIcon
              affinity={action.config.affinity}
              className="size-12 absolute top-0 -left-4 opacity-50"
            />
          </div>
        </ItemMedia>
      )}
      <ItemContent className="gap-0 py-0.5 min-w-0 overflow-hidden">
        <ItemTitle
          className={cn(
            'text-white gap-1 font-serif',
            action.is_disabled && 'text-white/60'
          )}
        >
          {action.config.name}
        </ItemTitle>
        <ItemDescription className="block min-w-0 max-w-full overflow-hidden text-ellipsis whitespace-nowrap text-left font-serif text-foreground/70">
          {!!actor_command && status === 'idle' && (
            <span
              className="hover:underline isolate p-0 min-h-9 cursor-pointer"
              onClick={() => {
                sendContextMessage({
                  request_ID: v4(),
                  type: 'cancel-action',
                  client_ID: client?.ID!,
                  context: actor_command.context,
                })
              }}
            >
              [Cancel]
            </span>
          )}{' '}
          {!action.config.power && (
            <span className="block truncate">
              {action.is_disabled && (
                <span className="text-red-300/40 font-serif">
                  [
                  {action.cooldown > 0
                    ? `Cooldown:${action.cooldown}`
                    : 'Disabled'}
                  ]{' '}
                </span>
              )}
              {action.config.description}
            </span>
          )}
          {!!action.config.power && (
            <span className="font-cinzel text-white/50">
              {action.is_disabled && (
                <span className="text-red-300/40 font-serif">
                  [
                  {action.cooldown > 0
                    ? `Cooldown ${action.cooldown}`
                    : 'Disabled'}
                  ]{' '}
                </span>
              )}
              <span className="mr-2 font-semibold">
                <span className="font-serif font-normal text-foreground/70">
                  {action.config.description}
                </span>
                {action.config.crit_chance && (
                  <span className="inline-flex items-baseline align-baseline">
                    <DNumber value={actor.stats['critical-chance']} r={100}>
                      {(action.config.crit_chance * 100).toFixed(0)}%
                    </DNumber>
                    <MdKeyboardDoubleArrowRight className="self-center" />
                    <DNumber value={actor.stats['critical-damage']} r={100}>
                      x{action.config.crit_modifier.toFixed(2)}
                    </DNumber>
                  </span>
                )}
              </span>
            </span>
          )}
        </ItemDescription>
      </ItemContent>
      {!!action.config.power && (
        <ItemActions className="relative flex flex-col gap-0 h-full items-end -mt-1 font-cinzel">
          <span
            className={cn(
              'text-xl leading-6 font-black font-cinzel-dec',
              affinityVariants({
                affinity: action.config.affinity,
              })
            )}
          >
            {action.config.power}
          </span>
          {action.config.accuracy && (
            <DNumber
              value={actor.stats['accuracy']}
              r={100}
              className={cn('text-xs font-semibold text-foreground/80')}
            >
              {Math.min(action.config.accuracy * 100, 100).toFixed(0)}%
            </DNumber>
          )}
        </ItemActions>
      )}
    </GothicFramedButton>
  )
}

function SystemActionButton({
  action,
  actor,
  disabled,
  ...props
}: React.ComponentProps<typeof GothicFramedButton> & {
  action: Action
  actor: Actor
}) {
  const status = useSelector(gameStore, (g) => g.status)
  return (
    <GothicFramedButton
      disabled={
        disabled ||
        action.is_disabled ||
        !actor.is_active ||
        status === 'running'
      }
      className="flex-1 w-12 p-0 text-stone-400"
      {...props}
    >
      {action.tags.includes('swap') && (
        <TbSwitchHorizontal className="size-6" />
      )}
      {action.tags.includes('retreat') && (
        <TbSwitchVertical className="size-6" />
      )}
      {action.tags.includes('back') && <TbArrowBigLeft className="size-6" />}
      {action.tags.includes('forward') && (
        <TbArrowBigRight className="size-6" />
      )}
      {action.tags.includes('front') && (
        <TbArrowBigRightLines className="size-6" />
      )}
    </GothicFramedButton>
  )
}

export { ActionButton, SystemActionButton }

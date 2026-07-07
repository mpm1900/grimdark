import type { Actor } from '#/lib/game/actor'
import type { PropsWithChildren, ReactNode } from 'react'
import { Dialog, DialogClose, DialogFooter } from './ui/dialog'
import type { Action } from '#/lib/game/action'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { useQuery } from '@tanstack/react-query'
import { NULL_CONTEXT } from '#/lib/game/context'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { useContext } from '#/hooks/use-context'
import { validateContextQuery } from '#/lib/queries/validate-context'
import { sendContextMessage } from '#/lib/stores/socket'
import { clientsStore } from '#/lib/stores/clients'
import { GothicFramedButton } from './gothic-ui/button'
import {
  GothicDialogContent,
  GothicDialogHeader,
  GothicDialogTitle,
} from './gothic-ui/dialog'
import { TargetsButtonGrid } from './targets-button-grid'
import { nextActiveActor } from '#/lib/stores/ui'
import { GothicBadge } from './gothic-ui/badge'
import { AffinityIcon } from './affinity-name'
import { StatIcon } from './stat-name'
import { cn } from '#/lib/utils'

function formatNumber(value: number) {
  return Number.isInteger(value) ? value.toFixed(0) : value.toFixed(1)
}

function formatPercent(value: number) {
  return `${Math.min(value * 100, 100).toFixed(0)}%`
}

function formatSigned(value: number) {
  return value > 0 ? `+${value}` : value.toString()
}

function pluralize(count: number, singular: string, plural = `${singular}s`) {
  return `${count} ${count === 1 ? singular : plural}`
}

function ActionDetailCard({
  label,
  value,
  className,
  valueClassName,
}: {
  label: string
  value: ReactNode
  className?: string
  valueClassName?: string
}) {
  return (
    <div
      className={cn(
        'relative isolate min-h-20 px-3 py-2.5 text-center',
        'before:pointer-events-none before:absolute before:left-1/2 before:top-0 before:size-1 before:-translate-x-1/2 before:rotate-45 before:bg-stone-500/40',
        className
      )}
    >
      <div className="relative font-cinzel text-xs font-semibold text-foreground/60">
        {label}
      </div>
      <div
        className={cn(
          'relative mt-1 font-cinzel-dec text-3xl leading-none font-bold text-stone-100',
          '[text-shadow:0_2px_0_var(--color-black)]',
          valueClassName
        )}
      >
        {value ?? '-'}
      </div>
    </div>
  )
}

function ActionDetails({ action }: { action: Action }) {
  const {
    accuracy,
    affinity,
    cooldown,
    crit_chance,
    crit_modifier,
    crit_stage,
    hits,
    lifesteal,
    power,
    priority,
    range,
    recoil,
    stat,
  } = action.config

  const has_power = power > 0
  const has_primary_details = has_power || accuracy !== null || range !== null
  const has_secondary_details =
    has_power &&
    (hits > 0 ||
      cooldown > 0 ||
      priority !== 0 ||
      crit_chance > 0 ||
      crit_stage !== 0 ||
      lifesteal > 0 ||
      recoil > 0)

  return (
    <div className="relative p-3">
      {affinity && (
        <AffinityIcon
          affinity={affinity}
          className="pointer-events-none absolute -left-7 -top-7 size-28 opacity-30"
        />
      )}
      {stat && (
        <StatIcon
          stat={stat}
          className="pointer-events-none absolute -right-7 -top-7 size-28 opacity-20"
        />
      )}

      {has_primary_details && (
        <div className="relative mx-auto grid grid-cols-3 max-w-80 gap-3 px-6 pt-3 sm:max-w-96 sm:px-10 [&>div]:flex-1">
          {(accuracy !== null || has_power) && (
            <ActionDetailCard
              label="Accuracy"
              value={accuracy === null ? 'Sure' : formatPercent(accuracy)}
              className="opacity-90"
              valueClassName={'text-white/60'}
            />
          )}
          {has_power && (
            <ActionDetailCard
              label="Power"
              value={formatNumber(power)}
              className="z-10 min-h-24 -translate-y-1 scale-110"
              valueClassName="text-4xl"
            />
          )}

          <ActionDetailCard
            label="Range"
            value={range}
            className="opacity-90"
            valueClassName={'text-white/60'}
          />
        </div>
      )}

      <div className="text-center text-white/50 p-4 italic">
        {action.config.description}
        {cooldown > 0 && (
          <span className="text-white/80">{` ${pluralize(cooldown, 'Turn')} cooldown.`}</span>
        )}
      </div>

      {has_secondary_details && (
        <div className="flex flex-wrap justify-center gap-1.5">
          {hits > 1 && (
            <GothicBadge variant="empty">{pluralize(hits, 'Hit')}</GothicBadge>
          )}
          {crit_stage > 0 && (
            <GothicBadge variant="empty">
              Crit {formatPercent(crit_chance)} / x{formatNumber(crit_modifier)}
            </GothicBadge>
          )}
          {lifesteal > 0 && (
            <GothicBadge variant="empty">
              Lifesteal {formatPercent(lifesteal)}
            </GothicBadge>
          )}
          {recoil > 0 && (
            <GothicBadge variant="empty">
              Recoil {formatPercent(recoil)}
            </GothicBadge>
          )}
          {priority !== 0 && (
            <GothicBadge variant="empty">
              Priority {formatSigned(priority)}
            </GothicBadge>
          )}
        </div>
      )}
    </div>
  )
}

function ActionContextDialog({
  actor,
  action,
  children,
  enabled,
}: PropsWithChildren<{
  actor: Actor
  action: Action
  enabled?: boolean
}>) {
  const client = useSelector(clientsStore, (s) => s.me!)
  const turn = useSelector(gameStore, (g) => g.turn)
  const targets_options = getTargetsQuery(
    actor.ID,
    actor.player_ID,
    action.ID,
    [turn]
  )
  targets_options.enabled = !!enabled
  const targets_query = useQuery(targets_options)
  const targets_context = targets_query.data ?? NULL_CONTEXT
  const context = useContext(targets_context)
  const validate_options = validateContextQuery(context.value)
  validate_options.enabled = !!enabled
  const validate_query = useQuery(validate_options)
  const is_loading = targets_query.isFetching || validate_query.isFetching

  return (
    <Dialog>
      {children}
      <GothicDialogContent className="pt-0">
        <div className="pointer-events-none absolute inset-x-0 top-0 -z-10 h-28 overflow-hidden">
          <div className="absolute left-1/2 h-full w-[calc(100%+5rem)] -translate-x-1/2 bg-[url('/gothic/DialogFlag.png')] bg-[length:100%_100%] bg-center bg-no-repeat opacity-80" />
        </div>

        <GothicDialogHeader>
          <GothicDialogTitle>{action.config.name}</GothicDialogTitle>
        </GothicDialogHeader>

        <div className="overflow-hidden min-h-32">
          <ActionDetails action={action} />

          <TargetsButtonGrid
            actor={actor}
            action={action}
            context={context}
            className="px-4"
          />
        </div>
        <DialogFooter className="p-0 -mr-1 -mb-0.5">
          <DialogClose asChild>
            <GothicFramedButton
              variant="red"
              disabled={!validate_query.data || is_loading}
              onClick={() => {
                sendContextMessage({
                  type: 'push-action',
                  client_ID: client.ID,
                  context: context.value,
                })

                nextActiveActor(actor)
                context.reset()
              }}
            >
              Confirm
            </GothicFramedButton>
          </DialogClose>
        </DialogFooter>
      </GothicDialogContent>
    </Dialog>
  )
}

export { ActionContextDialog }

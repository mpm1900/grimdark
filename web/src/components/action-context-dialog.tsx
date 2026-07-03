import type { Actor } from '#/lib/game/actor'
import type { PropsWithChildren } from 'react'
import { Dialog, DialogClose, DialogFooter } from './ui/dialog'
import type { Action } from '#/lib/game/action'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { useQuery } from '@tanstack/react-query'
import { Marker, MarkerContent } from './ui/marker'
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
import { STAT_LABELS } from '#/lib/game/core'
import { TargetsButtonGrid } from './targets-button-grid'

function AttackDetails({ action }: { action: Action }) {
  return (
    <div className="flex [&>div]:flex-1 px-3 pb-3 bg-right">
      {!!action.config.stat && (
        <div className="flex flex-col items-center">
          <Marker variant="separator">
            <MarkerContent>Stat</MarkerContent>
          </Marker>
          {STAT_LABELS[action.config.stat]}
        </div>
      )}
      {!!action.config.power && (
        <div className="flex flex-col items-center">
          <Marker variant="separator">
            <MarkerContent>Power</MarkerContent>
          </Marker>
          {action.config.power}
        </div>
      )}
      {action.config.accuracy && (
        <div className="flex flex-col items-center">
          <Marker variant="separator">
            <MarkerContent>Accuracy</MarkerContent>
          </Marker>
          {Math.min(action.config.accuracy * 100, 100)}%
        </div>
      )}
      {action.config.range !== null && (
        <div className="flex flex-col items-center">
          <Marker variant="separator">
            <MarkerContent>Range</MarkerContent>
          </Marker>
          {action.config.range}
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
      <GothicDialogContent>
        <div className="pointer-events-none absolute inset-x-0 top-0 -z-10 h-24 overflow-hidden">
          <div className="absolute left-1/2 h-full w-[calc(100%+5rem)] -translate-x-1/2 bg-[url('/gothic/DialogFlag.png')] bg-[length:100%_100%] bg-center bg-no-repeat opacity-70" />
        </div>

        <GothicDialogHeader>
          <GothicDialogTitle>{action.config.name}</GothicDialogTitle>
        </GothicDialogHeader>

        {!!action.config.power && <AttackDetails action={action} />}
        {action.config.description && (
          <div className="text-center text-white/60">
            {action.config.description}
          </div>
        )}
        <TargetsButtonGrid
          actor={actor}
          action={action}
          context={context}
          className="px-4"
        />
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

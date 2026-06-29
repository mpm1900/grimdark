import type { Actor } from '#/lib/game/actor'
import type { PropsWithChildren } from 'react'
import { Dialog, DialogClose, DialogFooter } from './ui/dialog'
import type { Action } from '#/lib/game/action'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { useQuery } from '@tanstack/react-query'
import { Marker, MarkerContent } from './ui/marker'
import { Field, FieldContent } from './ui/field'
import { Loader } from 'lucide-react'
import { getTargetsFromContext, NULL_CONTEXT } from '#/lib/game/context'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { useContext } from '#/hooks/use-context'
import { validateContextQuery } from '#/lib/queries/validate-context'
import { sendContextMessage } from '#/lib/stores/socket'
import { clientsStore } from '#/lib/stores/clients'
import { GothicBigButton, GothicFramedButton } from './gothic-ui/button'
import {
  GothicDialogContent,
  GothicDialogHeader,
  GothicDialogTitle,
} from './gothic-ui/dialog'

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
  const actors = useSelector(gameStore, (g) => g.actors)
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
  const targets = getTargetsFromContext(actors, targets_context)
  const context = useContext(targets_context)
  const validate_options = validateContextQuery(context.value)
  validate_options.enabled = !!enabled
  const validate_query = useQuery(validate_options)
  const is_loading = targets_query.isFetching || validate_query.isFetching

  return (
    <Dialog>
      {children}
      <GothicDialogContent>
        <GothicDialogHeader>
          <GothicDialogTitle>{action.config.name}</GothicDialogTitle>
        </GothicDialogHeader>
        {action.config.description && (
          <div className="text-center text-muted-foreground">
            {action.config.description}
          </div>
        )}
        <div className="grid grid-cols-3 px-3">
          {action.config.accuracy && (
            <Marker variant="separator">
              <MarkerContent>
                Accuracy: {action.config.accuracy * 100}%
              </MarkerContent>
            </Marker>
          )}
          {!!action.config.power && (
            <Marker variant="separator">
              <MarkerContent>Power: {action.config.power}</MarkerContent>
            </Marker>
          )}
          {!!action.config.recoil && (
            <Marker variant="separator">
              <MarkerContent>Recoil: {action.config.recoil}</MarkerContent>
            </Marker>
          )}
          {!action.config.recoil && !!action.config.lifesteal && (
            <Marker variant="separator">
              <MarkerContent>
                Lifesteal: {action.config.lifesteal * 100}%
              </MarkerContent>
            </Marker>
          )}
        </div>
        <div className="px-4">
          <div className="grid grid-cols-2">
            {targets.map((target) => (
              <GothicBigButton
                key={target.ID}
                variant={context.hasTarget(target) ? 'red' : 'basic'}
                // pressed={context.hasTarget(target)}
                onClick={(e) =>
                  !context.hasTarget(target)
                    ? context.addTarget(target)
                    : context.removeTarget(target)
                }
                disabled={!enabled}
              >
                {target.name}
              </GothicBigButton>
            ))}
          </div>
          {targets_query.isFetching && (
            <div className="grid place-items-center absolute inset-0">
              <Loader className="animate-spin" />
            </div>
          )}
          {targets.length === 0 && !is_loading && (
            <Marker variant="separator" className="px-6">
              <MarkerContent>
                {validate_query.data
                  ? "This action doesn't have targets."
                  : 'No targets available.'}
              </MarkerContent>
            </Marker>
          )}
        </div>
        <DialogFooter className="p-2">
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
              Choose
            </GothicFramedButton>
          </DialogClose>
        </DialogFooter>
      </GothicDialogContent>
    </Dialog>
  )
}

export { ActionContextDialog }

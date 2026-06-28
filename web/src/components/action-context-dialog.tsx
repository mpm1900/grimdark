import type { Actor } from '#/lib/game/actor'
import type { PropsWithChildren } from 'react'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from './ui/dialog'
import type { Action } from '#/lib/game/action'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { useQuery } from '@tanstack/react-query'
import { Marker, MarkerContent } from './ui/marker'
import { Field, FieldContent, FieldLabel } from './ui/field'
import { Button } from './ui/button'
import { ChevronRight } from 'lucide-react'
import { getTargetsFromContext, NULL_CONTEXT } from '#/lib/game/context'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { useContext } from '#/hooks/use-context'
import { Toggle } from './ui/toggle'
import { validateContextQuery } from '#/lib/queries/validate-context'

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
  const targets_options = getTargetsQuery(actor, action.ID)
  targets_options.enabled = !!enabled
  const targets_query = useQuery(targets_options)
  const targets_context = targets_query.data ?? NULL_CONTEXT
  const actors = useSelector(gameStore, (g) => g.actors)
  const targets = getTargetsFromContext(actors, targets_context)
  const context = useContext(targets_context)
  const validate_options = validateContextQuery(context.value)
  validate_options.enabled = !!enabled
  const validate_query = useQuery(validate_options)

  return (
    <Dialog>
      {children}
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{action.config.name}</DialogTitle>
          {action.config.description && (
            <DialogDescription>{action.config.description}</DialogDescription>
          )}
        </DialogHeader>
        <div className="grid grid-cols-3">
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
        <div>
          <Field>
            <FieldContent>
              <div className="gap-4 grid grid-cols-2">
                {targets.map((target) => (
                  <Toggle
                    key={target.ID}
                    variant="outline"
                    pressed={context.hasTarget(target)}
                    onPressedChange={(pressed) =>
                      pressed
                        ? context.addTarget(target)
                        : context.removeTarget(target)
                    }
                    disabled={!enabled}
                  >
                    {target.name}
                  </Toggle>
                ))}
              </div>
              {targets.length === 0 && (
                <Marker variant="separator">
                  <MarkerContent>
                    {validate_query.data
                      ? "This action doesn't have targets."
                      : 'No targets available.'}
                  </MarkerContent>
                </Marker>
              )}
            </FieldContent>
          </Field>
        </div>
        <DialogFooter>
          <DialogClose asChild>
            <Button
              disabled={!validate_query.data || validate_query.isFetching}
              onClick={() => {
                context.reset()
              }}
            >
              Choose <ChevronRight />
            </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

export { ActionContextDialog }

import { gameStore } from '#/lib/stores/game'
import { useSelector } from '@tanstack/react-store'
import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from './ui/alert-dialog'
import { clientsStore } from '#/lib/stores/clients'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { useQuery } from '@tanstack/react-query'
import { getTargetsFromContext, NULL_CONTEXT } from '#/lib/game/context'
import { useContext } from '#/hooks/use-context'
import { validateContextQuery } from '#/lib/queries/validate-context'
import { Toggle } from './ui/toggle'
import { ChevronRight, Loader } from 'lucide-react'
import { Marker, MarkerContent } from './ui/marker'
import { sendContextMessage } from '#/lib/stores/socket'
import { Button } from './ui/button'

function PromptController() {
  const prompt = useSelector(gameStore, (g) => g.prompts[0])
  const client = useSelector(clientsStore, (s) => s.me!)
  const actors = useSelector(gameStore, (g) => g.actors)
  const turn = useSelector(gameStore, (g) => g.turn)
  const targets_options = getTargetsQuery(null, prompt?.context.player_ID, prompt?.payload.ID, [
    prompt?.ID ?? '',
    turn,
  ])
  targets_options.enabled = !!prompt
  const targets_query = useQuery(targets_options)
  const targets_context = targets_query.data ?? NULL_CONTEXT
  const targets = getTargetsFromContext(actors, targets_context)
  const context = useContext(targets_context)
  const resolved_context = {
    ...context.value,
    action_ID: prompt?.payload.ID ?? null,
    player_ID: prompt?.context.player_ID ?? null,
    position_IDs: prompt?.context.position_IDs ?? [],
  }
  const validate_options = validateContextQuery(resolved_context)
  validate_options.enabled = !!prompt
  const validate_query = useQuery(validate_options)
  const is_loading = targets_query.isFetching || validate_query.isFetching

  return (
    <>
      <AlertDialog open={!!prompt}>
        <AlertDialogContent>
          {prompt && (
            <>
              <AlertDialogHeader>
                <AlertDialogTitle>
                  {prompt.payload.config.name}
                </AlertDialogTitle>
                <AlertDialogDescription>
                  {prompt.payload.config.description}
                </AlertDialogDescription>
              </AlertDialogHeader>
              <div>
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
                    >
                      {target.name}
                    </Toggle>
                  ))}
                </div>
                {targets_query.isFetching && (
                  <div className="grid place-items-center absolute inset-0">
                    <Loader className="animate-spin" />
                  </div>
                )}
                {targets.length === 0 && !is_loading && (
                  <Marker variant="separator">
                    <MarkerContent>
                      {validate_query.data
                        ? "This action doesn't have targets."
                        : 'No targets available.'}
                    </MarkerContent>
                  </Marker>
                )}
              </div>
            </>
          )}
          <AlertDialogFooter>
            <Button
              disabled={!validate_query.data || is_loading}
              onClick={() => {
                sendContextMessage({
                  type: 'resolve-prompt',
                  client_ID: client.ID,
                  context: resolved_context,
                })

                context.reset()
              }}
            >
              Choose <ChevronRight />
            </Button>
          </AlertDialogFooter>
        </AlertDialogContent>

      </AlertDialog>
    </>
  )
}

export { PromptController }

import { gameStore } from '#/lib/stores/game'
import { useSelector } from '@tanstack/react-store'
import { AlertDialog, AlertDialogFooter } from './ui/alert-dialog'
import { clientsStore } from '#/lib/stores/clients'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { useQuery } from '@tanstack/react-query'
import { NULL_CONTEXT } from '#/lib/game/context'
import { useContext } from '#/hooks/use-context'
import { validateContextQuery } from '#/lib/queries/validate-context'
import { sendContextMessage } from '#/lib/stores/socket'
import {
  GothicAlertDialogContent,
  GothicAlertDialogTitle,
  GothicDialogHeader,
} from './gothic-ui/dialog'
import { GothicFramedButton } from './gothic-ui/button'
import { TargetsButtonGrid } from './targets-button-grid'

function PromptController() {
  const prompt = useSelector(gameStore, (g) => g.prompts[0])
  const client = useSelector(clientsStore, (s) => s.me!)
  const turn = useSelector(gameStore, (g) => g.turn)
  const targets_options = getTargetsQuery(
    null,
    prompt?.context.player_ID,
    prompt?.payload.ID,
    [prompt?.ID ?? '', turn]
  )
  targets_options.enabled = !!prompt
  const targets_query = useQuery(targets_options)
  const targets_context = targets_query.data ?? NULL_CONTEXT
  const context = useContext(targets_context)
  const selected_actor_count = context.value.actor_IDs.filter(Boolean).length
  const resolved_context = {
    ...context.value,
    action_ID: prompt?.payload.ID ?? null,
    player_ID: prompt?.context.player_ID ?? null,
    position_IDs:
      prompt?.context.position_IDs.slice(0, selected_actor_count) ?? [],
  }
  const validate_options = validateContextQuery(resolved_context)
  validate_options.enabled = !!prompt
  const validate_query = useQuery(validate_options)
  const is_loading = targets_query.isFetching || validate_query.isFetching

  return (
    <>
      <AlertDialog open={!!prompt}>
        <GothicAlertDialogContent className="focus:outline-none focus-visible:outline-none">
          <div className="pointer-events-none absolute inset-x-0 top-0 -z-10 h-24 overflow-hidden">
            <div className="absolute left-1/2 h-full w-[calc(100%+5rem)] -translate-x-1/2 bg-[url('/gothic/DialogFlag.png')] bg-[length:100%_100%] bg-center bg-no-repeat opacity-70" />
          </div>
          {prompt && (
            <>
              <GothicDialogHeader>
                <GothicAlertDialogTitle>
                  {prompt.payload.config.name}
                </GothicAlertDialogTitle>
              </GothicDialogHeader>
              {prompt.payload.config.description && (
                <div className="text-center text-white/60">
                  {prompt.payload.config.description}
                </div>
              )}

              <TargetsButtonGrid
                actor={null}
                action={prompt.payload}
                context={{
                  ...context,
                  value: resolved_context,
                }}
              />
            </>
          )}
          <AlertDialogFooter className="p-0 -mr-1 -mb-0.5">
            <GothicFramedButton
              variant="red"
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
              Confirm
            </GothicFramedButton>
          </AlertDialogFooter>
        </GothicAlertDialogContent>
      </AlertDialog>
    </>
  )
}

export { PromptController }

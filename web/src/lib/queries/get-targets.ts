import { queryOptions } from '@tanstack/react-query'
import { contextToString, NULL_CONTEXT, type Context } from '../game/context'
import { clientsStore } from '../stores/clients'
import { sendContextMessage } from '../stores/socket'
import type { ID } from '../game/core'
import { subscribe } from '../socket/subscribe'

async function getTargets(context: Context): Promise<Context> {
  const client = clientsStore.get().me

  const promise: Promise<Context> = new Promise((resolve, reject) => {
    if (!client) {
      return reject('no client')
    }

    sendContextMessage({
      type: 'get-targets',
      client_ID: client.ID,
      context,
    })

    const unsub = subscribe((_event, message) => {
      if (message && message.context) {
        if (
          message.type === 'target-IDs' &&
          message.context.source_ID === context.source_ID &&
          message.context.action_ID === context.action_ID
        ) {
          resolve(message.context as Context)
          unsub()
        }
      } else {
        reject('no message')
        unsub()
      }
    })
  })

  return promise
}

function getTargetsQuery(
  source_ID: ID | null | undefined,
  player_ID: ID | null | undefined,
  action_ID: ID,
  deps: (boolean | number | string)[]
) {
  const context: Context = {
    ...NULL_CONTEXT,
    action_ID,
    source_ID: source_ID ?? null,
    parent_ID: source_ID ?? null,
    player_ID: player_ID ?? null,
  }
  return queryOptions<Context>({
    queryKey: ['get-targets', contextToString(context), ...deps],
    queryFn: async () => {
      return getTargets(context)
    },
  })
}

export { getTargets, getTargetsQuery }

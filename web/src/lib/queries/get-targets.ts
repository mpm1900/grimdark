import { queryOptions } from '@tanstack/react-query'
import { contextToString, NULL_CONTEXT, type Context } from '../game/context'
import { lobbyStore } from '../stores/clients'
import { sendContextMessage } from '../stores/socket'
import type { ID } from '../game/core'
import { subscribe } from '../socket/subscribe'
import { v4 } from 'uuid'

async function getTargets(context: Context): Promise<Context> {
  const client = lobbyStore.state.client

  const promise: Promise<Context> = new Promise((resolve, reject) => {
    if (!client) {
      return reject('no client')
    }

    const request_ID = v4()
    sendContextMessage({
      request_ID,
      type: 'get-targets',
      client_ID: client.ID,
      context,
    })

    const unsub = subscribe((_event, message) => {
      if (message && message.context) {
        if (message.request_ID === request_ID) {
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
  deps: (boolean | number | string)[],
  context_overrides: Partial<Context> = {}
) {
  const context: Context = {
    ...NULL_CONTEXT,
    source_ID: source_ID ?? null,
    parent_ID: source_ID ?? null,
    player_ID: player_ID ?? null,
    ...context_overrides,
    action_ID,
  }
  return queryOptions<Context>({
    queryKey: ['get-targets', contextToString(context), ...deps],
    queryFn: async () => {
      return getTargets(context)
    },
  })
}

export { getTargets, getTargetsQuery }

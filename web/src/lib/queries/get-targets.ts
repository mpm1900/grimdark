import { queryOptions } from '@tanstack/react-query'
import { contextToString, NULL_CONTEXT, type Context } from '../game/context'
import { subscribe } from '../socket/connect'
import { clientsStore } from '../stores/clients'
import { sendContextMessage } from '../stores/socket'
import type { Actor } from '../game/actor'
import type { ID } from '../game/core'

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
      if (message) {
        if (
          message.type === 'target-IDs' &&
          message.context?.source_ID === context.source_ID &&
          message.context?.action_ID === context.action_ID
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

function getTargetsQuery(actor: Actor, action_ID: ID) {
  const context: Context = {
    ...NULL_CONTEXT,
    action_ID,
    source_ID: actor.ID,
    parent_ID: actor.ID,
    player_ID: actor.player_ID,
  }
  return queryOptions<Context>({
    queryKey: ['get-targets', actor.ID, action_ID],
    queryFn: async () => {
      return getTargets(context)
    },
  })
}

export { getTargets, getTargetsQuery }

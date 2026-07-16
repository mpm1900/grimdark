import { queryOptions } from '@tanstack/react-query'
import { contextToString, type Context } from '../game/context'
import { lobbyStore } from '../stores/clients'
import { sendContextMessage } from '../stores/socket'
import { subscribe } from '../socket/subscribe'
import { v4 } from 'uuid'

async function validateContext(context: Context): Promise<boolean> {
  const client = lobbyStore.state.client
  const promise: Promise<boolean> = new Promise((resolve, reject) => {
    if (!client) {
      return reject('no client')
    }

    const request_ID = v4()
    sendContextMessage({
      request_ID,
      type: 'validate-context',
      client_ID: client.ID,
      context,
    })

    const unsub = subscribe((_event, message) => {
      if (message && message.context) {
        if (message.request_ID === request_ID) {
          resolve(!!message.valid)
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

function validateContextQuery(context: Context) {
  return queryOptions<boolean>({
    queryKey: ['validate-context', contextToString(context)],
    queryFn: async () => {
      return validateContext(context)
    },
    staleTime: 0,
    gcTime: 0,
  })
}

export { validateContext, validateContextQuery }

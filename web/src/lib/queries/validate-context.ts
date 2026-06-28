import { queryOptions } from '@tanstack/react-query'
import { contextToString, type Context } from '../game/context'
import { subscribe } from '../socket/connect'
import { clientsStore } from '../stores/clients'
import { sendContextMessage } from '../stores/socket'

async function validateContext(context: Context): Promise<boolean> {
  const client = clientsStore.get().me
  const promise: Promise<boolean> = new Promise((resolve, reject) => {
    if (!client) {
      return reject('no client')
    }

    sendContextMessage({
      type: 'validate-context',
      client_ID: client.ID,
      context,
    })

    const unsub = subscribe((_event, message) => {
      if (message && message.context) {
        if (
          message.type === 'validate-context' &&
          contextToString(context) == contextToString(message.context)
        ) {
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

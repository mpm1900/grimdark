import { queryOptions } from '@tanstack/react-query'
import { contextToString, type Context } from '../game/context'
import { subscribe } from '../socket/connect'
import { clientsStore } from '../stores/clients'
import { sendContextMessage } from '../stores/socket'

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
        if (message.type === 'target-IDs') {
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

function getTargetsQuery(context: Context) {
  return queryOptions({
    queryKey: ['get-targets', contextToString(context)],
    queryFn: async () => {
      return getTargets(context)
    },
  })
}

export { getTargets, getTargetsQuery }

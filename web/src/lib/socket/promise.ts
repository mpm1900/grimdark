import { v4 } from 'uuid'
import { lobbyStore } from '../stores/clients'
import { sendContextMessage } from '../stores/socket'
import type { SocketRequest, SocketResponse } from './request'
import { subscribe } from './subscribe'

export function promisify(
  request: Omit<SocketRequest, 'request_ID' | 'client_ID'>
): Promise<SocketResponse> {
  const client = lobbyStore.state.client
  const promise: Promise<SocketResponse> = new Promise((resolve, reject) => {
    if (!client) {
      return reject('no client')
    }

    const request_ID = v4()
    sendContextMessage({
      ...request,
      client_ID: client.ID,
      request_ID,
    })

    const unsub = subscribe((_event, message) => {
      if (message && message.context) {
        if (message.request_ID === request_ID) {
          resolve(message)
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

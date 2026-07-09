import { useMutation } from '@tanstack/react-query'
import { connect } from '../socket/connect'
import type { TeamConfig } from '../game/team'
import {
  sendContextMessage,
  socketStore,
  type SocketResponse,
} from '../stores/socket'
import { NULL_CONTEXT } from '../game/context'
import { lobbyStore } from '../stores/clients'
import { subscribe } from '../socket/subscribe'

async function postConnect(team_config: TeamConfig): Promise<SocketResponse> {
  const client = lobbyStore.state.client
  return new Promise((resolve, reject) => {
    if (!client) {
      return reject('no client')
    }

    sendContextMessage({
      type: 'post-connect',
      client_ID: client.ID,
      context: NULL_CONTEXT,
      team_config,
    })

    const unsub = subscribe((_event, message) => {
      if (message) {
        if (message.type === 'post-connect') {
          socketStore.setState((s) => ({
            ...s,
            instance_ID: message.game?.instance_ID ?? null,
          }))
          resolve(message)
          unsub()
        }
      } else {
        reject('no message')
        unsub()
      }
    })
  })
}

function useConnect() {
  return useMutation({
    mutationKey: ['connect'],
    mutationFn: async (team_config: TeamConfig) => {
      const connect_message = await connect()
      const message = await postConnect(team_config)
      console.log('message', message)
      return message
    },
  })
}

export { useConnect }

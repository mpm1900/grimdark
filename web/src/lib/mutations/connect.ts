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
import type { ID } from '../game/core'

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
    mutationFn: async (team_config: TeamConfig & { instance_ID?: ID }) => {
      await connect(team_config.instance_ID)
      const message = await postConnect(team_config)
      return message
    },
  })
}

export { useConnect }

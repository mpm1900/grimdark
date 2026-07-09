import { useMutation } from '@tanstack/react-query'
import { connect } from '../socket/connect'
import type { TeamConfig } from '../game/team'
import { sendContextMessage } from '../stores/socket'
import { NULL_CONTEXT } from '../game/context'

function useConnect() {
  return useMutation({
    mutationKey: ['connect'],
    mutationFn: async (team: TeamConfig) => {
      const message = await connect()
      sendContextMessage({
        type: 'load-team',
        client_ID: team.user.id,
        context: NULL_CONTEXT,
        team_config: team,
      })
      return message
    },
  })
}

export { useConnect }

import { api } from '#/integrations/axios/instance'
import { createServerFn } from '@tanstack/react-start'
import { TeamSchema, type Team } from '../game/team'
import {
  mutationOptions,
  useMutation,
  useQueryClient,
} from '@tanstack/react-query'

const saveTeam = createServerFn()
  .validator(TeamSchema)
  .handler(async ({ data }) => {
    const response = await api.post<Team>('/api/teams/save', data)
    return response.data
  })

function useSaveTeam() {
  const qc = useQueryClient()
  return useMutation(
    mutationOptions({
      mutationKey: ['save-team'],
      mutationFn: async (team: Team) => {
        const response = await saveTeam({ data: team })
        await qc.invalidateQueries({ queryKey: ['teams'] })
        return response
      },
    })
  )
}

export { saveTeam, useSaveTeam }

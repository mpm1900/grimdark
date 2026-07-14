import { api } from '#/integrations/axios/instance'
import { queryOptions } from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import type { Team } from '../game/team'

const getTeams = createServerFn().handler(async () => {
  const response = await api.get<Team[]>(`/api/teams`)
  return response.data
})

const teamsQuery = queryOptions({
  queryKey: ['teams'],
  queryFn: () => getTeams(),
})

export { getTeams, teamsQuery }

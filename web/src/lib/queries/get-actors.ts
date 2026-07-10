import { queryOptions } from '@tanstack/react-query'
import { api } from '#/integrations/axios/instance'
import { createServerFn } from '@tanstack/react-start'
import type { ActorClass } from '../game/actor-class'

const getActors = createServerFn().handler(async () => {
  const response = await api.get<ActorClass[]>(`/api/actors`)
  return response.data
})

const actorsQuery = queryOptions({
  queryKey: ['get-actors'],
  queryFn: () => getActors(),
  select: (data) => data.sort((a, b) => a.name.localeCompare(b.name)),
  staleTime: 60000,
  gcTime: 60000,
})

export { actorsQuery, getActors }

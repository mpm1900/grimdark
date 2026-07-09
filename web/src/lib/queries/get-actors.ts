import { queryOptions } from '@tanstack/react-query'
import { getApiBaseUrl } from '#/utils/get-api-base-url'
import { createServerFn } from '@tanstack/react-start'
import type { ActorClass } from '../game/actor-class'

const getActors = createServerFn().handler(async () => {
  const response = await fetch(`${getApiBaseUrl()}/api/actors`)
  const data = await response.json()
  return data as ActorClass[]
})

const actorsQuery = queryOptions({
  queryKey: ['get-actors'],
  queryFn: () => getActors(),
  select: (data) => data.sort((a, b) => a.name.localeCompare(b.name)),
  staleTime: 60000,
  gcTime: 60000,
})

export { actorsQuery, getActors }

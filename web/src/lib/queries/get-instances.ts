import { api } from '#/integrations/axios/instance'
import { queryOptions } from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import type { Instance } from '../game/instance'

const getInstances = createServerFn().handler(async () => {
  const response = await api.get<Instance[]>(`/api/instances`)
  return response.data
})

const instancesQuery = queryOptions({
  queryKey: ['instances'],
  queryFn: () => getInstances(),
})

export { getInstances, instancesQuery }

import { queryOptions } from '@tanstack/react-query'
import { api } from '#/integrations/axios/instance'
import { createServerFn } from '@tanstack/react-start'
import type { Item } from '../game/weapon'

const getItems = createServerFn().handler(async () => {
  const response = await api.get<Item[]>(`/api/items`)
  return response.data
})

const itemsQuery = queryOptions({
  queryKey: ['get-items'],
  queryFn: () => getItems(),
  select: (data) =>
    data.sort((a, b) => a.entity.name.localeCompare(b.entity.name)),
  staleTime: 60000,
  gcTime: 60000,
})

export { getItems, itemsQuery }

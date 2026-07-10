import { api } from '#/integrations/axios/instance'
import { queryOptions, useQuery } from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import axios from 'axios'

export type User = {
  id: string
  username: string
  email: string
  created_at: string
}

const getUser = createServerFn().handler(async () => {
  try {
    const response = await api.get<User>(`/api/auth/me`)
    return response.data
  } catch (error: unknown) {
    const status = axios.isAxiosError(error)
      ? error.response?.status
      : (error as { status?: number } | null)?.status

    if (status === 401) {
      return null
    }

    throw error
  }
})

export const userQuery = queryOptions({
  queryKey: ['user'],
  queryFn: () => getUser(),
})

export function useUser() {
  return useQuery(userQuery)
}

import { setResponseCookie } from '#/utils/set-cookie'
import {
  mutationOptions,
  useMutation,
  useQueryClient,
} from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import z from 'zod'
import type { User } from '../queries/auth'
import { api } from '#/integrations/axios/instance'

const requestSchema = z.object({
  email: z.string(),
  password: z.string(),
})

const login = createServerFn({ method: 'POST' })
  .validator(requestSchema)
  .handler(async ({ data }) => {
    const response = await api.post<User>('/api/auth/login', data)
    setResponseCookie(response)
    return response.data
  })

function useLogin() {
  const queryClient = useQueryClient()
  return useMutation(
    mutationOptions({
      mutationKey: ['login'],
      mutationFn: async (data: z.output<typeof requestSchema>) => {
        return await login({ data })
      },
      onSuccess: (user) => {
        queryClient.setQueryData(['me'], user)
        queryClient.removeQueries({ queryKey: ['teams'] })
      },
    })
  )
}

export { useLogin }

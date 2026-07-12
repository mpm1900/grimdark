import type { User } from '#/lib/queries/auth'
import { api } from '#/integrations/axios/instance'
import { setResponseCookie } from '#/utils/set-cookie'
import {
  mutationOptions,
  useMutation,
  useQueryClient,
} from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import z from 'zod'

const requestSchema = z.object({
  username: z.string(),
  email: z.string(),
  password: z.string(),
  secret: z.string(),
})

const signup = createServerFn({ method: 'POST' })
  .validator(requestSchema)
  .handler(async ({ data }) => {
    const response = await api.post<User>('/api/auth/signup', data)
    setResponseCookie(response)
  })

function useSignup() {
  const queryClient = useQueryClient()
  return useMutation(
    mutationOptions({
      mutationKey: ['signup'],
      mutationFn: async (data: z.output<typeof requestSchema>) => {
        return await signup({ data })
      },
      onSuccess: (user) => {
        queryClient.setQueryData(['me'], user)
        queryClient.removeQueries({ queryKey: ['teams'] })
      },
    })
  )
}

export { useSignup }

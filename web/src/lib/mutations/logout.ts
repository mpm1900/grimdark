import { api } from '#/integrations/axios/instance'
import { setResponseCookie } from '#/utils/set-cookie'
import {
  mutationOptions,
  useMutation,
  useQueryClient,
} from '@tanstack/react-query'
import { useRouter } from '@tanstack/react-router'
import { createServerFn } from '@tanstack/react-start'

const logout = createServerFn({ method: 'POST' }).handler(async () => {
  const response = await api.post('/api/auth/logout')
  setResponseCookie(response)
})

function useLogout() {
  const queryClient = useQueryClient()
  const router = useRouter()

  return useMutation(
    mutationOptions({
      mutationKey: ['logout'],
      mutationFn: async () => {
        await logout()
      },
      onSuccess: () => {
        queryClient.setQueryData(['me'], null)
        queryClient.removeQueries({ queryKey: ['teams'] })
        router.navigate({ to: '/login' })
      },
    })
  )
}

export { useLogout }

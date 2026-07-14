import { api } from '#/integrations/axios/instance'
import {
  mutationOptions,
  useMutation,
  useQueryClient,
} from '@tanstack/react-query'
import { createServerFn } from '@tanstack/react-start'
import z from 'zod'
import type { ID } from '../game/core'

const deleteTeam = createServerFn()
  .validator(z.uuid())
  .handler(async ({ data: id }) => {
    const response = await api.delete(`/api/teams/${id}`)
    return response.data
  })

function useDeleteTeam() {
  const qc = useQueryClient()
  return useMutation(
    mutationOptions({
      mutationKey: ['delete-team'],
      mutationFn: async (id: ID) => {
        const response = await deleteTeam({ data: id })
        await qc.invalidateQueries({ queryKey: ['teams'] })
        return response
      },
    })
  )
}

export { deleteTeam, useDeleteTeam }

import { createServerFn } from '@tanstack/react-start'
import z from 'zod'
import type { Instance } from './get-instances'
import { api } from '#/integrations/axios/instance'

const getInstance = createServerFn()
  .validator(z.uuid())
  .handler(async ({ data: id }) => {
    const response = await api.get<Instance>(`/api/instance/${id}`)
    return response.data
  })

export { getInstance }

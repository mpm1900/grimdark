import { createServerFn } from '@tanstack/react-start'
import z from 'zod'
import { api } from '#/integrations/axios/instance'
import axios, { AxiosError } from 'axios'
import type { Instance } from '../game/instance'

const getInstance = createServerFn()
  .validator(z.uuid())
  .handler(async ({ data: id }) => {
    try {
      const response = await api.get<Instance>(`/api/instance/${id}`)
      return response.data
    } catch (e) {
      if (axios.isAxiosError(e)) {
        const error: AxiosError = e
        if (error.status == 404) {
          return null
        }

        throw error
      }
    }
  })

export { getInstance }

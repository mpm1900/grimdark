import { getApiBaseUrl } from '#/utils/get-api-base-url'
import { createServerFn } from '@tanstack/react-start'
import z from 'zod'

const getInstance = createServerFn()
  .validator(z.string())
  .handler(async ({ data: id }) => {
    const response = await fetch(`${getApiBaseUrl()}/api/instances/${id}`)

    if (response.status === 404) {
      throw new Response('Instance not found', { status: 404 })
    }

    if (response.status === 204) {
      throw new Response('Instance not found', { status: 404 })
    }

    if (!response.ok) {
      throw new Error(`Failed to fetch instance: ${response.status}`)
    }

    return response.json()
  })

export { getInstance }

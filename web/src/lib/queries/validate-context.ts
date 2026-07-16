import { queryOptions } from '@tanstack/react-query'
import { contextToString, type Context } from '../game/context'
import { promisify } from '../socket/promise'

async function validateContext(context: Context): Promise<boolean> {
  const response = await promisify({
    type: 'validate-context',
    context,
  })
  return !!response.valid
}

function validateContextQuery(context: Context) {
  return queryOptions<boolean>({
    queryKey: ['validate-context', contextToString(context)],
    queryFn: async () => {
      return validateContext(context)
    },
    staleTime: 0,
    gcTime: 0,
  })
}

export { validateContext, validateContextQuery }

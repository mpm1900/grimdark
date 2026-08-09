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

function validateContextQuery(
  context: Context,
  deps: (boolean | number | string)[] = []
) {
  return queryOptions<boolean>({
    queryKey: ['validate-context', contextToString(context), ...deps],
    queryFn: async () => {
      return validateContext(context)
    },
    staleTime: Infinity,
    gcTime: 5 * 60 * 1000,
  })
}

export { validateContext, validateContextQuery }

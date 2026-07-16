import { queryOptions } from '@tanstack/react-query'
import { contextToString, NULL_CONTEXT, type Context } from '../game/context'
import type { ID } from '../game/core'
import { promisify } from '../socket/promise'

async function getTargets(context: Context): Promise<Context> {
  const response = await promisify({
    type: 'get-targets',
    context,
  })

  return response.context!
}

function getTargetsQuery(
  source_ID: ID | null | undefined,
  player_ID: ID | null | undefined,
  action_ID: ID,
  deps: (boolean | number | string)[],
  context_overrides: Partial<Context> = {}
) {
  const context: Context = {
    ...NULL_CONTEXT,
    source_ID: source_ID ?? null,
    parent_ID: source_ID ?? null,
    player_ID: player_ID ?? null,
    ...context_overrides,
    action_ID,
  }
  return queryOptions<Context>({
    queryKey: ['get-targets', contextToString(context), ...deps],
    queryFn: async () => {
      return getTargets(context)
    },
  })
}

export { getTargets, getTargetsQuery }

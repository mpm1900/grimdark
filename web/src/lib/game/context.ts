import z from 'zod'

const ContextSchema = z.object({
  action_ID: z.string(),
  source_player_ID: z.string(),

  parent_actor_ID: z.string(),
  source_actor_ID: z.string(),

  target_actor_IDs: z.array(z.string()),
  target_position_IDs: z.array(z.string()),
})

type Context = z.output<typeof ContextSchema>

function contextToString(c: Context): string {
  return `${c.action_ID}.${c.parent_actor_ID}.${c.source_actor_ID}.${c.source_player_ID}.${c.target_actor_IDs?.join('+')}.${c.target_position_IDs?.join('+')}`
}

const NULL_CONTEXT: Context = {
  action_ID: '',
  parent_actor_ID: '',
  source_actor_ID: '',
  source_player_ID: '',
  target_actor_IDs: [],
  target_position_IDs: [],
}

export { ContextSchema, contextToString, NULL_CONTEXT }
export type { Context }

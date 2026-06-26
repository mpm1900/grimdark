import z from 'zod'

const ContextSchema = z.object({
  action_ID: z.string().nullable(),
  effect_ID: z.string().nullable().optional(),
  player_ID: z.string().nullable(),

  parent_ID: z.string().nullable(),
  source_ID: z.string().nullable(),

  actor_IDs: z.array(z.string().nullable()),
  position_IDs: z.array(z.string().nullable()),
})

type Context = z.output<typeof ContextSchema>

function contextToString(c: Context): string {
  return `${c.action_ID ?? ''}.${c.parent_ID ?? ''}.${c.source_ID ?? ''}.${c.player_ID ?? ''}.${c.actor_IDs?.filter(Boolean).join('+')}.${c.position_IDs?.filter(Boolean).join('+')}`
}

const NULL_CONTEXT: Context = {
  action_ID: null,
  effect_ID: null,
  parent_ID: null,
  source_ID: null,
  player_ID: null,
  actor_IDs: [],
  position_IDs: [],
}

export { ContextSchema, contextToString, NULL_CONTEXT }
export type { Context }

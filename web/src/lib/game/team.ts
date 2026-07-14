import z from 'zod'

export const ActorConfigSchema = z.object({
  class: z.uuid().nullable(),
  items: z.uuid().array(),
  name: z.string(),
  weapon_l: z.uuid().nullable(),
  weapon_r: z.uuid().nullable(),
})

export const TeamConfigSchema = z.object({
  name: z.string(),
  actors: z.array(ActorConfigSchema),
})

export const TeamSchema = z.object({
  ID: z.uuid().nullable(),
  config: TeamConfigSchema,
})

export type ActorConfig = z.infer<typeof ActorConfigSchema>
export type TeamConfig = z.infer<typeof TeamConfigSchema>
export type Team = z.infer<typeof TeamSchema>

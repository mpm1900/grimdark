import z from 'zod'

export const authSchema = z.object({
  email: z.email(),
  password: z.string().min(4),
  secret: z.string(),
})

export const signupSchema = authSchema.extend({
  username: z.string().min(1),
  secret: z.string().min(1),
})

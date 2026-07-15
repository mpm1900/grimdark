import type { Bindable, ID } from './core'

export type Effect = {
  ID: ID
  name: string
  description: string
  delay: number | null
  duration: number | null
  priority: number
}

export type Modifier = Bindable<Effect>

import type { Bindable, Entity, ID } from './core'

export type Effect = {
  ID: ID
  entity: Entity
  delay: number | null
  duration: number | null
  priority: number
}

export type Modifier = Bindable<Effect>

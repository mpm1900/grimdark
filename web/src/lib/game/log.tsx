import type React from 'react'
import type { Bindable } from './core'

export type Log = {
  template: string
  terms: Record<string, string>
}

export function RenderLog(log: Bindable<Log>): React.ReactNode {
  const { template, terms } = log.payload
  const keys = Object.keys(terms).sort((a, b) => b.length - a.length)

  if (keys.length === 0) {
    return template
  }

  const nodes: React.ReactNode[] = [template]

  for (const key of keys) {
    for (let i = 0; i < nodes.length; i++) {
      const node = nodes[i]
      if (typeof node !== 'string') {
        continue
      }

      const parts = node.split(key)
      if (parts.length === 1) {
        continue
      }

      nodes.splice(
        i,
        1,
        ...parts.flatMap((part, partIndex) => {
          if (partIndex === parts.length - 1) {
            return [part]
          }

          return [
            part,
            <span key={`${log.ID}-${key}-${i}-${partIndex}`}>{terms[key]}</span>,
          ]
        })
      )
    }
  }

  return <>{nodes}</>
}

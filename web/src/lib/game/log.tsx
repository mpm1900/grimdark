import type React from 'react'
import type { Bindable } from './core'
import { gameStore } from '../stores/game'
import { clientsStore } from '../stores/clients'
import { cn } from '../utils'
import { getTargetsFromContext, type Context } from './context'
import type { Game } from './game'
import { Marker, MarkerContent } from '#/components/ui/marker'
import { setHoverPosition } from '../stores/ui'

export type Log = {
  depth: number
  template: string
  terms: Record<string, string>
  type: string
}

function RenderTerm({
  context,
  game,
  terms,
  term_key,
}: {
  context: Context
  game: Game
  terms: Record<string, string>
  term_key: string
}) {
  const client_ID = clientsStore.get().me?.ID
  switch (term_key) {
    case '$source$':
      const source = game.actors.find((a) => a.ID === context.source_ID)
      return (
        <span
          onMouseEnter={() => {
            setHoverPosition(source?.position_ID ?? null)
          }}
          onMouseLeave={() => {
            setHoverPosition(null)
          }}
          className={cn('hover:underline cursor-default', {
            'text-ally/60': source?.player_ID === client_ID,
            'text-enemy/60': source?.player_ID !== client_ID,
          })}
        >
          {terms[term_key]}
        </span>
      )
    case '$target$':
      const targets = getTargetsFromContext(game.actors, context)
      const target = targets[0]
      return (
        <span
          onMouseEnter={() => {
            setHoverPosition(target.position_ID)
          }}
          onMouseLeave={() => {
            setHoverPosition(null)
          }}
          className={cn('hover:underline cursor-default', {
            'text-ally/60': target?.player_ID === client_ID,
            'text-enemy/60': target?.player_ID !== client_ID,
          })}
        >
          {terms[term_key]}
        </span>
      )
    default:
      return <span>{terms[term_key]}</span>
  }
}

export function RenderLog(log: Bindable<Log>): React.ReactNode {
  const game = gameStore.get()
  const { template, terms, type } = log.payload
  const keys = Object.keys(terms).sort((a, b) => b.length - a.length)

  if (type === 'trigger') {
    // return <span className='font-cinzel font-semibold text-center'>{template.replace(/-/, ' ')}:</span>
  }

  if (type === 'turn') {
    return (
      <Marker variant="separator">
        <MarkerContent>{template}</MarkerContent>
      </Marker>
    )
  }

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
            <RenderTerm
              key={`${log.ID}-${key}-${i}-${partIndex}`}
              context={log.context}
              game={game}
              terms={terms}
              term_key={key}
            />,
          ]
        })
      )
    }
  }

  return <>{nodes}</>
}

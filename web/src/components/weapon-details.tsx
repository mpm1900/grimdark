import type { Stat } from '#/lib/game/core'
import type { Weapon } from '#/lib/game/weapon'
import { cn, sign } from '#/lib/utils'
import { entries } from '#/utils/maps'
import { GothicFrame } from './gothic-ui/frame'
import { Item, ItemContent, ItemDescription, ItemTitle } from './ui/item'

function InlineAuxStats({
  aux_stats,
}: {
  aux_stats: Partial<Record<Stat, number>>
}) {
  return (
    <span className="flex gap-2">
      {entries(aux_stats).map(([stat, aux], i) => (
        <span
          key={i}
          className={cn('capitalize', {
            'text-positive': aux > 0,
            'text-red-400': aux < 0,
          })}
        >
          {sign(aux)}
          {stat}: {Math.abs(aux)}
        </span>
      ))}
    </span>
  )
}

function WeaponDetails({ weapon }: { weapon: Weapon }) {
  return (
    <GothicFrame>
      <ItemContent>
        <span className="font-cinzel font-semibold">{weapon.name}</span>
        <div className="font-serif">
          <ItemDescription>
            <span>Actions:</span>{' '}
            {weapon.actions.map((a) => a.config.name).join(', ')}
          </ItemDescription>
          <ItemDescription>
            Effects:{' '}
            {weapon.effects.map((e) => e.name).join(', ') || (
              <span className="italic opacity-50">None</span>
            )}
          </ItemDescription>
          <ItemDescription>
            <InlineAuxStats aux_stats={weapon.aux_stats} />
          </ItemDescription>
        </div>
      </ItemContent>
    </GothicFrame>
  )
}

export { WeaponDetails }

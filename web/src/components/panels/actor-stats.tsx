import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { ActorFrame } from '../actor-frame'
import { AffinitiesTable } from '../affinities-table'
import { GothicBadge } from '../gothic-ui/badge'
import { GothicHighlightFrame } from '../gothic-ui/frame'
import { OtherStatsTable, StatsTable } from '../stats-table'
import { Marker, MarkerContent } from '../ui/marker'
import { WeaponDetails } from '../weapon-details'
import { PanelHeader } from './panel-header'
import { gameStore } from '#/lib/stores/game'
import { getAppliedEffects } from '#/lib/game/game'
import { EffectTooltip } from '../effect-tooltip'

function ActorStatsPanel({
  actor,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor: Actor }) {
  const game = useSelector(gameStore, (g) => g)
  const applied_effects = getAppliedEffects(game, actor)
  return (
    <div className={cn('grid grid-cols-3 relative', className)} {...props}>
      <PanelHeader>{actor.name}'s Stats</PanelHeader>
      <div className="font-serif text-foreground/80">
        <ActorFrame actor={actor} className="-ml-1 -mt-1.5 -mr-px z-0" />
        <div className="py-4 p-2 hidden">
          {actor.weapon_l && <WeaponDetails weapon={actor.weapon_l} />}
        </div>
        <div className="px-1 py-2">
          <Marker variant="separator">
            <MarkerContent>Other stats</MarkerContent>
          </Marker>
          <OtherStatsTable actor={actor} />
        </div>
      </div>
      <GothicHighlightFrame className="-mx-3 -mt-px gap-4">
        <StatsTable actor={actor} />
        <div className="px-2 max-w-72 font-serif">
          <Marker variant="separator">
            <MarkerContent>Active Effects</MarkerContent>
          </Marker>

          <div className="min-w-0 flex flex-row flex-wrap gap-2 p-2">
            {applied_effects.map((effect) => (
              <EffectTooltip key={effect.ID} effect={effect} asChild>
                <GothicBadge variant="empty" className="capitalize">
                  {effect.name}
                  {effect.count > 1 && `(${effect.count})`}
                </GothicBadge>
              </EffectTooltip>
            ))}
            {applied_effects.length === 0 && (
              <span className="italic text-muted-foreground">None</span>
            )}
          </div>
        </div>
      </GothicHighlightFrame>
      <AffinitiesTable actor={actor} className="pl-3.5" />
    </div>
  )
}

export { ActorStatsPanel }

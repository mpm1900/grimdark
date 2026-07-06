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
          {actor.weapon && <WeaponDetails weapon={actor.weapon} />}
        </div>
        <div className='px-1 pt-2'>
          <Marker variant="separator">
            <MarkerContent>Other stats</MarkerContent>
          </Marker>
          <OtherStatsTable actor={actor} />
        </div>
        <div className='px-1'>

          <Marker variant="separator">
            <MarkerContent>Active Effects</MarkerContent>
          </Marker>

          <div className="min-w-0 flex flex-row flex-wrap gap-2 p-2">
            {applied_effects.map((effect) => (
              <GothicBadge variant="empty" key={effect.ID} className="capitalize">
                {effect.name}
                {effect.count > 1 && `(${effect.count})`}
              </GothicBadge>
            ))}
            {applied_effects.length === 0 && (
              <span className="italic text-muted-foreground">None</span>
            )}
          </div>
        </div>
      </div>
      <GothicHighlightFrame className="-mx-3 -mt-px">
        <StatsTable actor={actor} />
      </GothicHighlightFrame>
      <AffinitiesTable actor={actor} className="pl-3.5" />
    </div>
  )
}

export { ActorStatsPanel }

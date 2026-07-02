import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { ActorFrame } from '../actor-frame'
import { AffinitiesTable } from '../affinities-table'
import { GothicHighlightFrame } from '../gothic-ui/frame'
import { StatsTable } from '../stats-table'
import { WeaponDetails } from '../weapon-details'
import { PanelHeader } from './panel-header'

function ActorStatsPanel({
  actor,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor: Actor }) {
  return (
    <div className={cn('grid grid-cols-3 relative', className)} {...props}>
      <PanelHeader>{actor.name}'s Stats</PanelHeader>
      <div className="font-serif text-foreground/60">
        <ActorFrame
          actor={actor}
          className="-ml-1 -mt-1.5 -mr-px z-0"
        />
        <div className="py-4 p-2">
          {actor.weapon && <WeaponDetails weapon={actor.weapon} />}
        </div>
      </div>
      <GothicHighlightFrame className="-mx-3 -mt-px">
        <StatsTable actor={actor} />
      </GothicHighlightFrame>
      <AffinitiesTable actor={actor} className='pl-3.5' />
    </div>
  )
}

export { ActorStatsPanel }

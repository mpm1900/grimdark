import { useSelector } from '@tanstack/react-store'
import { Platform, PlatformParent } from './platform'
import { teamStore } from '#/lib/stores/team'

function TeamPlatforms() {
  const team = useSelector(teamStore, (s) => s)
  return (
    <PlatformParent className="absolute inset-0 bottom-3/10 z-0 perspective-origin-top">
      {team.actors.map((_, i) => (
        <Platform
          key={i}
          rank={i}
          variant={
            i === team.active_actor
              ? 'player-active'
              : i === team.hover_actor
                ? 'player-hover'
                : 'player'
          }
          className="flex-1"
        >
          {''}
        </Platform>
      ))}
    </PlatformParent>
  )
}

export { TeamPlatforms }

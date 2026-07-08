import { HiLink } from 'react-icons/hi'
import { ActorAvatar } from '../actor-avatar'
import { TinyBadge } from '../gothic-ui/badge'
import { GothicCard } from '../gothic-ui/card'
import { GothicFrame } from '../gothic-ui/frame'
import { WeaponFrame } from '../weapon-details'
import { ActionContextDialog } from '../action-context-dialog'
import { DialogTrigger } from '../ui/dialog'
import { ActionButton, SystemActionButton } from '../action-button'
import { BattleLog } from '../battle-log'
import { useSelector } from '@tanstack/react-store'
import { uiStore } from '#/lib/stores/ui'
import { gameStore } from '#/lib/stores/game'

function BattlePanel() {
  const active_actor_id = useSelector(uiStore, (s) => s.active_actor)
  const game = useSelector(gameStore, (g) => g)
  const active_actor = game.actors.find((a) => a.ID === active_actor_id)
  const acitve_actor_command = game.commands.find(
    (c) => c.context.source_ID === active_actor?.ID
  )
  const weapons = [
    active_actor?.weapon_l ?? null,
    active_actor?.weapon_r ?? null,
  ].sort((a, b) => (b?.hands ?? 0) - (a?.hands ?? 0))
  const main_weapon = weapons[0]
  const secondary_weapon = weapons[1]

  if (!active_actor) return null
  return (
    <div className="h-48 absolute bottom-0 left-0 right-0 flex items-start justify-center z-10">
      {active_actor && <ActorAvatar actor={active_actor} />}
      <GothicCard className="h-full flex flex-row">
        <div className="relative h-full grid grid-cols-1 grid-rows-3 w-13">
          <TinyBadge
            variant="default"
            className="absolute z-10 px-1 -top-1 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
          >
            Items
          </TinyBadge>
          <GothicFrame></GothicFrame>
          <GothicFrame></GothicFrame>
          <GothicFrame></GothicFrame>
        </div>
        <div className="relative h-full grid grid-cols-1 grid-rows-2">
          <TinyBadge
            variant="default"
            className="absolute z-10 px-1 -top-1 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
          >
            Weapons
          </TinyBadge>
          {main_weapon ? (
            <WeaponFrame disabled={false} weapon={main_weapon} />
          ) : (
            <GothicFrame></GothicFrame>
          )}
          {main_weapon?.hands == 2 ? (
            <WeaponFrame disabled={true} weapon={main_weapon} />
          ) : secondary_weapon ? (
            <WeaponFrame disabled={false} weapon={secondary_weapon} />
          ) : (
            <GothicFrame></GothicFrame>
          )}
          {main_weapon?.hands == 2 && (
            <div className="pointer-events-none absolute inset-0 grid place-items-center">
              <HiLink className="rotate-136 size-6 fill-neutral-500" />
            </div>
          )}
        </div>
      </GothicCard>
      <GothicCard className="relative flex-row h-full max-w-1/3 bg-neutral-950 z-10">
        <TinyBadge
          variant="default"
          className="absolute z-10 px-1 -top-1 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
        >
          Actions
        </TinyBadge>
        <div className="grid grid-cols-2 grid-rows-3">
          {active_actor?.actions
            .filter((a) => a.type === 'actor')
            .map((action) => (
              <ActionContextDialog
                key={action.ID}
                actor={active_actor}
                action={action}
                enabled={!action.is_disabled}
              >
                <DialogTrigger asChild>
                  <ActionButton
                    action={action}
                    actor={active_actor}
                    disabled={!!acitve_actor_command}
                  />
                </DialogTrigger>
              </ActionContextDialog>
            ))}
        </div>
        <div className="flex flex-col justify-between">
          {active_actor?.actions
            .filter((a) => a.type === 'system')
            .map((action) => (
              <ActionContextDialog
                key={action.ID}
                actor={active_actor}
                action={action}
                enabled={!action.is_disabled}
              >
                <DialogTrigger asChild>
                  <SystemActionButton
                    action={action}
                    actor={active_actor}
                    disabled={!!acitve_actor_command}
                  />
                </DialogTrigger>
              </ActionContextDialog>
            ))}
        </div>
      </GothicCard>
      <GothicCard className="h-full min-w-0 max-w-1/4 flex-1 flex bg-neutral-950 p-0">
        <TinyBadge
          variant="default"
          className="absolute z-10 px-1 -top-1 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
        >
          Battle Log
        </TinyBadge>
        <BattleLog />
      </GothicCard>
    </div>
  )
}

export { BattlePanel }

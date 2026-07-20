import { HiLink } from 'react-icons/hi'
import { ActorAvatar } from '../actor-avatar'
import { TinyBadge } from '../gothic-ui/badge'
import { GothicCard } from '../gothic-ui/card'
import { GothicFrame } from '../gothic-ui/frame'
import { WeaponFrame, WeaponFrameExt } from '../weapon-details'
import { ActionContextDialog } from '../action-context-dialog'
import { DialogTrigger } from '../ui/dialog'
import { ActionButton, SystemActionButton } from '../action-button'
import { BattleLog } from '../battle-log'
import { useSelector } from '@tanstack/react-store'
import { uiStore } from '#/lib/stores/ui'
import { gameStore } from '#/lib/stores/game'
import { ActorLore } from '../actor-lore'
import type { Actor } from '#/lib/game/actor'
import { sendContextMessage } from '#/lib/stores/socket'
import { GothicFramedButton } from '../gothic-ui/button'
import { NULL_CONTEXT } from '#/lib/game/context'
import { lobbyStore } from '#/lib/stores/clients'
import { Check, ChevronsRight, Loader } from 'lucide-react'
import { v4 } from 'uuid'

function ActionsPanel({ active_actor }: { active_actor: Actor }) {
  const game = useSelector(gameStore, (g) => g)
  const acitve_actor_commands = game.commands.filter(
    (c) => c.context.source_ID === active_actor?.ID
  )
  const remaining_ap = active_actor.stats.actions - acitve_actor_commands.length
  return (
    <GothicCard className="relative flex-row h-full max-w-1/3 bg-neutral-950 z-10">
      <TinyBadge
        variant="default"
        className="absolute z-10 px-1 -top-1.5 border-t-0 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
      >
        Actions ({active_actor.stats.actions})
      </TinyBadge>
      <div className="grid grid-cols-2 grid-rows-3">
        {active_actor?.actions
          .filter((a) => a.tags.includes('actor'))
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
                  disabled={remaining_ap === 0 || active_actor.is_stunned}
                />
              </DialogTrigger>
            </ActionContextDialog>
          ))}
      </div>
      <div className="flex flex-col justify-between">
        {active_actor?.actions
          .filter((a) => a.tags.includes('system'))
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
                  disabled={remaining_ap === 0 || active_actor.is_stunned}
                />
              </DialogTrigger>
            </ActionContextDialog>
          ))}
      </div>
    </GothicCard>
  )
}

function BattlePanel() {
  const client = useSelector(lobbyStore, (s) => s.client)
  const active_actor_id = useSelector(uiStore, (s) => s.active_actor)
  const game = useSelector(gameStore, (g) => g)
  const game_status = game.status
  const player = game.players.find((p) => p.ID === client?.ID)
  const active_actor = game.actors.find((a) => a.ID === active_actor_id)
  const commands = game.commands.filter(
    (c) => c.context.player_ID === client?.ID
  )
  const actions = game.actors
    .filter((a) => !a.is_stunned && a.player_ID === client?.ID && a.position_ID)
    .map((a) => a.stats.actions)
    .reduce((sum, c) => sum + c, 0)

  const weapons = [
    active_actor?.weapon_l ?? null,
    active_actor?.weapon_r ?? null,
  ].sort((a, b) => (b?.weight ?? 0) - (a?.weight ?? 0))
  const main_weapon = weapons[0]
  const secondary_weapon = weapons[1]

  return (
    <div className="h-48 absolute bottom-0 left-0 right-0 flex items-start justify-center z-10">
      {active_actor && <ActorAvatar actor={active_actor} />}
      {active_actor?.is_player && (
        <GothicCard className="h-full flex flex-row z-10">
          <div className="relative h-full grid grid-cols-1 grid-rows-3 w-13">
            <TinyBadge
              variant="default"
              className="absolute z-10 px-1 -top-1.5 border-t-0 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
            >
              Items
            </TinyBadge>
            <GothicFrame></GothicFrame>
            <GothicFrame></GothicFrame>
            <GothicFrame></GothicFrame>
          </div>
          <div className="relative">
            <TinyBadge
              variant="default"
              className="absolute z-20 px-1 -top-1.5 border-t-0 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
            >
              Weapons
            </TinyBadge>
            <div className="relative h-full grid grid-cols-1 grid-rows-2 overflow-hidden min-w-18">
              {main_weapon ? (
                <WeaponFrame disabled={false} weapon={main_weapon} />
              ) : (
                <GothicFrame></GothicFrame>
              )}
              {main_weapon?.weight == 2 ? (
                <WeaponFrameExt weapon={main_weapon} disabled />
              ) : secondary_weapon ? (
                <WeaponFrame disabled={false} weapon={secondary_weapon} />
              ) : (
                <GothicFrame></GothicFrame>
              )}
              {main_weapon?.weight == 2 && (
                <div className="pointer-events-none absolute inset-0 grid place-items-center z-20">
                  <HiLink className="rotate-136 size-6 fill-foreground/60" />
                </div>
              )}
            </div>
          </div>
        </GothicCard>
      )}
      {active_actor?.is_player && <ActionsPanel active_actor={active_actor} />}
      {active_actor && !active_actor.is_player && (
        <GothicCard className="h-full w-1/3 z-10">
          <ActorLore actor={active_actor} className="" />
        </GothicCard>
      )}
      <GothicCard className="h-full min-w-0 max-w-1/4 flex-1 flex bg-neutral-950 p-0">
        <TinyBadge
          variant="default"
          className="absolute z-10 px-1 -top-1.5 border-t-0 left-1/2 -translate-x-1/2 rounded-xs rounded-b-sm font-cinzel border-white/30 text-foreground/60"
        >
          Battle Log
        </TinyBadge>
        <BattleLog />
      </GothicCard>
      {client && (
        <div className="grid place-items-center h-full -ml-1">
          <GothicFramedButton
            variant="red"
            className="flex flex-col p-0 px-1 h-18 w-18 text-lg gap-0 font-serif text-foreground/60"
            disabled={game_status != 'idle' || !player || player.ready}
            onClick={() => {
              sendContextMessage({
                request_ID: v4(),
                type: 'turn-ready',
                client_ID: client.ID,
                context: NULL_CONTEXT,
              })
            }}
          >
            {game.status === 'idle' && (
              <span>
                {commands.length}/{actions}
              </span>
            )}
            {game.status === 'idle' && !player?.ready && <ChevronsRight />}
            {game.status === 'idle' && player?.ready && <Check />}
            {game.status === 'running' && <Loader className="animate-spin" />}
            {game.status === 'waiting' && <Loader className="animate-spin" />}
          </GothicFramedButton>
        </div>
      )}
    </div>
  )
}

export { BattlePanel }

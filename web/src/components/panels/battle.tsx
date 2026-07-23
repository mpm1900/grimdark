import { HiLink } from 'react-icons/hi'
import { useEffect, useMemo, useState } from 'react'
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
import type { Action } from '#/lib/game/action'
import { sendContextMessage } from '#/lib/stores/socket'
import { GothicFramedButton } from '../gothic-ui/button'
import { NULL_CONTEXT } from '#/lib/game/context'
import { lobbyStore } from '#/lib/stores/clients'
import {
  Check,
  ChevronLeft,
  ChevronRight,
  ChevronsRight,
  Loader,
} from 'lucide-react'
import { v4 } from 'uuid'
import {
  Carousel,
  type CarouselApi,
  CarouselContent,
  CarouselItem,
} from '../ui/carousel'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '../ui/tabs'

const ACTIONS_PER_PAGE = 6

function chunkActions(actions: Action[]) {
  return Array.from(
    { length: Math.ceil(actions.length / ACTIONS_PER_PAGE) },
    (_, index) =>
      actions.slice(
        index * ACTIONS_PER_PAGE,
        index * ACTIONS_PER_PAGE + ACTIONS_PER_PAGE
      )
  )
}

function ActionsPanel({ active_actor }: { active_actor: Actor }) {
  const game = useSelector(gameStore, (g) => g)
  const acitve_actor_commands = game.commands.filter(
    (c) => c.context.source_ID === active_actor?.ID
  )
  const remaining_ap = active_actor.stats.actions - acitve_actor_commands.length
  const actor_actions = useMemo(
    () => active_actor.actions.filter((a) => a.tags.includes('actor')),
    [active_actor.actions]
  )
  const action_pages = useMemo(
    () => chunkActions(actor_actions),
    [actor_actions]
  )
  const [carousel_api, set_carousel_api] = useState<CarouselApi>()
  const [can_scroll_previous, set_can_scroll_previous] = useState(false)
  const [can_scroll_next, set_can_scroll_next] = useState(false)

  useEffect(() => {
    if (!carousel_api) return

    const update_scroll_buttons = () => {
      set_can_scroll_previous(carousel_api.canScrollPrev())
      set_can_scroll_next(carousel_api.canScrollNext())
    }

    update_scroll_buttons()
    carousel_api.on('select', update_scroll_buttons)
    carousel_api.on('reInit', update_scroll_buttons)

    return () => {
      carousel_api.off('select', update_scroll_buttons)
      carousel_api.off('reInit', update_scroll_buttons)
    }
  }, [carousel_api])

  useEffect(() => {
    carousel_api?.scrollTo(0)
  }, [active_actor.ID, carousel_api])

  return (
    <GothicCard className="relative flex-row h-full w-1/3 bg-neutral-950 z-10">
      <Tabs defaultValue="actions" className="relative h-full w-full">
        <TabsList className="absolute z-10 select-none -top-1.75 border-t-0 left-1/2 -translate-x-1/2 flex -space-x-px h-auto rounded-none bg-transparent p-0 group-data-[orientation=horizontal]/tabs:h-auto">
          <TabsTrigger
            value="actions"
            asChild
            className="w-18 text-center bg-neutral-950! data-[state=active]:bg-neutral-900! text-[10px] leading-3"
          >
            <TinyBadge
              variant="default"
              className="block rounded-xs py-0 px-1 rounded-b-sm font-cinzel border-white/30! text-foreground/60"
            >
              Actions
            </TinyBadge>
          </TabsTrigger>
          <TabsTrigger
            value="info"
            asChild
            className="w-18 text-center bg-neutral-950! data-[state=active]:bg-neutral-900! text-[10px] leading-3"
          >
            <TinyBadge
              variant="default"
              className="block rounded-xs py-0 px-1 rounded-b-sm font-cinzel border-white/30! text-foreground/60"
            >
              Info
            </TinyBadge>
          </TabsTrigger>
        </TabsList>
        <TabsContent value="actions" className="relative h-full flex">
          <Carousel
            setApi={set_carousel_api}
            className="h-full min-w-0 flex-1 *:data-[slot=carousel-content]:h-full"
          >
            <CarouselContent className="ml-0 h-full">
              {action_pages.map((actions, page_index) => (
                <CarouselItem key={page_index} className="pl-0">
                  <div className="grid h-full grid-cols-2 grid-rows-3">
                    {actions.map((action) => (
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
                            disabled={
                              remaining_ap === 0 || active_actor.is_stunned
                            }
                          />
                        </DialogTrigger>
                      </ActionContextDialog>
                    ))}
                  </div>
                </CarouselItem>
              ))}
            </CarouselContent>
            <div className="absolute -bottom-3 left-1/2 -translate-x-1/2 flex">
              <GothicFramedButton
                variant="red"
                className="p-0 size-6"
                disabled={!can_scroll_previous}
                onClick={() => carousel_api?.scrollPrev()}
              >
                <ChevronLeft />
              </GothicFramedButton>
              <GothicFramedButton
                variant="red"
                className="p-0 size-6"
                disabled={!can_scroll_next}
                onClick={() => carousel_api?.scrollNext()}
              >
                <ChevronRight />
              </GothicFramedButton>
            </div>
          </Carousel>
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
        </TabsContent>
        <TabsContent value="info" className="relative w-full">
          <ActorLore actor={active_actor} className="" />
        </TabsContent>
      </Tabs>
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

  const weapons = Object.values(active_actor?.weapons ?? {}).sort(
    (a, b) => (b?.weight ?? 0) - (a?.weight ?? 0)
  )
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

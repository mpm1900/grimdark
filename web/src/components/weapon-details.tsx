import { STAT_LABELS, type Stat } from '#/lib/game/core'
import type { Weapon } from '#/lib/game/weapon'
import { cn, sign } from '#/lib/utils'
import { entries } from '#/utils/maps'
import { cva } from 'class-variance-authority'
import { GothicFrame, GothicShadowFrame } from './gothic-ui/frame'
import { ItemDescription } from './ui/item'
import { statVariants } from './stat-name'
import { HoverCard, HoverCardTrigger } from './ui/hover-card'
import { GothicHoverCardContent } from './gothic-ui/hover-card'
import { DNumber } from './dnumber'
import { WEAPON_ICONS } from '#/icons/weapons'
import { EffectTooltip } from './effect-tooltip'
import { ActionTooltip } from './action-tooltip'

function InlineOffsetStats({
  offset_stats,
  className,
  ...props
}: React.ComponentProps<'span'> & {
  offset_stats: Partial<Record<Stat, number>>
}) {
  return (
    <span className={cn('flex flex-wrap gap-x-2', className)} {...props}>
      {entries(offset_stats).map(([stat, aux], i) => (
        <DNumber key={i} value={aux}>
          {sign(aux)}
          {Math.abs(aux)} {STAT_LABELS[stat]}
        </DNumber>
      ))}
    </span>
  )
}

const weaponWrapper = cva('relative p-0.5 font-serif', {
  variants: {
    rarity: {
      common: 'bg-gradient-to-b from-foreground/60 to-foreground/0',
      rare: 'bg-gradient-to-b from-emerald-500/60 to-emerald-500/0',
    },
  },
  defaultVariants: {
    rarity: 'common',
  },
})
const weaponBody = cva('relative bg-neutral-900 pb-1', {
  variants: {
    rarity: {
      common: 'shadow-[inset_2px_2px_18px_0px_rgba(0,_0,_0,_0.8)]',
      rare: 'shadow-[inset_-8px_8px_16px_0px_color-mix(in_oklab,var(--color-emerald-400)_10%,transparent)]',
    },
  },
  defaultVariants: {
    rarity: 'common',
  },
})
const weaponTitle = cva('font-cinzel font-semibold block text-md', {
  variants: {
    rarity: {
      common: 'text-foreground',
      rare: 'text-emerald-200/80',
    },
  },
  defaultVariants: {
    rarity: 'common',
  },
})

const weaponIcon = cva('absolute top-1 h-auto block', {
  variants: {
    rarity: {
      common: 'text-foreground/50',
      rare: 'text-emerald-300/50',
    },
    weapon_type: {
      sword: 'size-24 top-4 rotate-135',
      'big-sword': 'size-9 top-2',
      pistol: 'size-14 left-1/2 top-4',
      rifle: 'size-26 top-6 rotate-130',
    },
  },
  defaultVariants: {
    rarity: 'common',
    weapon_type: 'sword',
  },
})

function WeaponDetails({ weapon }: { weapon: Weapon }) {
  const rarity = 'rare'
  const Icon = WEAPON_ICONS[weapon.weapon_type]
  return (
    <div className={weaponWrapper({ rarity: rarity })}>
      <div className={weaponBody({ rarity: rarity })}>
        <div className="p-2">
          <div className="absolute z-10 bottom-0 left-2/3 right-0 top-2 overflow-hidden">
            {Icon && (
              <Icon
                className={weaponIcon({
                  rarity,
                  weapon_type: weapon.weapon_type,
                  className: 'left-1/2 -translate-x-1/2 fading-image z-30',
                })}
              />
            )}
          </div>
          <div className="pr-6">
            <span className={weaponTitle({ rarity: rarity })}>
              {weapon.name}
            </span>
            <span className="text-foreground/60 block text-xs leading-none">
              Common {weapon.weapon_type} ({weapon.weight})
            </span>
          </div>
        </div>
        <div className="text-foreground/80 italic text-xs px-6 py-2">
          {weapon.description}
        </div>
        <GothicShadowFrame className="z-10 mt-6 m-1 space-y-1">
          <ItemDescription className="text-foreground/80">
            <span className="text-foreground/40 block font-cinzel font-semibold">
              Actions
            </span>
            <span className="space-x-2 flex flex-wrap">
              {weapon.actions.map((a) => (
                <ActionTooltip
                  key={a.ID}
                  action={a}
                  className={statVariants({
                    stat: a.config.stat,
                    className: 'cursor-default hover:underline',
                  })}
                >
                  {a.config.name}
                </ActionTooltip>
              ))}
            </span>
          </ItemDescription>
          <ItemDescription className="text-foreground/80">
            <span className="text-foreground/40 block font-cinzel font-semibold">
              Effects
            </span>
            {weapon.effects.length > 0 && (
              <span className="space-x-2 flex flex-wrap">
                {weapon.effects.map((e) => (
                  <EffectTooltip
                    key={e.ID}
                    effect={e}
                    className="cursor-default hover:underline"
                  >
                    {e.name}
                  </EffectTooltip>
                ))}
              </span>
            )}
            <InlineOffsetStats offset_stats={weapon.offset_stats} />
          </ItemDescription>
        </GothicShadowFrame>
      </div>
    </div>
  )
}

function WeaponFrame({
  disabled,
  weapon,
}: {
  disabled: boolean
  weapon: Weapon
}) {
  const rarity = 'rare'
  const Icon = WEAPON_ICONS[weapon.weapon_type]
  return (
    <HoverCard>
      <HoverCardTrigger asChild>
        <GothicFrame
          className={cn(
            'relative w-20 overflow-visible',
            weapon.weight === 2 && 'z-10',
            disabled && 'pointer-events-none'
          )}
        >
          <div
            className={weaponWrapper({ rarity: rarity, className: 'h-full' })}
          >
            <div
              className={weaponBody({
                rarity: rarity,
                className: cn(
                  'h-full relative',
                  weapon.weight === 2 ? 'overflow-visible' : 'overflow-hidden'
                ),
              })}
            >
              <Icon
                className={weaponIcon({
                  rarity,
                  weapon_type: weapon.weapon_type,
                  className: 'left-1/2 -translate-x-1/2',
                })}
              />
              {disabled && (
                <div className="absolute inset-0 bg-neutral-300/30" />
              )}
            </div>
          </div>
          <GothicFrame className="pointer-events-none absolute inset-0 z-10 bg-transparent [border-image-slice:22]" />
        </GothicFrame>
      </HoverCardTrigger>
      <GothicHoverCardContent sideOffset={0}>
        <WeaponDetails weapon={weapon} />
      </GothicHoverCardContent>
    </HoverCard>
  )
}

function WeaponFrameExt({
  disabled,
  weapon,
}: {
  disabled: boolean
  weapon: Weapon
}) {
  const rarity = 'rare'
  return (
    <GothicFrame
      className={cn(
        'relative w-20 overflow-visible',
        disabled && 'pointer-events-none'
      )}
    >
      <div className={weaponWrapper({ rarity: rarity, className: 'h-full' })}>
        <div
          className={weaponBody({
            rarity: rarity,
            className: cn(
              'h-full relative',
              weapon.weight === 2 ? 'overflow-visible' : 'overflow-hidden'
            ),
          })}
        >
          {disabled && <div className="absolute inset-0 bg-neutral-300/30" />}
        </div>
      </div>
      <GothicFrame className="pointer-events-none absolute inset-0 z-20 bg-transparent [border-image-slice:22]" />
    </GothicFrame>
  )
}

export { WeaponDetails, WeaponFrame, WeaponFrameExt, weaponIcon }

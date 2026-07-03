import { STAT_LABELS, type Stat } from '#/lib/game/core'
import type { Weapon } from '#/lib/game/weapon'
import { cn, sign } from '#/lib/utils'
import { entries } from '#/utils/maps'
import { cva } from 'class-variance-authority'
import { GothicFrame, GothicShadowFrame } from './gothic-ui/frame'
import { ItemDescription } from './ui/item'
import { Separator } from './ui/separator'
import { statVariants } from './stat-name'

function InlineAuxStats({
  aux_stats,
  className,
  ...props
}: React.ComponentProps<'span'> & {
  aux_stats: Partial<Record<Stat, number>>
}) {
  return (
    <span className={cn('flex gap-2', className)} {...props}>
      {entries(aux_stats).map(([stat, aux], i) => (
        <span
          key={i}
          className={cn({
            'text-positive': aux > 0,
            'text-negative': aux < 0,
          })}
        >
          {sign(aux)}
          {Math.abs(aux)} {STAT_LABELS[stat]}
        </span>
      ))}
    </span>
  )
}

const weaponWrapper = cva('p-px', {
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

function WeaponDetails({ weapon }: { weapon: Weapon }) {
  const rarity = 'rare'
  return (
    <GothicFrame className="font-serif">
      <div className={weaponWrapper({ rarity: rarity })}>
        <div className={weaponBody({ rarity: rarity })}>
          <div className="p-2">
            <div className="absolute z-0 bottom-0 -right-3 -top-4 overflow-hidden">
              <img
                alt="weapon"
                className="right-0 fading-image"
                src="/img/SwordIcon.png"
              />
            </div>
            <div>
              <span className={weaponTitle({ rarity: rarity })}>
                {weapon.name}
              </span>
              <span className="text-foreground/60 block text-xs leading-none">
                Common {weapon.hands}-handed {weapon.weapon_type}
              </span>
            </div>
          </div>
          <div className="px-3">
            <Separator />
          </div>
          <div className="text-foreground/80 italic text-xs px-6 py-2">
            {weapon.description}
          </div>
          <GothicShadowFrame className="z-10 mt-6 m-1 space-y-1">
            <ItemDescription className="text-foreground/80">
              <span className="text-foreground/40 block font-cinzel font-semibold">
                Actions
              </span>
              <span className="pl-4 space-x-2">
                {weapon.actions.map((a) => (
                  <span
                    key={a.ID}
                    className={statVariants({ stat: a.config.stat })}
                  >
                    {a.config.name}
                  </span>
                ))}
              </span>
            </ItemDescription>
            <ItemDescription className="text-foreground/80">
              <span className="text-foreground/40 block font-cinzel font-semibold">
                Effects
              </span>
              {weapon.effects.length > 0 && (
                <span className="pl-4">
                  {weapon.effects.map((e) => e.name).join(', ')}
                </span>
              )}
              <InlineAuxStats className="pl-4" aux_stats={weapon.aux_stats} />
            </ItemDescription>
          </GothicShadowFrame>
        </div>
      </div>
    </GothicFrame>
  )
}

function WeaponFrame({ weapon }: { weapon: Weapon }) {
  const rarity = 'rare'
  return (
    <GothicFrame className='w-20'>
      <div className={weaponWrapper({ rarity: rarity, className: 'h-full' })}>
        <div className={weaponBody({ rarity: rarity, className: 'h-full overflow-hidden' })}>
          <img
            alt="weapon"
            className="absolute top-0 left-1/2 -translate-x-1/2 fading-image"
            src="/img/SwordIcon.png"
          />
        </div>
      </div>
    </GothicFrame>
  )
}

export { WeaponDetails, WeaponFrame }

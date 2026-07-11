import type { ID } from '#/lib/game/core'
import type { Weapon } from '#/lib/game/weapon'
import {
  Combobox,
  ComboboxContent,
  ComboboxEmpty,
  ComboboxInput,
  ComboboxItem,
  ComboboxList,
  ComboboxTrigger,
  ComboboxValue,
} from './ui/combobox'
import { GothicFramedButton } from './gothic-ui/button'
import { HoverCard, HoverCardTrigger } from './ui/hover-card'
import { GothicHoverCardContent } from './gothic-ui/hover-card'
import { WeaponDetails, weaponIcon } from './weapon-details'
import { WEAPON_ICONS } from '#/icons/weapons'

function WeaponCombobox({
  disabled,
  options,
  value,
  other,
  onValueChange,
}: {
  disabled?: boolean
  options: Array<Weapon>
  value: ID | null
  other: ID | null
  onValueChange: (value: ID | null) => void
}) {
  const weapon = options.find((o) => o.ID === value)
  const Icon = weapon ? WEAPON_ICONS[weapon?.weapon_type] : undefined
  return (
    <Combobox
      items={options}
      value={value}
      onValueChange={(v) =>
        v === value ? onValueChange(null) : onValueChange(v)
      }
      disabled={disabled}
    >
      <ComboboxTrigger
        render={
          <GothicFramedButton className="justify-between">
            <ComboboxValue>
              {weapon ? (
                <div className="flex items-center gap-2">
                  <div className="flex size-6 shrink-0 items-center justify-center overflow-hidden">
                    {Icon && (
                      <Icon
                        className={weaponIcon({
                          rarity: 'common',
                          weapon_type: weapon.weapon_type,
                          className: '!static max-h-5 max-w-5',
                        })}
                      />
                    )}
                  </div>
                  <div className="truncate">{weapon.name}</div>
                </div>
              ) : (
                <span className="text-foreground/60">Select Weapon</span>
              )}
            </ComboboxValue>
          </GothicFramedButton>
        }
      />
      <ComboboxContent>
        <ComboboxInput showTrigger={false} placeholder="Search" />
        <ComboboxEmpty>No weapons found.</ComboboxEmpty>
        <ComboboxList>
          {(weapon: Weapon) => (
            <HoverCard key={weapon.ID}>
              <HoverCardTrigger asChild>
                <ComboboxItem
                  value={weapon.ID}
                  disabled={weapon.hands > 1 && !!other}
                >
                  {weapon.name}
                </ComboboxItem>
              </HoverCardTrigger>
              <GothicHoverCardContent sideOffset={0} side="left">
                <WeaponDetails weapon={weapon} />
              </GothicHoverCardContent>
            </HoverCard>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { WeaponCombobox }

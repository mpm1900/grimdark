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
import { WeaponDetails } from './weapon-details'

function WeaponCombobox({
  remaining_weight,
  disabled,
  options,
  value,
  onValueChange,
}: {
  remaining_weight: number
  disabled?: boolean
  options: Array<Weapon>
  value: ID | null
  onValueChange: (value: ID | null) => void
}) {
  const weapon = options.find((o) => o.ID === value)
  const available_weight = remaining_weight + (weapon?.weight ?? 0)
  return (
    <Combobox
      items={options}
      value={value}
      onValueChange={(v) =>
        v === value ? onValueChange(null) : onValueChange(v)
      }
      disabled={disabled || (remaining_weight <= 0 && !weapon)}
    >
      <HoverCard open={!weapon ? false : undefined}>
        <HoverCardTrigger asChild>
          <ComboboxTrigger
            render={
              <GothicFramedButton className="justify-between">
                <ComboboxValue>
                  {weapon ? (
                    <div className="flex items-center gap-2 truncate">
                      <div className="truncate">{weapon.name}</div>
                    </div>
                  ) : (
                    <span className="text-foreground/60">Select Weapon</span>
                  )}
                </ComboboxValue>
              </GothicFramedButton>
            }
          />
        </HoverCardTrigger>
        <GothicHoverCardContent sideOffset={0} side="left" className="w-80">
          {weapon && <WeaponDetails weapon={weapon} />}
        </GothicHoverCardContent>
      </HoverCard>
      <ComboboxContent>
        <ComboboxInput showTrigger={false} placeholder="Search" />
        <ComboboxEmpty>No weapons found.</ComboboxEmpty>
        <ComboboxList>
          {(w: Weapon) => (
            <HoverCard key={w.ID}>
              <HoverCardTrigger asChild>
                <ComboboxItem
                  value={w.ID}
                  disabled={w.weight > available_weight && w.ID !== weapon?.ID}
                >
                  {w.name}
                </ComboboxItem>
              </HoverCardTrigger>
              <GothicHoverCardContent
                sideOffset={0}
                side="left"
                className="w-80"
              >
                <WeaponDetails weapon={w} />
              </GothicHoverCardContent>
            </HoverCard>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { WeaponCombobox }

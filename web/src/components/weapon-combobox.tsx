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
                <div className="truncate">{weapon.name}</div>
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
            <ComboboxItem
              key={weapon.ID}
              value={weapon.ID}
              disabled={weapon.hands > 1 && !!other}
            >
              {weapon.name}
            </ComboboxItem>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { WeaponCombobox }

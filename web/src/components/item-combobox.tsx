import type { ActorClass } from '#/lib/game/actor-class'
import type { ID } from '#/lib/game/core'
import { itemsQuery } from '#/lib/queries/get-items'
import { useQuery } from '@tanstack/react-query'
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
import type { Item } from '#/lib/game/weapon'
import { ItemTooltip } from './item-tooltip'

function ItemCombobox({
  actor_class,
  value,
  onValueChange,
}: {
  actor_class?: ActorClass
  value: ID | null
  onValueChange: (value: ID | null) => void
}) {
  const class_options = actor_class?.options.items ?? []
  const query = useQuery(itemsQuery)
  const options = [...class_options, ...(query.data ?? [])]
  const item = options.find((o) => o.ID === value)

  return (
    <Combobox
      items={options}
      value={value}
      onValueChange={(v) =>
        v === value ? onValueChange(null) : onValueChange(v)
      }
    >
      <ItemTooltip
        item={item}
        hover_card={{ open: !item ? false : undefined }}
        asChild
      >
        <ComboboxTrigger
          render={
            <GothicFramedButton className="justify-between">
              <ComboboxValue>
                {item ? (
                  <div className="flex items-center gap-2 truncate">
                    <div className="truncate">{item.name}</div>
                  </div>
                ) : (
                  <span className="text-foreground/60">Select Item</span>
                )}
              </ComboboxValue>
            </GothicFramedButton>
          }
        />
      </ItemTooltip>
      <ComboboxContent>
        <ComboboxInput showTrigger={false} placeholder="Search" />
        <ComboboxEmpty>No items found.</ComboboxEmpty>
        <ComboboxList>
          {(item: Item) => (
            <ItemTooltip key={item.ID} item={item}>
              <ComboboxItem value={item.ID}>{item.name}</ComboboxItem>
            </ItemTooltip>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { ItemCombobox }

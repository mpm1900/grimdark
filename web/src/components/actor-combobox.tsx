import type { ID } from '#/lib/game/core'
import { actorsQuery } from '#/lib/queries/get-actors'
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
import { Button } from './ui/button'
import type { ActorClass } from '#/lib/game/actor-class'
import { ChevronsUpDown } from 'lucide-react'

function ActorCombobox({
  value,
  onValueChange,
}: {
  value: ID | null
  onValueChange: (value: ID) => void
}) {
  const query = useQuery(actorsQuery)

  return (
    <Combobox
      items={query.data}
      value={value}
      onValueChange={(v) => {
        if (!v) {
          return
        }

        onValueChange(v)
      }}
    >
      <ComboboxTrigger
        render={
          <Button variant="outline" className="justify-between">
            <ComboboxValue>
              {(value) =>
                query.data?.find((a) => a.ID === value)?.name ?? (
                  <span className="text-foreground/60">Select a Class</span>
                )
              }
            </ComboboxValue>
            <ChevronsUpDown />
          </Button>
        }
      />
      <ComboboxContent>
        <ComboboxInput showTrigger={false} placeholder="Search" />
        <ComboboxEmpty>No classes found.</ComboboxEmpty>
        <ComboboxList>
          {(item: ActorClass) => (
            <ComboboxItem key={item.ID} value={item.ID}>
              {item.name}
            </ComboboxItem>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { ActorCombobox }

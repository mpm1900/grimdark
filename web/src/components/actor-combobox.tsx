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
import type { ActorClass } from '#/lib/game/actor-class'
import { GothicBigButton } from './gothic-ui/button'
import { ClassTooltip } from './class-tooltip'

function ActorCombobox({
  value,
  onValueChange,
}: {
  value: ID | null
  onValueChange: (value: ID) => void
}) {
  const query = useQuery(actorsQuery)
  const actor_class = query.data?.find((a) => a.ID === value)

  return (
    <Combobox
      items={query.data ?? []}
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
          <ClassTooltip
            actor_class={actor_class}
            hover_card={{ open: !actor_class ? false : undefined }}
            asChild
          >
            <GothicBigButton>
              <ComboboxValue>
                {(value) =>
                  query.data?.find((a) => a.ID === value)?.name ?? (
                    <span className="text-foreground/60">Select a Class</span>
                  )
                }
              </ComboboxValue>
            </GothicBigButton>
          </ClassTooltip>
        }
      />
      <ComboboxContent>
        <ComboboxInput showTrigger={false} placeholder="Search" />
        <ComboboxEmpty>No classes found.</ComboboxEmpty>
        <ComboboxList>
          {(item: ActorClass) => (
            <ClassTooltip key={item.ID} actor_class={item} asChild>
              <ComboboxItem value={item.ID}>{item.name}</ComboboxItem>
            </ClassTooltip>
          )}
        </ComboboxList>
      </ComboboxContent>
    </Combobox>
  )
}

export { ActorCombobox }

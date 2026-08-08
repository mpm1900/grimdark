import type { Item } from '#/lib/game/weapon'
import { EffectTooltip } from './effect-tooltip'
import { ItemDescription } from './ui/item'

function ItemDetails({ item }: { item: Item }) {
  return (
    <div className="p-1 font-serif">
      <div>{item.name}</div>
      <div className="text-foreground/80 italic text-xs px-6 py-2">
        {item.description}
      </div>
      <ItemDescription className="text-foreground/80">
        <span className="text-foreground/40 block font-cinzel font-semibold">
          Effects
        </span>
        {item.effects.length > 0 && (
          <span className="space-x-2 flex flex-wrap">
            {item.effects.map((e, i) => (
              <EffectTooltip
                key={`${e.ID}-${e.name}-${i}`}
                effect={e}
                className="cursor-default hover:underline capitalize"
              >
                {e.name}
              </EffectTooltip>
            ))}
          </span>
        )}
      </ItemDescription>
    </div>
  )
}

export { ItemDetails }

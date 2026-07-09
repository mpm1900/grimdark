import { gameStore } from '#/lib/stores/game'
import { useSelector } from '@tanstack/react-store'
import { ActiveContext } from './active-context'

function TurnContext() {
  const turn = useSelector(gameStore, (g) => g.turn)
  const active_context = useSelector(gameStore, (g) => g.active_context)
  return (
    <ActiveContext asChild active_context={active_context}>
      <div className="absolute z-30 -top-4 left-1/2 -translate-x-1/2 bg-[url(/gothic/TitleHeroFrame.png)] bg-cover bg-no-repeat h-24 leading-18 w-48 text-xl text-center font-cinzel-dec font-bold text-foreground/60">
        <div>
          T<span className="font-cinzel">urn {turn}</span>
        </div>
      </div>
    </ActiveContext>
  )
}

export { TurnContext }

import { RenderLog } from '#/lib/game/log'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
import { useSelector } from '@tanstack/react-store'
import { useEffect, useRef } from 'react'

function BattleLog({ className, ...props }: React.ComponentProps<'div'>) {
  const logs = useSelector(gameStore, (g) => g.logs)
  const endRef = useRef<HTMLLIElement | null>(null)
  const lastLogID = logs.length > 0 ? logs[logs.length - 1].ID : null

  useEffect(() => {
    endRef.current?.scrollIntoView({ block: 'end' })
  }, [lastLogID])

  return (
    <div
      className={cn(
        'flex-1 overflow-auto font-serif text-foreground/40',
        className
      )}
      {...props}
    >
      <ul className="px-1">
        {logs.map((log, index) => (
          <li
            key={log.ID}
            className="leading-4.5 text-sm truncate"
            ref={index === logs.length - 1 ? endRef : undefined}
          >
            {RenderLog(log)}
          </li>
        ))}
      </ul>
    </div>
  )
}

export { BattleLog }

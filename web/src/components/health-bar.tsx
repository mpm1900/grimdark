import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { GothicProgress } from './gothic-ui/progress'
import { Progress } from './ui/progress'

type HealthBarType = 'value' | 'ratio'

function HealthBar({
  actor,
  type,
  hide_numbers,
  className,
  ...props
}: React.ComponentProps<typeof Progress> & {
  actor: Actor
  hide_numbers?: boolean
  type: HealthBarType
}) {
  const health = actor.stats['health']
  const remaining = health - actor.wounds
  const ratio = remaining / health
  return (
    <GothicProgress
      {...props}
      className={cn('h-6 font-cinzel font-semibold rounded-xs', className)}
      value={ratio * 100}
    >
      {!hide_numbers && (
        <div className="absolute inset-0 flex items-center justify-end pr-1 text-foreground/70 [text-shadow:1px_1px_0_var(--color-black)] text-sm">
          {type === 'value' && (
            <span>
              {remaining}/{health}
            </span>
          )}
          {type === 'ratio' && (
            <span>{((remaining * 100) / health).toFixed(0)}% HP</span>
          )}
        </div>
      )}
    </GothicProgress>
  )
}

export { HealthBar }

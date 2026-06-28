import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { Progress } from './ui/progress'

type HealthBarType = 'value' | 'ratio'

function HealthBar({
  actor,
  type,
  ...props
}: React.ComponentProps<typeof Progress> & {
  actor: Actor
  type: HealthBarType
}) {
  const health = actor.stats['health']
  const remaining = health - actor.wounds
  const ratio = remaining / health
  return (
    <Progress
      {...props}
      className="h-5 rounded-sm"
      indicator={{
        className: cn('bg-red-400'),
      }}
      value={ratio * 100}
    >
      <div className="absolute inset-y-0 right-2 font-bold [text-shadow:1px_1px_0_var(--color-black)] text-sm">
        {type === 'value' && (
          <span>
            {remaining}/{health} HP
          </span>
        )}
        {type === 'ratio' && (
          <span>{((remaining * 100) / health).toFixed(0)}% HP</span>
        )}
      </div>
    </Progress>
  )
}

export { HealthBar }

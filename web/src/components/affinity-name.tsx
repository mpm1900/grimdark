import { AFFINITY_ICONS } from '#/icons/affinity'
import type { Affinity } from '#/lib/game/core'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'

const affinityVariants = cva('capitalize', {
  variants: {
    affinity: {
      arcane: 'text-arcane',
      cryo: 'text-cryo',
      fire: 'text-fire',
      kinetic: 'text-kinetic',
      lightning: 'text-lightning',
      poison: 'text-poison',
      psychic: 'text-psychic',
    },
  },
  defaultVariants: {
    affinity: 'kinetic',
  },
})

function AffinityName({
  affinity,
  className,
  ...props
}: React.ComponentProps<'span'> & {
  affinity: Affinity
}) {
  const Icon = AFFINITY_ICONS[affinity]
  return (
    <span {...props} className={cn(affinityVariants({ affinity }), className)}>
      <Icon className="size-5" />
    </span>
  )
}

export { AffinityName, affinityVariants }

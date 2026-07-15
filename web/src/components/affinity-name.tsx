import { AFFINITY_ICONS } from '#/icons/affinity'
import type { Affinity } from '#/lib/game/core'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'
import type { IconType } from 'react-icons/lib'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'

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

function AffinityIcon({
  affinity,
  className,
  ...props
}: React.ComponentProps<IconType> & {
  affinity: Affinity
}) {
  const Icon = AFFINITY_ICONS[affinity]
  if (!Icon) return null
  return (
    <Tooltip delayDuration={1000}>
      <TooltipTrigger asChild>
        <Icon
          className={cn(affinityVariants({ affinity }), className)}
          {...props}
        />
      </TooltipTrigger>
      <TooltipContent>{affinity}</TooltipContent>
    </Tooltip>
  )
}

export { AffinityIcon, affinityVariants }

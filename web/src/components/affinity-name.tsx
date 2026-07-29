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
      blood: 'text-blood',
      cryo: 'text-cryo',
      holy: 'text-holy',
      fire: 'text-fire',
      physical: 'text-physical',
      lightning: 'text-lightning',
      poison: 'text-poison',
      psychic: 'text-psychic',
    },
  },
  defaultVariants: {
    affinity: 'physical',
  },
})

function AffinityIcon({
  affinity,
  children,
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
        {children ? (
          <span className={cn(affinityVariants({ affinity }), className)}>
            {children}
          </span>
        ) : (
          <Icon
            className={cn(affinityVariants({ affinity }), className)}
            {...props}
          />
        )}
      </TooltipTrigger>
      <TooltipContent className="capitalize font-serif font-semibold text-sm">
        {affinity}
      </TooltipContent>
    </Tooltip>
  )
}

export { AffinityIcon, affinityVariants }

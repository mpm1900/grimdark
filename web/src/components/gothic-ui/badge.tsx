import * as React from 'react'
import { cva, type VariantProps } from 'class-variance-authority'
import { Slot } from 'radix-ui'

import { cn } from '#/lib/utils.ts'

const gothicBadgeVariants = cva(
  cn(
    'inline-flex w-fit shrink-0 items-center justify-center gap-1 border-[3px] border-solid border-transparent bg-transparent px-1 py-0.5 transition-colors [image-rendering:pixelated] [&>svg]:pointer-events-none [&>svg]:size-3',
    'text-xs font-semibold font-cinzel whitespace-nowrap leading-none text-foreground/80 [text-shadow:0_1px_0_var(--color-black)]'
  ),
  {
    variants: {
      variant: {
        default:
          '[border-image-slice:8_fill] [border-image-repeat:stretch] [border-image-source:url("/gothic/ButtonLittleGray_Normal.png")]',
        disabled:
          'text-foreground/50 [border-image-slice:8_fill] [border-image-repeat:stretch] [border-image-source:url("/gothic/ButtonLittleGray_Disabled.png")]',
        light:
          '[border-image-slice:8_fill] [border-image-repeat:stretch] [border-image-source:url("/gothic/ButtonLittleGray_Hovered.png")]',
        empty:
          '[border-image-slice:8_fill] [border-image-repeat:stretch] [border-image-source:url("/gothic/ButtonLittleGray_Empty.png")]',
        'empty-strong':
          '[border-image-slice:8_fill] [border-image-repeat:stretch] [border-image-source:url("/gothic/ButtonLittleGray_EmptyStrong.png")]',
      },
    },
    defaultVariants: {
      variant: 'default',
    },
  }
)

function GothicBadge({
  className,
  variant = 'default',
  asChild = false,
  ...props
}: React.ComponentProps<'span'> &
  VariantProps<typeof gothicBadgeVariants> & { asChild?: boolean }) {
  const Comp = asChild ? Slot.Root : 'span'

  return (
    <Comp
      data-slot="badge"
      data-variant={variant}
      className={cn(gothicBadgeVariants({ variant }), className)}
      {...props}
    />
  )
}

const tinyBadgeVariants = cva(
  cn(
    'text-[10px] py-0 px-0.5 h-3 leading-3 border rounded-sm font-serif font-bold flex inline-block truncate'
  ),
  {
    variants: {
      variant: {
        default: 'bg-neutral-950 text-foreground border',
        positive: 'bg-neutral-950 text-lime-200 border-lime-900/50',
        negative: 'bg-neutral-950 text-red-200 border-red-900/50',

        // statuses
        bleeding: 'bg-red-950 text-red-300 border-red-800/50',
        burned: 'bg-orange-900 text-orange-300 border-orange-950',
        frozen: 'bg-cyan-950 text-cyan-300 border-cyan-800/50',
        poisoned: 'bg-green-950 text-green-200 border-green-900/50',
        shocked: 'bg-yellow-950 text-blue-200 border-yellow-600/50',
        sleeping: 'bg-neutral-950 text-purple-300 border-purple-800/50',
      },
    },
    defaultVariants: {
      variant: 'default',
    },
  }
)

function TinyBadge({
  className,
  variant,
  asChild = false,
  ...props
}: React.ComponentProps<'span'> &
  VariantProps<typeof tinyBadgeVariants> & { asChild?: boolean }) {
  const Comp = asChild ? Slot.Root : 'span'

  return (
    <Comp
      data-slot="tiny-badge"
      className={cn(tinyBadgeVariants({ variant }), className)}
      {...props}
    />
  )
}

export { GothicBadge, gothicBadgeVariants, TinyBadge }

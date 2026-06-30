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
    'text-[10px] px-0.5 leading-0 py-1.5 border rounded-sm font-serif font-bold'
  ),
  {
    variants: {
      variant: {
        default: 'bg-neutral-950 text-foreground border',
        positive: 'bg-neutral-950 text-emerald-300 border-emerald-900/50',
        negative: 'bg-neutral-950 text-red-300 border-red-900/50',
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

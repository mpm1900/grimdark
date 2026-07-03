import * as React from 'react'

import { cn } from '#/lib/utils'
import { Slot } from 'radix-ui'
import { cva, type VariantProps } from 'class-variance-authority'

const gothicButtonVariants = cva(
  cn(
    'font-cinzel [text-shadow:0_1px_0_var(--color-black)] text-sm font-bold select-none whitespace-nowrap',
    'inline-flex shrink-0 items-center justify-center gap-2 outline-none ring-1 ring-black'
  ),
  {
    variants: {
      variant: {
        basic: cn(
          'px-3 py-1 border-[6px] border-solid border-transparent bg-transparent text-foreground/80',
          '[border-image-slice:16_fill] [border-image-repeat:repeat-infinite]',
          "[border-image-source:url('/gothic/ButtonFramedGray_Normal.png')]",
          "active:[border-image-source:url('/gothic/ButtonFramedGray_Pressed.png')] active:text-foreground/60",
          "disabled:[border-image-source:url('/gothic/ButtonFramedGray_Disabled.png')] disabled:pointer-events-none disabled:text-foreground/60",
          "hover:[border-image-source:url('/gothic/ButtonFramedGray_Hovered.png')] hover:text-foreground"
        ),
        red: cn(
          'px-3 py-1 border-[4px] border-solid border-transparent bg-transparent text-foreground/80',
          '[border-image-slice:18_fill] [border-image-repeat:repeat-infinite]',
          "[border-image-source:url('/gothic/ButtonFramedRed_Normal.png')]",
          "active:[border-image-source:url('/gothic/ButtonFramedRed_Pressed.png')] active:text-foreground/60",
          "disabled:[border-image-source:url('/gothic/ButtonFramedRed_Disabled.png')] disabled:pointer-events-none disabled:text-foreground/60",
          "hover:[border-image-source:url('/gothic/ButtonFramedRed_Hovered.png')] hover:text-foreground"
        ),
      },
    },
    defaultVariants: {
      variant: 'basic',
    },
  }
)

function GothicFramedButton({
  className,
  variant,
  asChild = false,
  ...props
}: React.ComponentProps<'button'> &
  VariantProps<typeof gothicButtonVariants> & {
    asChild?: boolean
  }) {
  const Comp = asChild ? Slot.Root : 'button'
  return (
    <Comp
      data-slot="gothic-button"
      data-variant={variant}
      className={cn(gothicButtonVariants({ variant }), className)}
      {...props}
    />
  )
}

const gothicBigButtonVariants = cva(
  cn(
    'font-cinzel [text-shadow:0_1px_0_var(--color-black)] text-sm font-bold select-none whitespace-nowrap',
    'inline-flex shrink-0 items-center justify-center gap-2 outline-none'
  ),
  {
    variants: {
      variant: {
        basic: cn(
          'min-h-12 px-8 py-2 border-x-[16px] border-y-[6px] border-solid border-transparent bg-transparent text-foreground/80',
          '[border-image-slice:12_40_12_40_fill] [border-image-repeat:stretch]',
          "[border-image-source:url('/gothic/ButtonStandartGray_Normal.png')]",
          "active:[border-image-source:url('/gothic/ButtonStandartGray_Pressed.png')] active:text-foreground/60",
          "disabled:[border-image-source:url('/gothic/ButtonStandartGray_Disabled.png')] disabled:pointer-events-none disabled:text-foreground/40",
          "hover:[border-image-source:url('/gothic/ButtonStandartGray_Hovered.png')] hover:text-foreground"
        ),
        red: cn(
          'min-h-12 px-8 py-2 border-x-[16px] border-y-[6px] border-solid border-transparent bg-transparent text-foreground/80',
          '[border-image-slice:12_40_12_40_fill] [border-image-repeat:stretch]',
          "[border-image-source:url('/gothic/ButtonStandartRed_Normal.png')]",
          "active:[border-image-source:url('/gothic/ButtonStandartRed_Pressed.png')] active:text-foreground/60",
          "disabled:[border-image-source:url('/gothic/ButtonStandartRed_Disabled.png')] disabled:pointer-events-none disabled:text-foreground/60",
          "hover:[border-image-source:url('/gothic/ButtonStandartRed_Hovered.png')] hover:text-foreground"
        ),
        green: cn(
          'min-h-12 px-8 py-2 border-x-[16px] border-y-[6px] border-solid border-transparent bg-transparent text-foreground/80',
          '[border-image-slice:12_40_12_40_fill] [border-image-repeat:stretch]',
          "[border-image-source:url('/gothic/ButtonStandartGreen_Normal.png')]",
          "active:[border-image-source:url('/gothic/ButtonStandartGreen_Pressed.png')] active:text-foreground/60",
          "disabled:[border-image-source:url('/gothic/ButtonStandartGreen_Disabled.png')] disabled:pointer-events-none disabled:text-foreground/60",
          "hover:[border-image-source:url('/gothic/ButtonStandartGreen_Hovered.png')] hover:text-foreground"
        ),
      },
    },
    defaultVariants: {
      variant: 'basic',
    },
  }
)

function GothicBigButton({
  className,
  variant,
  asChild = false,
  ...props
}: React.ComponentProps<'button'> &
  VariantProps<typeof gothicBigButtonVariants> & {
    asChild?: boolean
  }) {
  const Comp = asChild ? Slot.Root : 'button'
  return (
    <Comp
      data-slot="gothic-big-button"
      data-variant={variant}
      className={cn(gothicBigButtonVariants({ variant }), className)}
      {...props}
    />
  )
}

function GothicMiniButton({
  className,
  asChild = false,
  ...props
}: React.ComponentProps<'button'> & {
  asChild?: boolean
}) {
  const Comp = asChild ? Slot.Root : 'button'
  return (
    <Comp
      data-slot="button"
      className={cn(
        'inline-flex w-fit shrink-0 items-center justify-center gap-1 border-[3px] border-solid border-transparent bg-transparent px-1 py-0.5 transition-colors [image-rendering:pixelated] [&>svg]:pointer-events-none [&>svg]:size-3',
        'text-xs font-semibold font-cinzel whitespace-nowrap leading-none text-foreground/80 [text-shadow:0_1px_0_var(--color-black)]',
        '[border-image-slice:8_fill] [border-image-repeat:stretch] [border-image-source:url("/gothic/ButtonLittleGray_Normal.png")] text-foreground/80',
        'hover:[border-image-slice:8_fill] hover:[border-image-repeat:stretch] hover:[border-image-source:url("/gothic/ButtonLittleGray_Hovered.png")] hover:text-foreground',
        'active:[border-image-slice:8_fill] active:[border-image-repeat:stretch] active:[border-image-source:url("/gothic/ButtonLittleGray_Pressed.png")] active:text-foreground/70',
        className
      )}
      {...props}
    />
  )
}

function GothicCloseButton({
  className,
  asChild = false,
  ...props
}: React.ComponentProps<'button'> & {
  asChild?: boolean
}) {
  const Comp = asChild ? Slot.Root : 'button'
  return (
    <Comp
      data-slot="button"
      className={cn(
        'm-0 inline-flex size-6 shrink-0 items-center justify-center border-0 bg-transparent p-0 outline-none',
        'bg-contain bg-center bg-no-repeat',
        "bg-[url('/gothic/SquareRedButtonExit_Normal.png')]",
        "hover:bg-[url('/gothic/SquareRedButtonExit_Hovered.png')]",
        "active:bg-[url('/gothic/SquareRedButtonExit_Pressed.png')]",
        "disabled:bg-[url('/gothic/SquareRedButtonExit_Disabled.png')] disabled:pointer-events-none",
        className
      )}
      {...props}
    />
  )
}

export { GothicFramedButton, GothicBigButton, GothicMiniButton, GothicCloseButton }

import * as React from 'react'
import { Progress as ProgressPrimitive } from 'radix-ui'

import { cn } from '#/lib/utils.ts'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'

type GothicProgressProps = React.ComponentProps<
  typeof ProgressPrimitive.Root
> & {
  indicator?: Partial<React.ComponentProps<typeof ProgressPrimitive.Indicator>>
  resetKey?: React.Key
  track?: React.ComponentProps<'div'>
  frame?: React.ComponentProps<'div'>
}

function GothicProgress({
  children,
  className,
  frame: { className: frameClassName, ...frame } = {},
  indicator: {
    className: indicatorClassName,
    style: indicatorStyle,
    ...indicator
  } = {},
  track: { className: trackClassName, ...track } = {},
  resetKey,
  value,
  ...props
}: GothicProgressProps) {
  const turn = useSelector(gameStore, (g) => g.turn)
  const clampedValue = Math.min(Math.max(value ?? 0, 0), 100)
  const delta = React.useRef({ resetKey, turn, value: clampedValue })

  if (delta.current.turn !== turn || delta.current.resetKey !== resetKey) {
    delta.current = { resetKey, turn, value: clampedValue }
  }

  return (
    <ProgressPrimitive.Root
      data-slot="progress"
      className={cn(
        'relative w-full overflow-hidden bg-neutral-950 border-2 border-foreground/30 ring ring-black',
        className
      )}
      value={value}
      {...props}
    >
      <div
        data-slot="progress-track"
        className={cn(
          'absolute inset-x-0 inset-y-0 overflow-hidden bg-neutral-950 my-px',
          trackClassName
        )}
        {...track}
      >
        <ProgressPrimitive.Indicator
          data-slot="progress-delta-indicator"
          className={cn(
            'absolute inset-0 h-full w-full bg-white/30 transition-[clip-path] duration-300 ease-out',
            'm-px'
          )}
          style={{
            clipPath: `inset(0 ${100 - delta.current.value}% 0 0)`,
          }}
        />
        <ProgressPrimitive.Indicator
          data-slot="progress-indicator"
          className={cn(
            'absolute inset-0 h-full w-full bg-[url("/gothic/HealthBarMini_Fill.png")] bg-[length:100%_100%] bg-center bg-no-repeat transition-[clip-path] duration-300 ease-out',
            'm-px',
            indicatorClassName
          )}
          style={{
            clipPath: `inset(0 ${100 - clampedValue}% 0 0)`,
            ...indicatorStyle,
          }}
          {...indicator}
        />
        <div
          aria-hidden
          data-slot="progress-glass"
          className={cn(
            'pointer-events-none absolute inset-0 left-px right-px',
            '[border-image-source:url("/gothic/HealthBarMini_Glass.png")]',
            '[border-image-slice:15_fill]',
            '[border-image-repeat:stretch]',
            frameClassName
          )}
          {...frame}
        />
      </div>
      {children}
    </ProgressPrimitive.Root>
  )
}

export { GothicProgress }

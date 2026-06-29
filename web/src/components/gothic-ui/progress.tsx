import * as React from 'react'
import { Progress as ProgressPrimitive } from 'radix-ui'

import { cn } from '#/lib/utils.ts'

type GothicProgressProps = React.ComponentProps<
  typeof ProgressPrimitive.Root
> & {
  indicator?: Partial<React.ComponentProps<typeof ProgressPrimitive.Indicator>>
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
  value,
  ...props
}: GothicProgressProps) {
  const clampedValue = Math.min(Math.max(value ?? 0, 0), 100)

  return (
    <ProgressPrimitive.Root
      data-slot="progress"
      className={cn('relative w-full overflow-hidden bg-black', className)}
      value={value}
      {...props}
    >
      <div
        data-slot="progress-track"
        className={cn(
          'absolute inset-x-[2%] inset-y-[18%] overflow-hidden',
          trackClassName
        )}
        {...track}
      >
        <ProgressPrimitive.Indicator
          data-slot="progress-indicator"
          className={cn(
            'h-full w-full bg-[url("/gothic/HealthBarMini_Fill_Middle.png")] bg-[length:100%_100%] bg-center bg-no-repeat transition-[clip-path] duration-300 ease-out',
            indicatorClassName
          )}
          style={{
            clipPath: `inset(0 ${100 - clampedValue}% 0 0)`,
            ...indicatorStyle,
          }}
          {...indicator}
        />
      </div>
      <div
        aria-hidden
        data-slot="progress-frame"
        className={cn(
          'pointer-events-none absolute inset-0 border-[3px] border-solid border-transparent',
          '[border-image-source:url("/gothic/HealthBarMini_frame.png")]',
          '[border-image-slice:15_fill]',
          '[border-image-repeat:stretch]',
          frameClassName
        )}
        {...frame}
      />
      {children}
    </ProgressPrimitive.Root>
  )
}

export { GothicProgress }

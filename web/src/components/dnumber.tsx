import { cn } from '#/lib/utils'

function DNumber({
  value,
  r = 0,
  perfect,
  className,
  ...props
}: React.ComponentProps<'span'> & {
  value: number
  r?: number
  perfect?: boolean
}) {
  return (
    <span
      className={cn(
        {
          'text-positive': value > r,
          'text-negative': value < r,
          'text-perfect': perfect,
        },
        className
      )}
      {...props}
    />
  )
}

export { DNumber }

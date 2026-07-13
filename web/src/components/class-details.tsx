import { cn } from '#/lib/utils'

function ClassDetails({ className, ...props }: React.ComponentProps<'div'>) {
  return <div className={cn(className)} {...props}></div>
}

export { ClassDetails }

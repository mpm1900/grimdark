import type { ErrorRouteComponent } from '@tanstack/react-router'
import { GothicCard } from './gothic-ui/card'

const PageError: ErrorRouteComponent = (props) => (
  <div className="h-full w-full grid place-items-center">
    <GothicCard className="px-8 py-4 font-serif text-foreground">
      <div>{String(props.error)}</div>
      <div>{props.info?.componentStack}</div>
    </GothicCard>
  </div>
)

export { PageError }

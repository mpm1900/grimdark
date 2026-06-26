import { AppHeader } from '#/components/app-header'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({ component: Home })

function Home() {
  return (
    <div>
      <ClientOnly>
        <AppHeader />
      </ClientOnly>
    </div>
  )
}

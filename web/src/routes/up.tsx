import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/up')({
  component: Up,
})

function Up() {
  return 'ok'
}

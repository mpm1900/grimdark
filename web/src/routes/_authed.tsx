import {
  Navigate,
  Outlet,
  createFileRoute,
  redirect,
} from '@tanstack/react-router'

export const Route = createFileRoute('/_authed')({
  beforeLoad: ({ context }) => {
    if (!context.auth.user) {
      throw redirect({ to: '/login' })
    }
  },
  component: AuthenticatedLayout,
  notFoundComponent: () => <Navigate to="/" replace />,
})

function AuthenticatedLayout() {
  return <Outlet />
}

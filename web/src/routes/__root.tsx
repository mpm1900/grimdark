import { HeadContent, Scripts, createRootRouteWithContext } from '@tanstack/react-router'

import app_css from '../styles.css?url'
import type { QueryClient } from '@tanstack/react-query'
import { userQuery, type User } from '#/lib/queries/auth'
import { Login } from './login'
import { Toaster } from '#/components/ui/sonner'
import { TooltipProvider } from '#/components/ui/tooltip'

interface RouterContext {
  queryClient: QueryClient
  auth: {
    user: User | null
  }
}

export const Route = createRootRouteWithContext<RouterContext>()({
  beforeLoad: async ({ context, location }) => {
    if (location.pathname === '/up') {
      return { auth: { user: null } } as any
    }

    const user = await context.queryClient.fetchQuery(userQuery)
    return {
      auth: {
        user,
      },
    }
  },
  head: () => ({
    meta: [
      {
        charSet: 'utf-8',
      },
      {
        name: 'viewport',
        content: 'width=device-width, initial-scale=1',
      },
      {
        title: 'GrimDark',
      },
    ],
    links: [
      {
        rel: 'stylesheet',
        href: app_css,
      },
    ],
  }),
  shellComponent: RootDocument,
  errorComponent: Login,
})

function RootDocument({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" className='dark'>
      <head>
        <HeadContent />
      </head>
      <body>
        <TooltipProvider>{children}</TooltipProvider>
        <Toaster position="bottom-center" />
        <Scripts />
      </body>
    </html>
  )
}

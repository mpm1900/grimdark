import {
  HeadContent,
  Scripts,
  createRootRouteWithContext,
} from '@tanstack/react-router'

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

const gothicPreloadImages = [
  ...['Gray', 'Red'].flatMap((color) =>
    ['Normal', 'Hovered', 'Pressed', 'Disabled'].map(
      (state) => `/gothic/ButtonFramed${color}_${state}.png`
    )
  ),
  ...['Gray', 'Green', 'Red'].flatMap((color) =>
    ['Normal', 'Hovered', 'Pressed', 'Disabled'].map(
      (state) => `/gothic/ButtonStandart${color}_${state}.png`
    )
  ),
  ...['Normal', 'Hovered', 'Pressed', 'Disabled'].map(
    (state) => `/gothic/SquareRedButtonExit_${state}.png`
  ),
]

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
      ...gothicPreloadImages.map((href) => ({
        rel: 'preload',
        as: 'image',
        type: 'image/png',
        href,
      })),
    ],
  }),
  shellComponent: RootDocument,
  errorComponent: Login,
})

function RootDocument({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" className="dark">
      <head>
        <HeadContent />
      </head>
      <body>
        <GothicImagePreloader />
        <TooltipProvider>{children}</TooltipProvider>
        <Toaster position="bottom-center" />
        <Scripts />
      </body>
    </html>
  )
}

function GothicImagePreloader() {
  return (
    <div
      aria-hidden
      className="pointer-events-none fixed size-0 overflow-hidden opacity-0"
    >
      {gothicPreloadImages.map((src) => (
        <img key={src} src={src} alt="" width={1} height={1} />
      ))}
    </div>
  )
}

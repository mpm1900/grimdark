import { GothicCard, GothicCardContent } from '#/components/gothic-ui/card'
import { GothicTabsList, GothicTabsTrigger } from '#/components/gothic-ui/tabs'
import { LoginForm } from '#/components/login/login-form'
import { SignupForm } from '#/components/login/signup-form'
import { PanelHeader } from '#/components/panels/panel-header'
import { Tabs, TabsContent } from '#/components/ui/tabs'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { GiWingedSword } from 'react-icons/gi'

export const Route = createFileRoute('/login')({
  beforeLoad: ({ context }) => {
    if (context.auth.user) {
      throw redirect({ to: '/' })
    }
  },
  component: Login,
})

function Login() {
  return (
    <div className="flex min-h-svh w-full items-start justify-center overflow-y-auto px-4 py-8 sm:px-6 md:p-10 [@media(min-height:760px)]:items-center">
      <div className="w-full max-w-sm space-y-2">
        <div className="relative text-foreground/60 flex mb-8 items-center justify-center rounded-md z-20 md:mb-12">
          <GiWingedSword className="size-12 md:size-16" />
        </div>

        <Tabs defaultValue="login" className="flex flex-col gap-6">
          <GothicCard>
            <PanelHeader className="mt-0 pt-2 grid place-items-center">
              <GothicTabsList className="flex flex-row gap-2 self-center">
                <GothicTabsTrigger value="login">Log In</GothicTabsTrigger>
                <GothicTabsTrigger value="signup">Sign Up</GothicTabsTrigger>
              </GothicTabsList>
            </PanelHeader>
            <GothicCardContent className="pt-8">
              <TabsContent value="signup">
                <SignupForm />
              </TabsContent>
              <TabsContent value="login">
                <LoginForm />
              </TabsContent>
            </GothicCardContent>
          </GothicCard>
        </Tabs>
        <div className="text-foreground/40 font-serif *:[a]:hover:text-primary text-center text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
          Copyright © 2026 snaxm games. All rights reserved.
        </div>
      </div>
    </div>
  )
}

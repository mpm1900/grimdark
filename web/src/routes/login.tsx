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
    <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm space-y-2">
        <a className="flex items-center justify-center gap-2 self-center font-medium p-4 mb-9">
          <div className="text-foreground/60 flex size-16 items-center justify-center rounded-md">
            <GiWingedSword className="size-16" />
          </div>
        </a>

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

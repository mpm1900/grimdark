import { Field, FieldGroup, FieldLabel, FieldSet } from '../ui/field'
import { Input } from '../ui/input'
import { useNavigate } from '@tanstack/react-router'
import { useForm } from '@tanstack/react-form'
import { signupSchema } from '#/lib/auth/auth-schema'
import { toast } from 'sonner'
import { useSignup } from '#/lib/mutations/signup'
import { cn } from '#/lib/utils'
import { GothicFramedButton } from '../gothic-ui/button'
import { Marker, MarkerContent } from '../ui/marker'

function SignupForm() {
  const signup = useSignup()
  const navigate = useNavigate({ from: '/login' })
  const form = useForm({
    defaultValues: {
      username: '',
      email: '',
      password: '',
      secret: '',
    },
    validators: {
      onChange: signupSchema,
    },
    onSubmit: async ({ value }) => {
      try {
        await signup.mutateAsync(value)
        await navigate({ to: '/' })
      } catch (e) {
        console.error(e)
        toast.error(String(e))
      }
    },
  })
  return (
    <form className={cn('flex flex-col')} onSubmit={form.handleSubmit}>
      <FieldSet className="w-full">
        <FieldGroup className="gap-4">
          <form.Field name="username">
            {(field) => (
              <Field className="gap-1">
                <FieldLabel htmlFor="signup-username">
                  <Marker variant="separator">
                    <MarkerContent>Username</MarkerContent>
                  </Marker>
                </FieldLabel>
                <Input
                  id="signup-username"
                  type="text"
                  placeholder="username"
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                />
              </Field>
            )}
          </form.Field>
          <form.Field name="email">
            {(field) => (
              <Field className="gap-1">
                <FieldLabel htmlFor="signup-email">
                  <Marker variant="separator">
                    <MarkerContent>Email</MarkerContent>
                  </Marker>
                </FieldLabel>
                <Input
                  id="signup-email"
                  type="email"
                  placeholder="you@email.com"
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                />
              </Field>
            )}
          </form.Field>
          <form.Field name="password">
            {(field) => (
              <Field className="gap-1">
                <FieldLabel htmlFor="signup-password">
                  <Marker variant="separator">
                    <MarkerContent>Password</MarkerContent>
                  </Marker>
                </FieldLabel>
                <Input
                  id="signup-password"
                  type="password"
                  placeholder="••••••••"
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                />
              </Field>
            )}
          </form.Field>
          <form.Field name="secret">
            {(field) => (
              <Field className="gap-1">
                <FieldLabel htmlFor="signup-secret">
                  <Marker variant="separator">
                    <MarkerContent>Secret</MarkerContent>
                  </Marker>
                </FieldLabel>
                <Input
                  id="signup-secret"
                  type="password"
                  placeholder="Invite secret"
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                />
              </Field>
            )}
          </form.Field>
          <form.Subscribe>
            {({ canSubmit }) => (
              <div className="flex justify-end -mr-4 -mb-1">
                <GothicFramedButton
                  variant="red"
                  type="submit"
                  className="-mr-px"
                  disabled={!canSubmit}
                  onClick={() => form.handleSubmit()}
                >
                  Sign Up
                </GothicFramedButton>
              </div>
            )}
          </form.Subscribe>
        </FieldGroup>
      </FieldSet>
    </form>
  )
}

export { SignupForm }

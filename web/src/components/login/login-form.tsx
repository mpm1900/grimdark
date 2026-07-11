import { Field, FieldGroup, FieldLabel, FieldSet } from '../ui/field'
import { Input } from '../ui/input'
import { useLogin } from '#/lib/mutations/login'
import { useNavigate } from '@tanstack/react-router'
import { useForm } from '@tanstack/react-form'
import { authSchema } from '#/lib/auth/auth-schema'
import { toast } from 'sonner'
import { cn } from '#/lib/utils'
import { GothicFramedButton } from '../gothic-ui/button'
import { Marker, MarkerContent } from '../ui/marker'

function LoginForm() {
  const login = useLogin()
  const navigate = useNavigate({ from: '/login' })
  const form = useForm({
    defaultValues: {
      email: '',
      password: '',
      secret: '',
    },
    validators: {
      onChange: authSchema,
    },
    onSubmit: async ({ value }) => {
      try {
        await login.mutateAsync(value)
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
          <form.Field name="email">
            {(field) => (
              <Field className="gap-1">
                <FieldLabel htmlFor="email">
                  <Marker variant="separator">
                    <MarkerContent>Email</MarkerContent>
                  </Marker>
                </FieldLabel>
                <Input
                  id="email"
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
                <FieldLabel htmlFor="password">
                  <Marker variant="separator">
                    <MarkerContent>Password</MarkerContent>
                  </Marker>
                </FieldLabel>
                <Input
                  id="password"
                  type="password"
                  placeholder="••••••••"
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
                  Log In
                </GothicFramedButton>
              </div>
            )}
          </form.Subscribe>
        </FieldGroup>
      </FieldSet>
    </form>
  )
}

export { LoginForm }

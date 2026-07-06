import type { Actor } from '#/lib/game/actor'
import { ActorFlag } from './actor-flag'
import { Field, FieldContent, FieldLabel } from './ui/field'
import { Marker, MarkerContent } from './ui/marker'
import { WeaponDetails } from './weapon-details'
import { ActionContextDialog } from './action-context-dialog'
import { DialogTrigger } from './ui/dialog'
import { ActionButton } from './action-button'
import { getAppliedEffects } from '#/lib/game/game'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { GothicBadge } from './gothic-ui/badge'
import { ActorFrame } from './actor-frame'
import { GothicHighlightFrame } from './gothic-ui/frame'
import { StatsTable } from './stats-table'
import { AffinitiesTable } from './affinities-table'

function ActorCharacter({ actor }: { actor: Actor }) {
  return (
    <div className="flex flex-col gap-4">
      <ActorFrame actor={actor} className="-mt-1.5 -ml-1" />

      <div className="font-serif space-y-2 text-center px-3">
        <div className="grid grid-cols-3 text-sm">
          <Field className="gap-0">
            <FieldLabel className="text-foreground/60 font-cinzel">
              Race
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">
              {actor.race}
            </FieldContent>
          </Field>
          <Field className="gap-0">
            <FieldLabel className="text-foreground/60 font-cinzel">
              Faction
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">
              {actor.faction}
            </FieldContent>
          </Field>
          <Field className="gap-0">
            <FieldLabel className="text-foreground/60 font-cinzel">
              Class
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">
              --
            </FieldContent>
          </Field>
        </div>
        <div className="grid grid-cols-2 text-sm">
          <Field className="gap-0 text-foreground">
            <FieldLabel className="text-foreground/60 font-cinzel">
              State
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">
              {actor.state}
            </FieldContent>
          </Field>
          <Field className="gap-0">
            <FieldLabel className="text-foreground/60 font-cinzel">
              Status
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">
              {actor.status}
            </FieldContent>
          </Field>
        </div>
      </div>
    </div>
  )
}
function ActorState({
  actor,
  applied_effects,
}: {
  actor: Actor
  applied_effects: ReturnType<typeof getAppliedEffects>
}) {
  return (
    <div className="flex flex-col gap-4 px-2">
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Abilities</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent className="min-w-0 flex-row flex-wrap gap-2">
          {actor.effects
            .filter((e) => !!e.name)
            .map((effect) => (
              <GothicBadge
                variant="default"
                key={effect.ID}
                className="capitalize"
              >
                {effect.name}
              </GothicBadge>
            ))}
        </FieldContent>
      </Field>
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Flags</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent className="flex flex-col gap-4 font-serif">
          <div className="min-w-0 flex flex-row flex-wrap gap-2">
            <ActorFlag actor={actor} flag="is_active">
              Active
            </ActorFlag>
            <ActorFlag actor={actor} flag="is_alive">
              Alive
            </ActorFlag>
            <ActorFlag actor={actor} flag="is_hidden">
              Hidden
            </ActorFlag>
            <ActorFlag actor={actor} flag="is_protected">
              Protected
            </ActorFlag>
            <ActorFlag actor={actor} flag="is_stunned">
              Stunned
            </ActorFlag>
          </div>
        </FieldContent>
      </Field>
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Active Effects</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent className="min-w-0 flex-row flex-wrap gap-2">
          {applied_effects.map((effect) => (
            <GothicBadge variant="empty" key={effect.ID} className="capitalize">
              {effect.name}
              {effect.count > 1 && `(${effect.count})`}
            </GothicBadge>
          ))}
          {applied_effects.length === 0 && (
            <span className="italic text-muted-foreground">None</span>
          )}
        </FieldContent>
      </Field>
    </div>
  )
}

function ActorStats({ actor }: { actor: Actor }) {
  return (
    <div className="flex min-w-2/5 xl:min-w-1/3 flex-col gap-4 text-foreground">
      <GothicHighlightFrame className="-mx-2">
        <StatsTable actor={actor} />
        <AffinitiesTable actor={actor} />
      </GothicHighlightFrame>
    </div>
  )
}
function ActorWeaponsAndActions({ actor }: { actor: Actor }) {
  return (
    <div className="flex min-w-0 flex-col">
      {actor.weapon && <WeaponDetails weapon={actor.weapon} />}
      <Marker variant="separator" className="hidden">
        <MarkerContent>Actions</MarkerContent>
      </Marker>
      <div className="flex flex-col gap-0">
        {actor.actions.map((action) => (
          <ActionContextDialog
            key={action.ID}
            actor={actor}
            action={action}
            enabled={!action.is_disabled}
          >
            <DialogTrigger asChild>
              <ActionButton action={action} actor={actor} />
            </DialogTrigger>
          </ActionContextDialog>
        ))}
      </div>
    </div>
  )
}

function ActorDetails({ actor }: { actor: Actor }) {
  const game = useSelector(gameStore, (g) => g)
  const applied_effects = getAppliedEffects(game, actor)

  return (
    <div className="flex max-w-4xl">
      <div className="flex min-w-0 flex-col gap-4">
        <ActorCharacter actor={actor} />
        <ActorState actor={actor} applied_effects={applied_effects} />
      </div>
      <ActorStats actor={actor} />
      <ActorWeaponsAndActions actor={actor} />
    </div>
  )
}

export { ActorDetails }

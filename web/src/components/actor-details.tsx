import type { Actor } from '#/lib/game/actor'
import { AFFINITIES, mapStage, STATS, type Stat } from '#/lib/game/core'
import { ActorFlag } from './actor-flag'
import { AffinityName } from './affinity-name'
import { HealthBar } from './health-bar'
import { AspectRatio } from './ui/aspect-ratio'
import { Badge } from './ui/badge'
import { Field, FieldContent, FieldLabel } from './ui/field'
import { Marker, MarkerContent } from './ui/marker'
import { Table, TableBody, TableCell, TableRow } from './ui/table'
import {
  AffinityDamageValue,
  AffinityMultiplier,
  AffinityResistanceValue,
} from './affinity-value'
import { WeaponDetails } from './weapon-details'
import { ActionContextDialog } from './action-context-dialog'
import { DialogTrigger } from './ui/dialog'
import { ActionButton } from './action-button'
import { StatValue } from './stat-value'
import { cn } from '#/lib/utils'
import { getAppliedEffects } from '#/lib/game/game'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'

function StatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  const mod = stat === 'accuracy' || stat === 'evasion' ? 3 : 2
  const mult = mapStage(stage, mod, 1)
  return (
    <TableRow>
      <TableCell className="capitalize">{stat}</TableCell>
      <TableCell className="text-end">
        <StatValue actor={actor} stat={stat} />
      </TableCell>
      <TableCell className="text-end">
        {stage != 0 && <span>{stage}</span>}
      </TableCell>
      <TableCell>
        {stage != 0 && (
          <span className={cn(mult === 1 && 'opacity-45')}>
            x{mult.toFixed(2)}
          </span>
        )}
      </TableCell>
    </TableRow>
  )
}
function StatsTable({ actor }: { actor: Actor }) {
  return (
    <Table>
      <TableBody>
        {STATS.filter((s) => s !== 'health').map((stat) => (
          <StatRow key={stat} actor={actor} stat={stat} />
        ))}
      </TableBody>
    </Table>
  )
}

function AffinityDamageTable({ actor }: { actor: Actor }) {
  return (
    <Table>
      <TableBody>
        {AFFINITIES.map((affinity) => (
          <TableRow key={affinity}>
            <TableCell>
              <AffinityName affinity={affinity} />
            </TableCell>
            <TableCell>
              <AffinityDamageValue actor={actor} affinity={affinity} />
            </TableCell>
            <TableCell>
              <AffinityMultiplier value={actor.affinity_damage[affinity]} />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}
function AffinityResistanceTable({ actor }: { actor: Actor }) {
  return (
    <Table>
      <TableBody>
        {AFFINITIES.map((affinity) => (
          <TableRow key={affinity}>
            <TableCell>
              <AffinityName affinity={affinity} />
            </TableCell>
            <TableCell>
              <AffinityResistanceValue actor={actor} affinity={affinity} />
            </TableCell>
            <TableCell>
              <AffinityMultiplier
                value={actor.affinity_resistance[affinity] * -1}
              />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}

function ActorCharacter({ actor }: { actor: Actor }) {
  return (
    <div className="flex flex-col gap-2">
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Character Lv.{actor.level}</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent className="flex flex-row justify-between items-center">
          <span className="text-md font-black">{actor.name}</span>
          <div className="flex gap-1">
            {AFFINITIES.filter((a) => actor.affinities.includes(a)).map((a) => (
              <AffinityName key={a} affinity={a} />
            ))}
          </div>
        </FieldContent>
      </Field>
      <div>
        <HealthBar type="value" actor={actor} />
      </div>
      <AspectRatio ratio={16 / 9} className="rounded-lg bg-muted" />
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
    <div className="flex flex-col gap-4">
      <div className="grid grid-cols-3 text-xs">
        <Field>
          <FieldLabel>
            <Marker variant="separator">
              <MarkerContent>Augment</MarkerContent>
            </Marker>
          </FieldLabel>
          <FieldContent className="text-center capitalize">
            {actor.augment}
          </FieldContent>
        </Field>
        <Field>
          <FieldLabel>
            <Marker variant="separator">
              <MarkerContent>State</MarkerContent>
            </Marker>
          </FieldLabel>
          <FieldContent className="text-center capitalize">
            {actor.state}
          </FieldContent>
        </Field>
        <Field>
          <FieldLabel>
            <Marker variant="separator">
              <MarkerContent>Status</MarkerContent>
            </Marker>
          </FieldLabel>
          <FieldContent className="text-center capitalize">
            {actor.status}
          </FieldContent>
        </Field>
      </div>
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Flags</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent className="flex-row flex-wrap gap-2 min-w-0">
          <ActorFlag actor={actor} flag="is_active">
            Active
          </ActorFlag>
          <ActorFlag actor={actor} flag="is_alive">
            Alive
          </ActorFlag>
          <ActorFlag actor={actor} flag="is_protected">
            Protected
          </ActorFlag>
          <ActorFlag actor={actor} flag="is_staggered">
            Staggered
          </ActorFlag>
          <ActorFlag actor={actor} flag="is_stunned">
            Stunned
          </ActorFlag>
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
            <Badge variant="outline" key={effect.ID} className="capitalize">
              {effect.name}
              {effect.count > 1 && `(${effect.count})`}
            </Badge>
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
    <div className="flex min-w-0 flex-col gap-4">
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Stats</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent>
          <StatsTable actor={actor} />
        </FieldContent>
      </Field>
      <div className="grid grid-cols-2">
        <Field>
          <FieldLabel>
            <Marker variant="separator">
              <MarkerContent>Resistances</MarkerContent>
            </Marker>
          </FieldLabel>
          <FieldContent>
            <AffinityResistanceTable actor={actor} />
          </FieldContent>
        </Field>
        <Field>
          <FieldLabel>
            <Marker variant="separator">
              <MarkerContent>Damage</MarkerContent>
            </Marker>
          </FieldLabel>
          <FieldContent>
            <AffinityDamageTable actor={actor} />
          </FieldContent>
        </Field>
      </div>
    </div>
  )
}
function ActorWeaponsAndActions({ actor }: { actor: Actor }) {
  return (
    <div className="flex min-w-0 flex-col gap-4">
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Weapon</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent>
          {actor.weapon && <WeaponDetails weapon={actor.weapon} />}
        </FieldContent>
      </Field>
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Source Effects</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent className="min-w-0 flex-row flex-wrap gap-2">
          {actor.effects
            .filter((e) => !!e.name)
            .map((effect) => (
              <Badge variant="outline" key={effect.ID} className="capitalize">
                {effect.name}
              </Badge>
            ))}
        </FieldContent>
      </Field>
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Actions</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent className="flex flex-col gap-2">
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
        </FieldContent>
      </Field>
    </div>
  )
}

function ActorDetails({ actor }: { actor: Actor }) {
  const game = useSelector(gameStore, (g) => g)
  const applied_effects = getAppliedEffects(game, actor)

  return (
    <div className="grid grid-cols-3 gap-3 max-w-4xl">
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

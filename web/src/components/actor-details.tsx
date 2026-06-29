import type { Actor } from '#/lib/game/actor'
import {
  AFFINITIES,
  mapStage,
  type Stat,
  MAIN_STATS,
  ACCURACY_STATS,
  CRITICAL_STATS,
} from '#/lib/game/core'
import { ActorFlag } from './actor-flag'
import { AffinityName } from './affinity-name'
import { HealthBar } from './health-bar'
import { Badge } from './ui/badge'
import { Field, FieldContent, FieldLabel } from './ui/field'
import { Marker, MarkerContent } from './ui/marker'
import { Table, TableBody, TableCell, TableRow } from './ui/table'
import {
  AffinityDamageMultiplier,
  AffinityDamageValue,
  AffinityResistanceMultiplier,
  AffinityResistanceValue,
} from './affinity-value'
import { WeaponDetails } from './weapon-details'
import { ActionContextDialog } from './action-context-dialog'
import { DialogTrigger } from './ui/dialog'
import { ActionButton } from './action-button'
import { StatValue } from './stat-value'
import { cn, sign } from '#/lib/utils'
import { getAppliedEffects } from '#/lib/game/game'
import { useSelector } from '@tanstack/react-store'
import { gameStore } from '#/lib/stores/game'
import { StatName } from './stat-name'

function MainStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  const mod = 2
  const mult = mapStage(stage, mod, 1)
  return (
    <TableRow>
      <TableCell className="capitalize">
        <StatName stat={stat}>{stat}</StatName>
      </TableCell>
      <TableCell className="text-end">
        <StatValue actor={actor} stat={stat}>
          {actor.stats[stat]}
        </StatValue>
      </TableCell>
      <TableCell className="text-end">
        {stage != 0 && (
          <StatValue actor={actor} stat={stat}>
            {stage}
          </StatValue>
        )}
      </TableCell>
      <TableCell className="text-end">
        {stage != 0 && (
          <span className={cn(mult === 1 && 'opacity-45')}>
            x{mult.toFixed(2)}
          </span>
        )}
      </TableCell>
    </TableRow>
  )
}
function AccuracyStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  const mod = 3
  const mult = mapStage(stage, mod, 1)
  return (
    <TableRow>
      <TableCell className="capitalize">{stat}</TableCell>
      <TableCell className="text-end">
        <StatValue actor={actor} stat={stat}>
          x{(actor.stats[stat] / 100).toFixed(2)}
        </StatValue>
      </TableCell>
      <TableCell className="text-end">
        {stage != 0 && (
          <StatValue actor={actor} stat={stat}>
            {stage}
          </StatValue>
        )}
      </TableCell>
      <TableCell className="text-end">
        {stage != 0 && (
          <span className={cn(mult === 1 && 'opacity-45')}>
            x{mult.toFixed(2)}
          </span>
        )}
      </TableCell>
    </TableRow>
  )
}
function CriticalStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  return (
    <TableRow>
      <TableCell className="capitalize">{stat}</TableCell>
      <TableCell className="text-end">
        <span
          className={cn({
            'text-green-400': stage > 0,
            'text-red-400': stage < 0,
            'opacity-45': stage === 0,
          })}
        >
          {sign(stage)}
          {stage}
        </span>
      </TableCell>
      <TableCell className="text-end">
        {stage != 0 && (
          <span
            className={cn({
              'text-green-400': stage > 0,
              'text-red-400': stage < 0,
              'opacity-45': stage === 0,
            })}
          >
            {stage}
          </span>
        )}
      </TableCell>
      <TableCell className="text-end">
        <span className="opacity-45">--</span>
      </TableCell>
    </TableRow>
  )
}
function StatsTable({ actor }: { actor: Actor }) {
  return (
    <Table>
      <TableBody>
        {MAIN_STATS.map((stat) => (
          <MainStatRow key={stat} actor={actor} stat={stat} />
        ))}
        {ACCURACY_STATS.map((stat) => (
          <AccuracyStatRow key={stat} actor={actor} stat={stat} />
        ))}
        <CriticalStatRow actor={actor} stat="critical-chance" />
        <AccuracyStatRow actor={actor} stat="critical-damage" />
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
            <TableCell className="text-end">
              <AffinityDamageValue actor={actor} affinity={affinity} />
            </TableCell>
            <TableCell className="text-end">
              <AffinityDamageMultiplier actor={actor} affinity={affinity} />
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
            <TableCell className="text-end">
              <AffinityResistanceValue actor={actor} affinity={affinity} />
            </TableCell>
            <TableCell className="text-end">
              <AffinityResistanceMultiplier
                actor={actor}
                affinity={affinity}
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
    <div className="flex flex-col gap-4">
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
      <div className="flex flex-col gap-4">
        <div className="grid grid-cols-2 text-sm">
          <Field>
            <FieldLabel className="text-muted-foreground">Race</FieldLabel>
            <FieldContent className="capitalize">{actor.race}</FieldContent>
          </Field>
          <Field>
            <FieldLabel className="text-muted-foreground">Faction</FieldLabel>
            <FieldContent className="capitalize">{actor.faction}</FieldContent>
          </Field>
        </div>
        <div className="grid grid-cols-3 text-sm">
          <Field>
            <FieldLabel className="text-muted-foreground">Class</FieldLabel>
            <FieldContent className="capitalize">--</FieldContent>
          </Field>
          <Field>
            <FieldLabel className="text-muted-foreground">Subclass</FieldLabel>
            <FieldContent className="capitalize">--</FieldContent>
          </Field>
          <Field>
            <FieldLabel className="text-muted-foreground">Augment</FieldLabel>
            <FieldContent className="capitalize">{actor.augment}</FieldContent>
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
    <div className="flex flex-col gap-4">
      <Field>
        <FieldLabel>
          <Marker variant="separator">
            <MarkerContent>Flags</MarkerContent>
          </Marker>
        </FieldLabel>
        <FieldContent className="flex flex-col gap-4">
          <div className="grid grid-cols-2 text-sm">
            <Field>
              <FieldLabel className="text-muted-foreground">State</FieldLabel>
              <FieldContent className="capitalize">{actor.state}</FieldContent>
            </Field>
            <Field>
              <FieldLabel className="text-muted-foreground">Status</FieldLabel>
              <FieldContent className="capitalize">{actor.status}</FieldContent>
            </Field>
          </div>
          <div className="flex flex-row flex-wrap gap-2 min-w-0">
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
            <ActorFlag actor={actor} flag="is_staggered">
              Staggered
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
            <MarkerContent>Weapon(s)</MarkerContent>
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
    <div className="grid grid-cols-3 gap-4 max-w-4xl">
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

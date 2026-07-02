import type { Actor } from '#/lib/game/actor'
import {
  AFFINITIES,
  mapStage,
  type Stat,
  MAIN_STATS,
  ACCURACY_STATS,
  STAT_LABELS,
} from '#/lib/game/core'
import { ActorFlag } from './actor-flag'
import { AffinityName } from './affinity-name'
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
import { GothicBadge } from './gothic-ui/badge'
import { ActorFrame } from './actor-frame'
import { GothicHighlightFrame } from './gothic-ui/frame'

function MainStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  const mod = 2
  const mult = mapStage(stage, mod, 1)
  return (
    <TableRow>
      <TableCell className="capitalize font-cinzel">
        <StatName stat={stat}>{STAT_LABELS[stat]}</StatName>
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
          <span
            className={cn(
              {
                'text-positive': stage > 0,
                'text-negative': stage < 0,
              },
              mult === 1 && 'text-foreground/20'
            )}
          >
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
      <TableCell className="capitalize font-cinzel">
        {STAT_LABELS[stat]}
      </TableCell>
      <TableCell className="text-end"></TableCell>
      <TableCell className="text-end">
        {!!stage && (
          <StatValue actor={actor} stat={stat}>
            {sign(stage)}
            {Math.abs(stage)}
          </StatValue>
        )}
      </TableCell>
      <TableCell className="text-end">
        <StatValue
          actor={actor}
          stat={stat}
          className={cn(actor.stats[stat] === 100 && 'text-foreground/20')}
        >
          x{mult.toFixed(2)}
        </StatValue>
      </TableCell>
    </TableRow>
  )
}
function CriticalStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  return (
    <TableRow>
      <TableCell className="capitalize font-cinzel">
        {STAT_LABELS[stat]}
      </TableCell>
      <TableCell className="text-end"></TableCell>
      <TableCell className="text-end">
        {!!stage && (
          <span
            className={cn({
              'text-positive': stage > 0,
              'text-negative': stage < 0,
            })}
          >
            {sign(stage)}
            {Math.abs(stage)}
          </span>
        )}
      </TableCell>
      <TableCell className="text-end text-foreground/20">--</TableCell>
    </TableRow>
  )
}
function StatsTable({ actor }: { actor: Actor }) {
  return (
    <Table className="font-mono">
      <TableBody>
        {MAIN_STATS.map((stat) => (
          <MainStatRow key={stat} actor={actor} stat={stat} />
        ))}
        {ACCURACY_STATS.map((stat) => (
          <AccuracyStatRow key={stat} actor={actor} stat={stat} />
        ))}
        <AccuracyStatRow actor={actor} stat="critical-damage" />
        <CriticalStatRow actor={actor} stat="critical-chance" />
      </TableBody>
    </Table>
  )
}

function AffinityDamageTable({ actor }: { actor: Actor }) {
  return (
    <Table className="font-mono">
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
    <Table className="font-mono">
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
      <ActorFrame actor={actor} className="-mt-1.5 -ml-1" />

      <div className="font-serif space-y-2 text-center px-3">
        <div className="grid grid-cols-3 text-sm">
          <Field className="gap-0">
            <FieldLabel className="text-foreground/60 font-cinzel">
              Race
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">{actor.race}</FieldContent>
          </Field>
          <Field className="gap-0">
            <FieldLabel className="text-foreground/60 font-cinzel">
              Faction
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">{actor.faction}</FieldContent>
          </Field>
          <Field className="gap-0">
            <FieldLabel className="text-foreground/60 font-cinzel">
              Class
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">--</FieldContent>
          </Field>
        </div>
        <div className="grid grid-cols-3 text-sm">
          <Field className="gap-0">
            <FieldLabel className="text-foreground/60 font-cinzel">
              Augment
            </FieldLabel>
            <FieldContent className="capitalize text-foreground">{actor.augment}</FieldContent>
          </Field>
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
        <div className="grid grid-cols-2 p-3 -mx-2">
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

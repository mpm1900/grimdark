import type { Actor } from '#/lib/game/actor'
import {
  ACCURACY_STATS,
  MAIN_STATS,
  mapStage,
  STAT_LABELS,
  type Stat,
} from '#/lib/game/core'
import { cn, sign } from '#/lib/utils'
import { DNumber } from './dnumber'
import { StatName } from './stat-name'
import { StatValue } from './stat-value'
import { Table, TableBody, TableCell, TableRow } from './ui/table'

function MainStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  const mult =
    (actor.stats[stat] - (actor.offset_stats[stat] ?? 0)) /
    actor.unmodified_stats[stat]
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
          <DNumber
            value={mult}
            r={1}
            className={cn(mult === 1 && 'text-foreground/20')}
          >
            x{mult.toFixed(2)}
          </DNumber>
        )}
      </TableCell>
    </TableRow>
  )
}
function MultiplierStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  const mod = 3
  const mult = mapStage(stage, mod, 1)

  return (
    <TableRow>
      <TableCell className="capitalize font-cinzel" colSpan={2}>
        {STAT_LABELS[stat]}
      </TableCell>
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
          className={cn(mult == 1 && 'text-foreground/20!')}
        >
          x{mult.toFixed(2)}
        </StatValue>
      </TableCell>
    </TableRow>
  )
}
function OffsetStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const stage = actor.stages[stat]
  return (
    <TableRow>
      <TableCell className="capitalize font-cinzel" colSpan={2}>
        {STAT_LABELS[stat]}
      </TableCell>
      <TableCell className="text-end">
        {!!stage && (
          <DNumber value={stage}>
            {sign(stage)}
            {Math.abs(stage)}
          </DNumber>
        )}
      </TableCell>
      <TableCell className="text-end text-foreground/20">--</TableCell>
    </TableRow>
  )
}
function RawStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  const value = actor.stats[stat]
  return (
    <TableRow>
      <TableCell className="capitalize font-cinzel" colSpan={2}>
        {STAT_LABELS[stat]}
      </TableCell>
      <TableCell className="text-end">
        <DNumber value={value} r={1}>
          {Math.abs(value)}
        </DNumber>
      </TableCell>
      <TableCell className="text-end text-foreground/20">--</TableCell>
    </TableRow>
  )
}

function StatsTable({
  actor,
  className,
  ...props
}: React.ComponentProps<typeof Table> & { actor: Actor }) {
  return (
    <Table className={cn('font-mono', className)} {...props}>
      <TableBody>
        {MAIN_STATS.map((stat) => (
          <MainStatRow key={stat} actor={actor} stat={stat} />
        ))}
      </TableBody>
    </Table>
  )
}

function OtherStatsTable({
  actor,
  className,
  ...props
}: React.ComponentProps<typeof Table> & { actor: Actor }) {
  return (
    <Table className={cn('font-mono', className)} {...props}>
      <TableBody>
        {ACCURACY_STATS.map((stat) => (
          <MultiplierStatRow key={stat} actor={actor} stat={stat} />
        ))}
        <MultiplierStatRow actor={actor} stat="critical-damage" />
        <OffsetStatRow actor={actor} stat="critical-chance" />
        <OffsetStatRow actor={actor} stat="damage-reflect" />
        <MultiplierStatRow actor={actor} stat="effect-chance" />
        <RawStatRow actor={actor} stat="actions" />
      </TableBody>
    </Table>
  )
}

export { StatsTable, OtherStatsTable }

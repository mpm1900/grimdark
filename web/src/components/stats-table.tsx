import type { Actor } from '#/lib/game/actor'
import {
  ACCURACY_STATS,
  MAIN_STATS,
  mapStage,
  STAT_LABELS,
  type Stat,
} from '#/lib/game/core'
import { cn, sign } from '#/lib/utils'
import { StatName } from './stat-name'
import { StatValue } from './stat-value'
import { Table, TableBody, TableCell, TableRow } from './ui/table'

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
function MultiplierStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
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
function OffsetStatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
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
        {ACCURACY_STATS.map((stat) => (
          <MultiplierStatRow key={stat} actor={actor} stat={stat} />
        ))}
        <MultiplierStatRow actor={actor} stat="critical-damage" />
        <OffsetStatRow actor={actor} stat="critical-chance" />
        <OffsetStatRow actor={actor} stat="range" />
      </TableBody>
    </Table>
  )
}

export { StatsTable }

import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { Marker, MarkerContent } from './ui/marker'
import { Table, TableBody, TableCell, TableRow } from './ui/table'
import { AffinityName } from './affinity-name'
import {
  AffinityDamageMultiplier,
  AffinityDamageValue,
  AffinityResistanceMultiplier,
  AffinityResistanceValue,
} from './affinity-value'
import { AFFINITIES } from '#/lib/game/core'

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

function AffinitiesTable({
  actor,
  className,
  ...props
}: React.ComponentProps<'div'> & { actor: Actor }) {
  return (
    <div
      className={cn('grid grid-cols-2 p-3 -mx-2 text-foreground/80', className)}
      {...props}
    >
      <div>
        <Marker variant="separator">
          <MarkerContent>Damage</MarkerContent>
        </Marker>
        <AffinityDamageTable actor={actor} />
      </div>
      <div>
        <Marker variant="separator">
          <MarkerContent>Resistances</MarkerContent>
        </Marker>
        <AffinityResistanceTable actor={actor} />
      </div>
    </div>
  )
}

export { AffinitiesTable }

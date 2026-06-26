import {
  AffinityDamageValue,
  AffinityResistanceValue,
} from '#/components/affinity-value'
import { AppHeader } from '#/components/app-header'
import { StatValue } from '#/components/stat-value'
import { Field, FieldContent, FieldLabel } from '#/components/ui/field'
import { Item, ItemContent, ItemTitle } from '#/components/ui/item'
import { Separator } from '#/components/ui/separator'
import { Table, TableBody, TableCell, TableRow } from '#/components/ui/table'
import type { Actor } from '#/lib/game/actor'
import { AFFINITIES, STATS, type Stat } from '#/lib/game/core'
import { getAppliedModifiers, type Game } from '#/lib/game/game'
import { RenderLog } from '#/lib/game/log'
import { gameStore } from '#/lib/stores/game'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'

export const Route = createFileRoute('/')({ component: Home })

function StatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  return (
    <TableRow>
      <TableCell className="capitalize">{stat}</TableCell>
      <TableCell>
        <StatValue actor={actor} stat={stat} />
      </TableCell>
      <TableCell>
        {actor.stages[stat] != 0 && <span>{actor.stages[stat]}</span>}
      </TableCell>
    </TableRow>
  )
}

function ActorTest({ game, actor }: { game: Game; actor: Actor }) {
  return (
    <Item variant="outline">
      <ItemContent className="gap-4">
        <div className="flex justify-between">
          <ItemTitle className="font-bold">
            {actor.name}, lv{actor.level}
          </ItemTitle>
          <div className="flex gap-1">
            {AFFINITIES.filter((a) => actor.affinities.includes(a)).map((a) => (
              <span key={a} className="capitalize">
                {a}
              </span>
            ))}
          </div>
        </div>
        <Separator />
        <div className="grid grid-cols-2 gap-3">
          <Field>
            <FieldLabel>Stats</FieldLabel>
            <FieldContent>
              <Table>
                <TableBody>
                  <TableRow>
                    <TableCell>Health</TableCell>
                    <TableCell>
                      {actor.stats.health - actor.damage}/{actor.stats.health}
                    </TableCell>
                  </TableRow>
                  {STATS.filter((s) => s !== 'health').map((stat) => (
                    <StatRow key={stat} actor={actor} stat={stat} />
                  ))}
                </TableBody>
              </Table>
            </FieldContent>
          </Field>
          <div className="flex flex-col gap-4">
            <div className="grid grid-cols-2">
              <Field>
                <FieldLabel>Affinity Res</FieldLabel>
                <FieldContent>
                  <Table>
                    <TableBody>
                      {AFFINITIES.map((affinity) => (
                        <TableRow key={affinity}>
                          <TableCell className="capitalize">
                            {affinity}
                          </TableCell>
                          <TableCell>
                            <AffinityResistanceValue
                              actor={actor}
                              affinity={affinity}
                            />
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </FieldContent>
              </Field>
              <Field>
                <FieldLabel>Affinity Dmg</FieldLabel>
                <FieldContent>
                  <Table>
                    <TableBody>
                      {AFFINITIES.map((affinity) => (
                        <TableRow key={affinity}>
                          <TableCell className="capitalize">
                            {affinity}
                          </TableCell>
                          <TableCell>
                            <AffinityDamageValue
                              actor={actor}
                              affinity={affinity}
                            />
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </FieldContent>
              </Field>
            </div>
            <Field>
              <FieldLabel>Applied Modifiers</FieldLabel>
              <FieldContent className="flex flex-row flex-wrap gap-1">
                {getAppliedModifiers(game, actor).map((mod) => (
                  <span key={mod.ID}>{mod.payload.name},</span>
                ))}
              </FieldContent>
            </Field>
          </div>
        </div>
      </ItemContent>
    </Item>
  )
}

function Home() {
  const game = useSelector(gameStore, (g) => g)
  const actors = useSelector(gameStore, (g) => g.actors)
  return (
    <ClientOnly>
      <div>
        <AppHeader />
        <div className="pt-12 flex">
          <div className="p-4 flex flex-col gap-3">
            {actors.map((actor) => (
              <ActorTest key={actor.ID} game={game} actor={actor} />
            ))}
          </div>
          <ul className="p-4">
            {game.logs.map((log) => (
              <li key={log.ID}>{RenderLog(log)}</li>
            ))}
          </ul>
        </div>
      </div>
    </ClientOnly>
  )
}

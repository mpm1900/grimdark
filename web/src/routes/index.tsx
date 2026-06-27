import { AffinityName } from '#/components/affinity-name'
import {
  AffinityDamageValue,
  AffinityMultiplier,
  AffinityResistanceValue,
} from '#/components/affinity-value'
import { AppHeader } from '#/components/app-header'
import { StatValue } from '#/components/stat-value'
import { Field, FieldContent, FieldLabel } from '#/components/ui/field'
import { Item, ItemContent, ItemTitle } from '#/components/ui/item'
import { Separator } from '#/components/ui/separator'
import { Table, TableBody, TableCell, TableRow } from '#/components/ui/table'
import type { Actor } from '#/lib/game/actor'
import { AFFINITIES, mapStage, STATS, type Stat } from '#/lib/game/core'
import { getAppliedEffects, type Game } from '#/lib/game/game'
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
              <AffinityName key={a} affinity={a} />
            ))}
          </div>
        </div>
        <Separator />
        <div className="grid grid-cols-2 gap-3">
          <div className="flex flex-col gap-4">
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
            <Field>
              <FieldLabel>Status</FieldLabel>
              <FieldContent>{actor.status}</FieldContent>
            </Field>
          </div>
          <div className="flex flex-col gap-4">
            <div className="grid grid-cols-2">
              <Field>
                <FieldLabel>Affinity Res</FieldLabel>
                <FieldContent>
                  <Table>
                    <TableBody>
                      {AFFINITIES.map((affinity) => (
                        <TableRow key={affinity}>
                          <TableCell>
                            <AffinityName affinity={affinity} />
                          </TableCell>
                          <TableCell>
                            <AffinityResistanceValue
                              actor={actor}
                              affinity={affinity}
                            />
                          </TableCell>
                          <TableCell>
                            <AffinityMultiplier
                              value={actor.affinity_resistance[affinity]}
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
                          <TableCell>
                            <AffinityName affinity={affinity} />
                          </TableCell>
                          <TableCell>
                            <AffinityDamageValue
                              actor={actor}
                              affinity={affinity}
                            />
                          </TableCell>
                          <TableCell>
                            <AffinityMultiplier
                              value={actor.affinity_damage[affinity]}
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
              <FieldLabel>Flags</FieldLabel>
              <FieldContent className="flex flex-row gap-x-2 flex-wrap">
                {actor.is_active && <span>IsActive</span>}
                {actor.is_alive && <span>IsAlive</span>}
                {actor.is_protected && <span>IsProtected</span>}
                {actor.is_staggered && <span>IsStaggered</span>}
              </FieldContent>
            </Field>
            <Field>
              <FieldLabel>Applied Effects</FieldLabel>
              <FieldContent className="flex flex-row flex-wrap gap-x-2 w-60!">
                {getAppliedEffects(game, actor).map((effect) => (
                  <span key={effect.ID} className="capitalize">
                    {effect.name}
                    {effect.count > 1 && `(${effect.count})`}
                  </span>
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

import { AppHeader } from '#/components/app-header'
import { StatValue } from '#/components/stat-value'
import { Field, FieldContent, FieldLabel } from '#/components/ui/field'
import { Item, ItemContent, ItemTitle } from '#/components/ui/item'
import { Separator } from '#/components/ui/separator'
import { Table, TableBody, TableCell, TableRow } from '#/components/ui/table'
import type { Actor } from '#/lib/game/actor'
import type { Stat } from '#/lib/game/core'
import { getAppliedModifiers, type Game } from '#/lib/game/game'
import { RenderLog } from '#/lib/game/log'
import { gameStore } from '#/lib/stores/game'
import { keys } from '#/utils/maps'
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
            {actor.affinities.map((a) => (
              <span key={a}>{a}</span>
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
                  <StatRow actor={actor} stat="melee" />
                  <StatRow actor={actor} stat="ranged" />
                  <StatRow actor={actor} stat="special" />
                  <StatRow actor={actor} stat="martial-defense" />
                  <StatRow actor={actor} stat="special-defense" />
                  <StatRow actor={actor} stat="speed" />
                </TableBody>
              </Table>
            </FieldContent>
          </Field>
          <div>
            <Field>
              <FieldLabel>Affinity Resistance</FieldLabel>
              <FieldContent>
                <Table>
                  <TableBody>
                    {keys(actor.affinity_resistance).map((affinity) => (
                      <TableRow key={affinity}>
                        <TableCell className="capitalize">{affinity}</TableCell>
                        <TableCell>
                          {actor.affinity_resistance[affinity]}
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </FieldContent>
            </Field>
            <div className="flex gap-1">
              {getAppliedModifiers(game, actor).map((mod) => (
                <span key={mod.ID}>{mod.payload.name},</span>
              ))}
            </div>
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
              <li>{RenderLog(log)}</li>
            ))}
          </ul>
        </div>
      </div>
    </ClientOnly>
  )
}

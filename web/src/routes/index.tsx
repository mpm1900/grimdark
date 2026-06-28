import { ActorFlag } from '#/components/actor-flag'
import { AffinityName } from '#/components/affinity-name'
import {
  AffinityDamageValue,
  AffinityMultiplier,
  AffinityResistanceValue,
} from '#/components/affinity-value'
import { AppHeader } from '#/components/app-header'
import { StatValue } from '#/components/stat-value'
import { Badge } from '#/components/ui/badge'
import { Button } from '#/components/ui/button'
import { Field, FieldContent, FieldLabel } from '#/components/ui/field'
import {
  Item,
  ItemContent,
  ItemDescription,
  ItemTitle,
} from '#/components/ui/item'
import { Marker, MarkerContent } from '#/components/ui/marker'
import { Table, TableBody, TableCell, TableRow } from '#/components/ui/table'
import type { Actor } from '#/lib/game/actor'
import { NULL_CONTEXT, type Context } from '#/lib/game/context'
import { AFFINITIES, STATS, type Stat } from '#/lib/game/core'
import { getAppliedEffects, type Game } from '#/lib/game/game'
import { RenderLog } from '#/lib/game/log'
import { getTargetsQuery } from '#/lib/queries/get-targets'
import { gameStore } from '#/lib/stores/game'
import { useQuery } from '@tanstack/react-query'
import { ClientOnly, createFileRoute } from '@tanstack/react-router'
import { useSelector } from '@tanstack/react-store'

export const Route = createFileRoute('/')({ component: Home })

function StatRow({ actor, stat }: { actor: Actor; stat: Stat }) {
  return (
    <TableRow>
      <TableCell className="capitalize">{stat}</TableCell>
      <TableCell className="text-end">
        <StatValue actor={actor} stat={stat} />
      </TableCell>
      <TableCell className="text-end">
        {actor.stages[stat] != 0 && <span>{actor.stages[stat]}</span>}
      </TableCell>
    </TableRow>
  )
}

function ActorTest({ game, actor }: { game: Game; actor: Actor }) {
  const applied_effects = getAppliedEffects(game, actor)
  const context: Context = {
    ...NULL_CONTEXT,
    action_ID: actor.actions[0].ID,
    source_ID: actor.ID,
    parent_ID: actor.ID,
    player_ID: actor.player_ID,
    position_IDs: [],
  }
  const targets_query = useQuery(getTargetsQuery(context))
  console.log('data', targets_query.data)
  return (
    <Item variant="outline">
      <ItemContent className="gap-4">
        <div className="flex justify-between">
          <ItemTitle className="font-bold">
            {actor.name}, Lv.{actor.level}
          </ItemTitle>
          <div className="flex gap-1">
            {AFFINITIES.filter((a) => actor.affinities.includes(a)).map((a) => (
              <AffinityName key={a} affinity={a} />
            ))}
          </div>
        </div>
        <div className="grid grid-cols-3 gap-3">
          <div className="flex flex-col gap-4">
            <Field>
              <FieldLabel>
                <Marker variant="separator">
                  <MarkerContent>Stats</MarkerContent>
                </Marker>
              </FieldLabel>
              <FieldContent>
                <Table>
                  <TableBody>
                    <TableRow>
                      <TableCell>Health</TableCell>
                      <TableCell className="text-end">
                        {actor.stats.health - actor.wounds}/{actor.stats.health}
                      </TableCell>
                    </TableRow>
                    {STATS.filter((s) => s !== 'health').map((stat) => (
                      <StatRow key={stat} actor={actor} stat={stat} />
                    ))}
                  </TableBody>
                </Table>
              </FieldContent>
            </Field>
            <div className="flex gap-0">
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
            </div>
          </div>
          <div className="flex flex-col gap-4">
            <div className="grid grid-cols-2">
              <Field>
                <FieldLabel>
                  <Marker variant="separator">
                    <MarkerContent>Resistances</MarkerContent>
                  </Marker>
                </FieldLabel>
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
                              value={actor.affinity_resistance[affinity] * -1}
                            />
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </FieldContent>
              </Field>
              <Field>
                <FieldLabel>
                  <Marker variant="separator">
                    <MarkerContent>Damage</MarkerContent>
                  </Marker>
                </FieldLabel>
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
              <FieldLabel>
                <Marker variant="separator">
                  <MarkerContent>Flags</MarkerContent>
                </Marker>
              </FieldLabel>
              <FieldContent className="flex flex-row gap-x-2 flex-wrap max-w-60">
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
                  <MarkerContent>Applied Effects</MarkerContent>
                </Marker>
              </FieldLabel>
              <FieldContent className="flex flex-row flex-wrap gap-x-2 w-60!">
                {applied_effects.map((effect) => (
                  <Badge
                    variant="outline"
                    key={effect.ID}
                    className="capitalize"
                  >
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
          <div className="flex flex-col gap-4">
            <Field>
              <FieldLabel>
                <Marker variant="separator">
                  <MarkerContent>Weapon</MarkerContent>
                </Marker>
              </FieldLabel>
              <FieldContent>
                <Item variant="outline" className="p-2">
                  <ItemContent>
                    <ItemTitle>{actor.weapon?.name}</ItemTitle>
                    <ItemDescription>
                      Actions:{' '}
                      {actor.weapon?.actions
                        .map((a) => a.config.name)
                        .join(', ')}
                    </ItemDescription>
                    <ItemDescription>
                      Effects:{' '}
                      {actor.weapon?.effects.map((e) => e.name).join(', ') || (
                        <span className="italic">None</span>
                      )}
                    </ItemDescription>
                  </ItemContent>
                </Item>
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
                  <Item key={action.ID} className="p-2" asChild>
                    <Button
                      variant="outline"
                      className="h-auto"
                      disabled={action.is_disabled}
                    >
                      <ItemContent>
                        <ItemTitle>{action.config.name}</ItemTitle>
                        <ItemDescription className="text-left">
                          {action.config.description}
                        </ItemDescription>
                      </ItemContent>
                    </Button>
                  </Item>
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

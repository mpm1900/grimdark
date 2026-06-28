import { ActionContextDialog } from '#/components/action-context-dialog'
import { ActorFlag } from '#/components/actor-flag'
import { AffinityName, affinityVariants } from '#/components/affinity-name'
import {
  AffinityDamageValue,
  AffinityMultiplier,
  AffinityResistanceValue,
} from '#/components/affinity-value'
import { AppHeader } from '#/components/app-header'
import { StatValue } from '#/components/stat-value'
import { Badge } from '#/components/ui/badge'
import { Button } from '#/components/ui/button'
import { DialogTrigger } from '#/components/ui/dialog'
import { Field, FieldContent, FieldLabel } from '#/components/ui/field'
import {
  Item,
  ItemActions,
  ItemContent,
  ItemDescription,
  ItemMedia,
  ItemTitle,
} from '#/components/ui/item'
import { Marker, MarkerContent } from '#/components/ui/marker'
import { Table, TableBody, TableCell, TableRow } from '#/components/ui/table'
import { WeaponDetails } from '#/components/weapon'
import type { Actor } from '#/lib/game/actor'
import { AFFINITIES, STATS, type Stat } from '#/lib/game/core'
import { getAppliedEffects, type Game } from '#/lib/game/game'
import { RenderLog } from '#/lib/game/log'
import { gameStore } from '#/lib/stores/game'
import { cn } from '#/lib/utils'
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
                {actor.weapon && <WeaponDetails weapon={actor.weapon} />}
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
                    <Item className="p-2" asChild>
                      <DialogTrigger asChild>
                        <Button
                          variant="outline"
                          className="h-auto"
                          disabled={action.is_disabled}
                        >
                          <ItemMedia>
                            <AffinityName affinity={action.config.affinity} />
                          </ItemMedia>
                          <ItemContent>
                            <ItemTitle>{action.config.name}</ItemTitle>
                            <ItemDescription className="text-left">
                              {action.config.description || (
                                <span className="flex gap-2">
                                  {action.config.accuracy && (
                                    <span>
                                      Accuracy {action.config.accuracy * 100}%
                                    </span>
                                  )}
                                </span>
                              )}
                            </ItemDescription>
                          </ItemContent>
                          {!!action.config.power && (
                            <ItemActions>
                              <span
                                className={cn(
                                  'text-xl font-black',
                                  affinityVariants({
                                    affinity: action.config.affinity,
                                  })
                                )}
                              >
                                {action.config.power}
                              </span>
                            </ItemActions>
                          )}
                        </Button>
                      </DialogTrigger>
                    </Item>
                  </ActionContextDialog>
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

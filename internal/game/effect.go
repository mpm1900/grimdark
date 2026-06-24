package game

import "github.com/google/uuid"

const EffectPriorityBaseStats = 0
const EffectPriorityAuxStats = 0
const EffectPriorityMapBaseStats = 1
const EffectPriorityStats = 2
const EffectPriorityStages = 2
const EffectPriorityMapStages = 3

type Effect struct {
	Mutation
	ID       uuid.UUID
	Name     string
	Delay    *int
	Duration *int
	Priority int
	Triggers []Trigger
	GetLog   func(Game, Effect, Context) (Log, bool)
}
type Modifier struct {
	Bindable[Effect]
	Priority int
}

func (e Effect) Ready() bool {
	return e.Delay == nil || *e.Delay <= 0
}

func (e Effect) Filter(g Game, context Context) bool {
	if e.filter == nil {
		return true
	}

	return e.filter(g, context)
}
func (e Effect) Delta(g *Game, context Context) []uuid.UUID {
	if e.delta == nil {
		return []uuid.UUID{}
	}

	return e.delta(g, context)
}
func (e Effect) Bind(context Context) Modifier {
	bindable := bind(e, context)
	mod := Modifier{
		Bindable: bindable,
		Priority: e.Priority,
	}
	return mod
}
func (m *Modifier) Resolve(game *Game) []uuid.UUID {
	actorIDs := resolveMutation(game, m.Context, m.Payload)
	if len(actorIDs) > 0 {
		for _, actorID := range actorIDs {
			game.meta.Apply(actorID, m.Payload.ID)
		}
	}

	return actorIDs
}

func EffectSource(priority int, updater Updater[Actor]) Effect {
	return Effect{
		ID:       uuid.New(),
		Delay:    nil,
		Duration: nil,
		Priority: priority,
		Triggers: []Trigger{},
		Mutation: Mutation{
			delta: func(g *Game, context Context) []uuid.UUID {
				applied := []uuid.UUID{}
				if context.SourceID == nil {
					return applied
				}

				applied = append(applied, *context.SourceID)
				g.ModifyActor(*context.SourceID, func(a Actor) Actor {
					return updater(a, context)
				})

				return applied
			},
		},
	}
}
func EffectTargets(priority int, updater Updater[Actor]) Effect {
	return Effect{
		ID:       uuid.New(),
		Delay:    nil,
		Duration: nil,
		Priority: priority,
		Triggers: []Trigger{},
		Mutation: Mutation{
			delta: func(g *Game, context Context) []uuid.UUID {
				applied := []uuid.UUID{}

				for _, target := range g.GetTargets(context) {
					applied = append(applied, target.ID)
					g.ModifyActor(target.ID, func(a Actor) Actor {
						return updater(a, context)
					})
				}

				return applied
			},
		},
	}
}
func EffectActorsWhere(priority int, where Filter[Actor], updater Updater[Actor]) Effect {
	return Effect{
		ID:       uuid.New(),
		Delay:    nil,
		Duration: nil,
		Priority: priority,
		Triggers: []Trigger{},
		Mutation: Mutation{
			delta: func(g *Game, context Context) []uuid.UUID {
				applied := []uuid.UUID{}

				for _, target := range g.State().Actors {
					if where(target, context) {
						applied = append(applied, target.ID)
						g.ModifyActor(target.ID, func(a Actor) Actor {
							return updater(a, context)
						})
					}
				}

				return applied
			},
		},
	}
}
func EffectActorsAll(priority int, updater Updater[Actor]) Effect {
	return EffectActorsWhere(
		priority,
		func(a Actor, ctx Context) bool {
			return true
		},
		updater,
	)
}
func EffectAllies(priority int, updater Updater[Actor]) Effect {
	return Effect{
		ID:       uuid.New(),
		Delay:    nil,
		Duration: nil,
		Priority: priority,
		Triggers: []Trigger{},
		Mutation: Mutation{
			delta: func(g *Game, context Context) []uuid.UUID {
				applied := []uuid.UUID{}
				filter := func(a Actor, ctx Context) bool {
					return a.PlayerID == *context.PlayerID
				}

				for _, target := range g.State().FindActorsWhere(filter, context) {
					applied = append(applied, target.ID)
					g.ModifyActor(target.ID, func(a Actor) Actor {
						return updater(a, context)
					})
				}

				return applied
			},
		},
	}
}

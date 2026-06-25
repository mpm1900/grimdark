package game

import (
	"github.com/google/uuid"
)

const EffectPriorityAffinities = 0
const EffectPriorityBaseStats = 0
const EffectPriorityAuxStats = 0
const EffectPriorityMapBaseStats = 1
const EffectPriorityPreStageStats = 2
const EffectPriorityStages = 2
const EffectPriorityStagesOverwrite = 3
const EffectPriorityMapStages = 4
const EffectPriorityPostStagesStats = 5

var ET_DEFAULT = uuid.New()

type Effect struct {
	Mutation
	ID       uuid.UUID
	Name     string
	Delay    *int
	Duration *int
	Priority int
	Tags     map[uuid.UUID]struct{}
	Triggers []Trigger
	// check is ran on add
	Check GameFilter
	// success logs
	OnSuccess func(*Game, Effect, Context)
	// failure logs
	OnFailure func(*Game, Effect, Context)
}
type Modifier struct {
	Bindable[Effect]
	Priority int
}

func NewEffect() Effect {
	id := uuid.New()
	return Effect{
		ID:       id,
		Name:     "",
		Delay:    nil,
		Duration: nil,
		Priority: 0,
		Tags: map[uuid.UUID]struct{}{
			ET_DEFAULT: {},
			id:         {},
		},
	}
}

func (e *Effect) SetTag(tag uuid.UUID) {
	e.Tags[tag] = struct{}{}
}

func (e *Effect) SetID(id uuid.UUID) {
	old := e.ID
	delete(e.Tags, old)

	e.ID = id
	e.SetTag(id)
}

func (e Effect) Ready() bool {
	return e.Delay == nil || *e.Delay <= 0
}
func (e Effect) HasTag(tag uuid.UUID) bool {
	if e.Tags == nil {
		return false
	}

	_, ok := e.Tags[tag]
	return ok
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
func (m *Modifier) Resolve(g *Game) []uuid.UUID {
	for tag := range m.Payload.Tags {
		if _, immune := g.meta.modifier_immunities[tag]; immune {
			return []uuid.UUID{}
		}
	}

	actorIDs := resolveMutation(g, m.Context, m.Payload)
	if len(actorIDs) > 0 {
		for _, actorID := range actorIDs {
			g.meta.apply(actorID, m.Payload.ID)
		}
	}

	return actorIDs
}

func EffectSource(priority int, updater Updater[Actor]) Effect {
	effect := NewEffect()
	effect.Priority = priority
	effect.Mutation = Mutation{
		delta: func(g *Game, context Context) []uuid.UUID {
			applied := []uuid.UUID{}
			if context.SourceID == uuid.Nil {
				return applied
			}

			applied = append(applied, context.SourceID)
			g.ModifyActor(context.SourceID, func(a Actor) Actor {
				return updater(*g, a, context)
			})

			return applied
		},
	}

	return effect
}
func EffectTargets(priority int, updater Updater[Actor]) Effect {
	effect := NewEffect()
	effect.Priority = priority
	effect.Mutation = Mutation{
		delta: func(g *Game, context Context) []uuid.UUID {
			applied := []uuid.UUID{}

			for _, target := range g.GetTargets(context) {
				applied = append(applied, target.ID)
				g.ModifyActor(target.ID, func(a Actor) Actor {
					return updater(*g, a, context)
				})
			}

			return applied
		},
	}

	return effect
}
func EffectActorsWhere(priority int, where Filter[Actor], updater Updater[Actor]) Effect {
	effect := NewEffect()
	effect.Priority = priority
	effect.Mutation = Mutation{
		delta: func(g *Game, context Context) []uuid.UUID {
			applied := []uuid.UUID{}

			actors := g.FindActors(where, context)
			for _, target := range actors {
				applied = append(applied, target.ID)
				g.ModifyActor(target.ID, func(a Actor) Actor {
					return updater(*g, a, context)
				})
			}

			return applied
		},
	}

	return effect
}
func EffectActorsAll(priority int, updater Updater[Actor]) Effect {
	return EffectActorsWhere(
		priority,
		AllActors,
		updater,
	)
}
func EffectActorsActive(priority int, updater Updater[Actor]) Effect {
	return EffectActorsWhere(
		priority,
		CombineFilters(ActiveActors, AliveActors),
		updater,
	)
}
func EffectAllies(priority int, updater Updater[Actor]) Effect {
	return EffectActorsWhere(
		priority,
		CombineFilters(Allies),
		updater,
	)
}

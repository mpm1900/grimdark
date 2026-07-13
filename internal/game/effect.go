package game

import (
	"github.com/google/uuid"
)

const EffectPriorityAffinities = 0
const EffectPriorityBaseStats = 0
const EffectPriorityOffsetStats = 0
const EffectPriorityFlags = 0
const EffectPriorityTriggers = 0
const EffectPriorityOffsetOverwrite = 1
const EffectPriorityMapBaseStats = 2
const EffectPriorityPreStageStats = 3
const EffectPriorityStages = 3
const EffectPriorityStagesOverwrite = 4
const EffectPriorityMapStages = 5
const EffectPriorityPostStagesStats = 6
const EffectPriorityActionState = 6

var ET_DEFAULT = uuid.New()

type EffectState struct {
	Delay    *int
	Duration *int
}

type Effect struct {
	Mutation
	ID       uuid.UUID              `json:"ID"`
	Name     string                 `json:"name"`
	Delay    *int                   `json:"delay"`
	Duration *int                   `json:"duration"`
	Priority int                    `json:"priority"`
	Tags     map[uuid.UUID]struct{} `json:"-"`
	Triggers []Trigger              `json:"-"`
	// check is ran on add
	Check GameFilter `json:"-"`
	// success logs
	CheckSuccess func(*Game, Effect, Context) `json:"-"`
	// failure logs
	CheckFailure func(*Game, Effect, Context) `json:"-"`
}
type Modifier struct {
	Bindable[Effect]
	Priority int `json:"priority"`
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
func (e *Effect) ApplyState(s EffectState) {
	if s.Delay != nil {
		e.Delay = s.Delay
	}
	if s.Duration != nil {
		e.Duration = s.Duration
	}
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
			g.meta.apply(m.ID, actorID)
		}
	} else {
		g.meta.apply(m.ID, uuid.Nil)
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
func EffectActorsActiveOther(priority int, updater Updater[Actor]) Effect {
	return EffectActorsWhere(
		priority,
		OtherActors,
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

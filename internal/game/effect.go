package game

import (
	"github.com/google/uuid"
)

const EffectPriorityImmunities = -1
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

type EffectState struct {
	Delay    *int
	Duration *int
}

type Effect struct {
	Mutation
	ID          uuid.UUID           `json:"ID"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Delay       *int                `json:"delay"`
	Duration    *int                `json:"duration"`
	Priority    int                 `json:"priority"`
	Tags        map[string]struct{} `json:"-"`
	Triggers    []Trigger           `json:"-"`
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
	tags := map[string]struct{}{}
	tags[id.String()] = struct{}{}

	return Effect{
		ID:          id,
		Name:        "",
		Description: "",
		Delay:       nil,
		Duration:    nil,
		Priority:    0,
		Tags:        tags,
	}
}

func (e *Effect) SetTag(tag string) {
	e.Tags[tag] = struct{}{}
}

func (e *Effect) SetID(id uuid.UUID) {
	old := e.ID
	delete(e.Tags, old.String())

	e.ID = id
	e.SetTag(id.String())
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
func (e Effect) HasTag(tag string) bool {
	if e.Tags == nil {
		return false
	}

	_, ok := e.Tags[tag]
	return ok
}
func (e Effect) Clone() Effect {
	if e.Delay != nil {
		delay := *e.Delay
		e.Delay = &delay
	}
	if e.Duration != nil {
		duration := *e.Duration
		e.Duration = &duration
	}
	if e.Tags != nil {
		tags := make(map[string]struct{}, len(e.Tags))
		for tag := range e.Tags {
			tags[tag] = struct{}{}
		}
		e.Tags = tags
	}

	return e
}
func (e Effect) Filter(g *Game, context Context) bool {
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
	context = context.Clone()
	e = e.Clone()
	context.EffectID = e.ID
	bindable := bind(e, context)
	mod := Modifier{
		Bindable: bindable,
		Priority: e.Priority,
	}
	return mod
}
func (m *Modifier) Resolve(g *Game) []uuid.UUID {
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

			source, ok := g.GetSource(context)
			if !ok {
				return applied
			}
			if source.HasEffectImmunity(context.EffectID) {
				return applied
			}

			applied = append(applied, context.SourceID)
			g.ModifyActor(context.SourceID, func(a Actor) Actor {
				return updater(g, a, context)
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
				if target.HasEffectImmunity(context.EffectID) {
					continue
				}

				applied = append(applied, target.ID)
				g.ModifyActor(target.ID, func(a Actor) Actor {
					return updater(g, a, context)
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
				if target.HasEffectImmunity(context.EffectID) {
					continue
				}

				applied = append(applied, target.ID)
				g.ModifyActor(target.ID, func(a Actor) Actor {
					return updater(g, a, context)
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

package game

import (
	"github.com/google/uuid"
)

type Bindable[P any] struct {
	ID      uuid.UUID
	Context Context
	Payload P
}

type Mutation struct {
	filter Filter[Game]
	delta  Mutator
}
type Transaction Bindable[Mutation]

func bind[P any](payload P, context Context) Bindable[P] {
	return Bindable[P]{
		ID:      uuid.New(),
		Context: context,
		Payload: payload,
	}
}

func NewMutation(filter Filter[Game], delta Mutator) Mutation {
	return Mutation{
		filter,
		delta,
	}
}

func (m *Mutation) SetFilter(filter Filter[Game]) {
	m.filter = filter
}
func (m Mutation) Filter(g Game, context Context) bool {
	if m.filter == nil {
		return true
	}

	return m.filter(g, context)
}

func (m Mutation) Delta(g *Game, context Context) []uuid.UUID {
	if m.delta == nil {
		return []uuid.UUID{}
	}

	return m.delta(g, context)
}

func (m Mutation) Bind(context Context) Transaction {
	return Transaction(bind(m, context))
}

type resolvableMutation interface {
	Filter(Game, Context) bool
	Delta(*Game, Context) []uuid.UUID
}

func resolveMutation(g *Game, context Context, mutation resolvableMutation) []uuid.UUID {
	if !mutation.Filter(*g, context) {
		return []uuid.UUID{}
	}

	return mutation.Delta(g, context)
}
func (tx *Transaction) Resolve(g *Game) []uuid.UUID {
	return resolveMutation(g, tx.Context, tx.Payload)
}

func AddModifiers(modifiers ...Modifier) Mutation {
	return Mutation{
		delta: func(g *Game, ctx Context) []uuid.UUID {
			g.AddModifiers(modifiers...)
			return []uuid.UUID{}
		},
	}
}
func PushLog(log Log) Mutation {
	return Mutation{
		delta: func(g *Game, ctx Context) []uuid.UUID {
			g.PushLog(log.Bind(ctx))
			return []uuid.UUID{}
		},
	}
}
func MutateSource(updater Updater[Actor]) Mutation {
	return Mutation{
		delta: func(g *Game, context Context) []uuid.UUID {
			applied := []uuid.UUID{}
			if context.SourceID == nil {
				return applied
			}

			applied = append(applied, *context.SourceID)
			g.MutateActor(*context.SourceID, func(a Actor) Actor {
				return updater(a, context)
			})

			return applied
		},
	}
}
func MutateTargets(updater Updater[Actor]) Mutation {
	return Mutation{
		delta: func(g *Game, context Context) []uuid.UUID {
			applied := []uuid.UUID{}

			for _, target := range g.GetTargets(context) {
				applied = append(applied, target.ID)
				g.MutateActor(target.ID, func(a Actor) Actor {
					return updater(a, context)
				})
			}

			return applied
		},
	}
}
func DamageTargets(damage float64) Mutation {
	return Mutation{
		delta: func(g *Game, context Context) []uuid.UUID {
			applied := []uuid.UUID{}

			g.DamageTargets(context, damage)
			for _, target := range g.GetTargets(context) {
				applied = append(applied, target.ID)
			}

			return applied
		},
	}
}

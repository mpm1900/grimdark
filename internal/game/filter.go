package game

import (
	"slices"

	"github.com/google/uuid"
)

type Filter[T any] func(T, Context) bool
type Mutator func(*Game, Context) []uuid.UUID
type Updater[T any] func(Game, T, Context) T

func CombineFilters[T any](
	filters ...Filter[T],
) Filter[T] {
	return func(t T, ctx Context) bool {
		for _, fn := range filters {
			if !fn(t, ctx) {
				return false
			}
		}

		return true
	}
}

// filters
// game
func TrueGameFilter(g Game, context Context) bool {
	return true
}

// context
func ContextTargetLength(length int) Filter[Game] {
	return func(g Game, ctx Context) bool {
		targets := g.GetTargets(ctx)
		return len(targets) == length
	}
}

// actor
func NoneActors(actor Actor, context Context) bool {
	return false
}
func ActiveActors(actor Actor, context Context) bool {
	return actor.PositionID != uuid.Nil
}
func AliveActors(actor Actor, context Context) bool {
	return actor.IsAlive
}
func Allies(actor Actor, context Context) bool {
	return actor.PlayerID == context.PlayerID
}
func Enemies(actor Actor, context Context) bool {
	return actor.PlayerID != context.PlayerID
}
func SourceActor(actor Actor, context Context) bool {
	return actor.ID == context.SourceID
}
func TargetActors(actor Actor, context Context) bool {
	is_actor := slices.Contains(context.ActorIDs, actor.ID)
	is_pos := actor.PositionID != uuid.Nil && slices.Contains(context.PositionIDs, actor.PositionID)
	return is_actor || is_pos
}

// trigger validation filters
func TriggerTargetMatchesModifierParent(g Game, trigger Context, modifier Context) bool {
	for _, target := range g.GetTargets(trigger) {
		if target.ID == modifier.ParentID {
			return true
		}
	}

	return false
}

package game

import (
	"slices"

	"github.com/google/uuid"
)

type Filter[T any] func(Game, T, Context) bool
type Updater[T any] func(Game, T, Context) T

type GameFilter func(Game, Context) bool
type Mutator func(*Game, Context) []uuid.UUID

func CombineFilters[T any](
	filters ...Filter[T],
) Filter[T] {
	return func(g Game, t T, ctx Context) bool {
		for _, fn := range filters {
			if !fn(g, t, ctx) {
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
func ContextTargetLength(length int) GameFilter {
	return func(g Game, ctx Context) bool {
		targets := g.GetTargets(ctx)
		return len(targets) == length
	}
}

// actor
func NoneActors(g Game, actor Actor, context Context) bool {
	return false
}
func AllActors(g Game, actor Actor, context Context) bool {
	return true
}
func ActiveActors(g Game, actor Actor, context Context) bool {
	return actor.IsActive()
}
func AliveActors(g Game, actor Actor, context Context) bool {
	return actor.IsAlive
}
func Allies(g Game, actor Actor, context Context) bool {
	return actor.PlayerID == context.PlayerID
}
func Enemies(g Game, actor Actor, context Context) bool {
	return actor.PlayerID != context.PlayerID
}
func SourceActor(g Game, actor Actor, context Context) bool {
	return actor.ID == context.SourceID
}
func TargetActors(g Game, actor Actor, context Context) bool {
	is_actor := slices.Contains(context.ActorIDs, actor.ID)
	is_pos := actor.IsActive() && slices.Contains(context.PositionIDs, actor.PositionID)
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

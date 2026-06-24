package game

import (
	"slices"

	"github.com/google/uuid"
)

type Filter[T any] func(T, Context) bool
type Mutator func(*Game, Context) []uuid.UUID
type Updater[T any] func(T, Context) T

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
// actor
func ActiveActors(actor Actor, context Context) bool {
	return actor.PositionID != nil
}
func AliveActors(actor Actor, context Context) bool {
	return actor.IsAlive
}
func Allies(actor Actor, context Context) bool {
	if context.PlayerID == nil {
		return false
	}

	return actor.PlayerID == *context.PlayerID
}
func Enemies(actor Actor, context Context) bool {
	if context.PlayerID == nil {
		return false
	}

	return actor.PlayerID != *context.PlayerID
}
func SourceActor(actor Actor, context Context) bool {
	if context.SourceID == nil {
		return false
	}

	return actor.ID == *context.SourceID
}
func TargetActors(actor Actor, context Context) bool {
	is_actor := slices.Contains(context.ActorIDs, actor.ID)
	is_pos := actor.PositionID != nil && slices.Contains(context.PositionIDs, *actor.PositionID)
	return is_actor || is_pos
}

// trigger validation filters
func TriggerTargetMatchesModifierParent(g Game, trigger Context, modifier Context) bool {
	if modifier.ParentID == nil {
		return false
	}

	for _, target := range g.GetTargets(trigger) {
		if target.ID == *modifier.ParentID {
			return true
		}
	}

	return false
}

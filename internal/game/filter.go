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

// context validators
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
func InactiveActors(g Game, actor Actor, context Context) bool {
	return !actor.IsActive()
}
func AliveActors(g Game, actor Actor, context Context) bool {
	return actor.IsAlive
}
func DeadActors(g Game, actor Actor, context Context) bool {
	return !actor.IsAlive
}
func NonStunnedActors(g Game, actor Actor, context Context) bool {
	return !actor.IsStunned
}
func Allies(g Game, actor Actor, context Context) bool {
	return actor.PlayerID == context.PlayerID
}
func OtherAllies(g Game, actor Actor, context Context) bool {
	return Allies(g, actor, context) && NotSourceActor(g, actor, context)
}
func Enemies(g Game, actor Actor, context Context) bool {
	return actor.PlayerID != context.PlayerID
}
func SourceActor(g Game, actor Actor, context Context) bool {
	return actor.ID == context.SourceID
}
func NotSourceActor(g Game, actor Actor, context Context) bool {
	return actor.ID != context.SourceID
}
func OtherActors(g Game, actor Actor, context Context) bool {
	return CombineFilters(ActiveActors, AliveActors, NotSourceActor)(g, actor, context)
}
func TargetActors(g Game, actor Actor, context Context) bool {
	is_actor := slices.Contains(context.ActorIDs, actor.ID)
	is_pos := actor.IsActive() && slices.Contains(context.PositionIDs, actor.PositionID)
	return is_actor || is_pos
}
func PositionRank(rank int) Filter[Actor] {
	return func(g Game, a Actor, ctx Context) bool {
		position, ok := g.GetPosition(a.PositionID)
		if !ok {
			return false
		}

		return position.Rank == rank
	}
}
func ActionRange(action_range int) Filter[Actor] {
	return func(g Game, target Actor, ctx Context) bool {
		source, ok := g.GetSource(ctx)
		if !ok {
			return false
		}

		source_pos, ok := g.GetPosition(source.PositionID)
		if !ok {
			return false
		}
		target_pos, ok := g.GetPosition(target.PositionID)
		if !ok {
			return false
		}

		distance := source_pos.GetDistanceFrom(target_pos)
		state := source.ActionsState[ctx.ActionID]
		return action_range+state.RangeBonus >= distance
	}
}

// trigger validation filters
func TriggerModifierParentIsActive(g Game, trigger Context, modifier Context) bool {
	parent, ok := g.GetParent(modifier)
	if !ok {
		return false
	}

	return parent.IsActive()
}
func TriggerTargetMatchesModifierParent(g Game, trigger Context, modifier Context) bool {
	for _, target := range g.GetTargets(trigger) {
		if target.ID == modifier.ParentID {
			return true
		}
	}

	return false
}
func TriggerSourceMatchesModifierParent(g Game, trigger Context, modifier Context) bool {
	return trigger.SourceID == modifier.ParentID
}

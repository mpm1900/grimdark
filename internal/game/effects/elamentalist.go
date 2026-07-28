package effects

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func Elamentalist() game.Effect {
	effect := game.EffectActorsWhere(
		game.EffectPriorityFlags,
		func(g *game.Game, a game.Actor, ctx game.Context) bool {
			active_context := g.State().ActiveContext
			if active_context == nil || active_context.ActionID == uuid.Nil {
				return false
			}

			if active_context.SourceID != ctx.ParentID {
				return false
			}

			return true
		},
		func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
			active_context := g.State().ActiveContext
			if active_context == nil || active_context.ActionID == uuid.Nil {
				return a
			}

			if active_context.SourceID != a.ID {
				return a
			}

			action, ok := a.GetActionByID(active_context.ActionID)
			if !ok {
				return a
			}

			a.Class.Affinities = game.Set[game.Affinity]{
				action.Config.Affinity: {},
			}

			return a
		},
	)

	effect.Name = "Elementalist"
	effect.Description = "During an action, the user's affinity became that of the chosen action."
	return effect
}

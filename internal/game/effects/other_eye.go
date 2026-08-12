package effects

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var OtherEye = otherEye()

func otherEye() game.Effect {
	effect := StagesResetWhere(func(g *game.Game, a game.Actor, ctx game.Context) bool {
		active_context := g.State().ActiveContext
		if active_context == nil || active_context.ActionID == uuid.Nil {
			return false
		}

		if active_context.SourceID != ctx.ParentID {
			return false
		}

		return active_context.HasTarget(a)
	})
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Other Eye",
		"While attacking, this actor ignores all stat stage changes to targeted actors.",
	)

	return effect
}

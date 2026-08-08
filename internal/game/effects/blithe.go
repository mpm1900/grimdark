package effects

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func Blithe() game.Effect {
	effect := StagesResetWhere(func(g *game.Game, a game.Actor, ctx game.Context) bool {
		parent, ok := g.GetParent(ctx)
		if !ok {
			return false
		}

		active_ctx := g.State().ActiveContext
		if active_ctx == nil {
			return false
		}

		// if the parent (owner) of this effect is the source,
		// reset the stats of the target if they match
		if active_ctx.SourceID == parent.ID {
			return active_ctx.HasTarget(a)
		}

		// if the parent (owner) of this effect is a target,
		// reset the stats of the source if they match
		if active_ctx.HasTarget(parent) {
			return active_ctx.SourceID == a.ID
		}

		return false
	})
	effect.ID = uuid.MustParse("019fe205-d2fb-751c-a057-f638c9578d0c")
	effect.Name = "Blithe"
	effect.Description = "When this actor attacks, reset the stat stages of targets. When this actor is attacked, reset the stat changes of the source actor."

	return effect
}

package effects

import (
	"grimdark/internal/game"
	"slices"
)

var Immobile = immobile()

func immobile() game.Effect {
	effect := game.EffectSource(game.EffectPriorityActionState, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		for _, action := range a.GetActions() {
			if slices.Contains(action.Tags, game.ATMovement) {
				a.UpdateActionState(action.ID, func(as game.ActionState) game.ActionState {
					as.IsDisabled = true
					return as
				})
			}
		}

		return a
	})
	effect.Name = "Immobile"
	effect.Description = "Cannot use movement actions."
	effect.CheckSuccess = game.EffectGainTargetsOnSuccess

	return effect
}

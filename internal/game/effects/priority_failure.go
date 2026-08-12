package effects

import (
	"grimdark/internal/game"
	"slices"
)

func PriorityFailure() game.Effect {
	effect := game.EffectActorsActive(game.EffectPriorityActionState, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		for _, action := range a.GetActions() {
			is_system := slices.Contains(action.Tags, game.ATSystem)
			if action.Config.Power > 0 && !is_system {
				state := a.ActionStates[action.ID]
				priority := action.Config.Priority + state.PriorityBonus
				if priority > 0 {
					a.UpdateActionState(action.ID, func(as game.ActionState) game.ActionState {
						as.IsDisabled = true
						return as
					})
				}
			}
		}

		return a
	})
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Priority Failure",
		"All actions with greater than priority 0 are disabled.",
	)

	return effect
}

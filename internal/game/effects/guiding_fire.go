package effects

import "grimdark/internal/game"

func GuidingFire() game.Effect {
	effect := game.EffectParent(game.EffectPriorityActionState, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		for _, action := range a.GetActions() {
			if action.Config.Affinity != game.Fire {
				continue
			}

			a.UpdateActionState(action.ID, func(as game.ActionState) game.ActionState {
				as.BypassAccuracy = true
				return as
			})
		}
		return a
	})
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Guiding Fire",
		"All fire attacks bypass accuracy checks.",
	)

	return effect
}

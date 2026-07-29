package effects

import "grimdark/internal/game"

func Devout() game.Effect {
	effect := game.EffectSource(
		game.EffectPriorityActions,
		func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
			actions := a.GetActions()
			for _, action := range actions {
				a.UpdateActionState(action.ID, func(as game.ActionState) game.ActionState {
					if action.Config.Power > 0 {
						return as
					}

					as.PriorityBonus += 1
					return as
				})
			}

			return a
		},
	)
	effect.Name = "Devout"
	effect.Description = "Non-damaging actions have +1 priority."
	effect.CheckSuccess = game.EffectGainTargetsOnSuccess

	return effect
}

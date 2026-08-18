package effects

import "grimdark/internal/game"

func DivineBlessing() game.Effect {
	effect := game.EffectParent(game.EffectPriorityPostStagesStats, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		if a.Stacks[game.Wounds] != 0 {
			return a
		}

		a.Stats[game.MartialDefense] *= 2
		a.Stats[game.SpecialDefense] *= 2
		return a
	})
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Divine Blessing",
		"Doubles defenses, but only at full health.",
	)
	return effect
}

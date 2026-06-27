package effects

import (
	"grimdark/internal/game"
)

func StaggerTargets() game.Effect {
	effect := game.EffectTargets(
		game.EffectPriorityFlags,
		func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
			a.IsStaggered = true
			return a
		},
	)
	effect.Name = "Staggered"
	effect.Duration = game.P(1)

	return effect
}

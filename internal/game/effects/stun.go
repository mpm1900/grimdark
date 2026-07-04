package effects

import (
	"grimdark/internal/game"
)

func StunTargets() game.Effect {
	effect := game.EffectTargets(
		game.EffectPriorityFlags,
		func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
			a.IsStunned = true
			return a
		},
	)
	effect.Name = "Stunned"
	effect.Duration = game.P(1)

	return effect
}

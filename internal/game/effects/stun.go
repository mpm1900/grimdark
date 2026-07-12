package effects

import (
	"grimdark/internal/game"
)

var StaggerTargets = stunTargets(1)
var StunTargets = stunTargets(2)

func stunTargets(duration int) game.Effect {
	effect := game.EffectTargets(
		game.EffectPriorityFlags,
		func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
			a.IsStunned = true
			return a
		},
	)
	effect.Name = "Stunned"
	effect.Duration = game.P(duration)
	effect.CheckSuccess = EffectGainTargetsOnSuccess

	return effect
}

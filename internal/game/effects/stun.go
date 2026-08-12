package effects

import (
	"grimdark/internal/game"
)

var StaggerTargets = stunTargets(1)
var StunTargets = stunTargets(2)

func stunTargets(duration int) game.Effect {
	effect := game.EffectTargets(
		game.EffectPriorityFlags,
		func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
			a.IsDisabled = true
			return a
		},
	)
	effect.Entity = game.MakeEntity(
		effect.ID,
		"Stunned",
		"Cannot act.",
	)
	effect.Duration = game.P(duration)
	effect.CheckSuccess = game.EffectGainTargetsOnSuccess

	return effect
}

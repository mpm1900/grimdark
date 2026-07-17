package effects

import "grimdark/internal/game"

func ProtectedSource() game.Effect {
	effect := game.EffectSource(game.EffectPriorityFlags, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.IsProtected = true

		return a
	})
	effect.Name = "Protected"

	return effect
}

func ProtectedWhere(where game.Filter[game.Actor]) game.Effect {
	effect := game.EffectActorsWhere(game.EffectPriorityFlags, where, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.IsProtected = true

		return a
	})
	effect.Name = "Protected"

	return effect
}

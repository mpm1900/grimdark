package effects

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func protectedEntity(id uuid.UUID) game.Entity {
	return game.MakeEntity(
		id,
		"Protected",
		"Protected from attacks and actions.",
	)
}

func ProtectedSource() game.Effect {
	effect := game.EffectSource(game.EffectPriorityFlags, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.IsProtected = true

		return a
	})
	effect.Entity = protectedEntity(effect.ID)
	effect.Duration = game.P(1)
	effect.CheckSuccess = game.EffectGainSourceOnSuccess

	return effect
}

func ProtectedWhere(where game.Filter[game.Actor]) game.Effect {
	effect := game.EffectActorsWhere(game.EffectPriorityFlags, where, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.IsProtected = true

		return a
	})
	effect.Entity = protectedEntity(effect.ID)
	effect.Duration = game.P(1)
	effect.CheckSuccess = game.EffectGainSourceOnSuccess

	return effect
}

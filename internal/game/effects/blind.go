package effects

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func blindEntity(id uuid.UUID) game.Entity {
	return game.MakeEntity(
		id,
		"Blind",
		"All accuracy checks will fail.",
	)
}

func BlindParent() game.Effect {
	effect := game.EffectParent(game.EffectPriorityFlags, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.IsBlind = true

		return a
	})
	effect.Entity = blindEntity(effect.ID)
	effect.CheckSuccess = game.EffectGainSourceOnSuccess

	return effect
}

func BlindWhere(where game.Filter[game.Actor]) game.Effect {
	effect := game.EffectActorsWhere(game.EffectPriorityFlags, where, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.IsBlind = true

		return a
	})
	effect.Entity = blindEntity(effect.ID)
	effect.CheckSuccess = game.EffectGainSourceOnSuccess

	return effect
}

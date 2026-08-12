package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func arcaneRitualEntity(id uuid.UUID) game.Entity {
	return game.MakeEntity(
		id,
		"Arcane Ritual",
		"Inverts the Speed of all active actors for 5 turns.",
	)
}

func arcaneRitual() game.Effect {
	effect := game.EffectActorsAll(game.EffectPriorityNegations, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stats[game.Speed] *= -1
		return a
	})
	effect.Entity = arcaneRitualEntity(effect.ID)
	effect.Duration = game.P(5)
	return effect
}

func ArcaneRitual() game.Action {
	id := uuid.MustParse("019fc0d5-abdf-7117-a0dd-d1d32e418c21")

	return game.Action{
		ID:     id,
		Tags:   []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: arcaneRitualEntity(id),
		Config: game.ActionConfig{
			Affinity:    game.Arcane,
			TargetCount: 0,
			Cooldown:    5,
		},
		Resolve: game.AddGlobalEffects(
			game.StatusConfig{},
			1,
			arcaneRitual(),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
	}
}

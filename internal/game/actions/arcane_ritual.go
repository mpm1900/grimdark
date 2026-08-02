package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func arcaneRitual() game.Effect {
	effect := game.EffectActorsAll(game.EffectPriorityNegations, func(g *game.Game, a game.Actor, ctx game.Context) game.Actor {
		a.Stats[game.Speed] *= -1
		return a
	})
	effect.Name = "Arcane Ritual"
	effect.Description = "Speed is inverted."
	effect.Duration = game.P(5)
	return effect
}

var ArcaneRitual = game.Action{
	ID:   uuid.MustParse("019fc0d5-abdf-7117-a0dd-d1d32e418c21"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Arcane Ritual",
		Description: "Inverts the Speed of all active actors for 5 turns.",
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

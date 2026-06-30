package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var Slash = game.Action{
	ID: uuid.MustParse("019f0aec-8b34-72cc-bbcc-36350e9fa6fb"),
	Config: game.ActionConfig{
		Name:         "Slash",
		Description: "Slashes target and possibly applies Speed Down.",
		Affinity:     game.Psychic,
		Stat:         game.Melee,
		Accuracy:     game.P(0.90),
		Power:        70,
		Lifesteal:    0.12,
		Hits:         1,
		CritChance:   0,
		CritModifier: 1.5,
	},
	IsActive: true,
	Resolve: game.BasicAttack(game.AttackConfig{
		OnSuccessResult: func(g game.Game, context game.Context, this *game.ActionContext, result game.DamageResult) {
			game.AddResultEffects(
				0.5,
				effects.StatDownTargets(game.Speed, 1),
			)(g, context, this, result)
			game.AddResultEffects(
				0.5,
				effects.StaggerTargets(),
			)(g, context, this, result)
		},
	}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Allies),
}

package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var PinDown = game.Action{
	ID:   uuid.MustParse("019fafb3-be68-75e4-b32c-0c02f8564b78"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Pin Down",
		Description:  "Hits 3 times. If all 3 hits land, then this action lowers the target's Speed.",
		Affinity:     game.Physical,
		Stat:         game.Ranged,
		Accuracy:     game.P(0.8),
		Power:        20,
		Lifesteal:    0,
		Hits:         3,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Range:        game.P(3),
	},
	Resolve: game.MakeAttack(game.AttackConfig{
		OnSuccess: func(g *game.Game, context game.Context, this *game.ActionContext) {
			targets := g.GetTargets(context)
			for _, t := range targets {
				game.AddEffectsTarget(
					1.0,
					t,
					effects.StatDownTargets(game.Speed, 1),
				)(g, context, this)
			}
		},
	}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(3)),
}

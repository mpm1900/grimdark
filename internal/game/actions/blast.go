package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var Blast = game.Action{
	ID:   uuid.MustParse("019f287e-fdf6-7fc4-87b0-0a0060efc424"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Blast",
		Description:  "A blast of arcane energy.",
		Affinity:     game.Arcane,
		Stat:         game.Special,
		Accuracy:     game.P(1.0),
		Power:        80,
		Lifesteal:    0,
		Hits:         1,
		CritChance:   0,
		CritModifier: 1.5,
		TargetCount:  1,
	},
	Resolve: game.MakeAttack(game.AttackConfig{
		OnSuccessResult: func(g *game.Game, context game.Context, this *game.ActionContext, result game.DamageResult) {

		},
	}),
	MapContext:       game.CtxTargetPreCollateral(),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Enemies, game.PositionRank(2)),
}

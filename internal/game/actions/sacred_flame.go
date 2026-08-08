package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var SacredFlame = game.Action{
	ID:   uuid.MustParse("019f9c82-435c-7f98-83b9-fed6b1f23ec4"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Sacred Flame",
		Description:  "50% chance to apply Burned target. This action is not usable from 1st position.",
		Affinity:     game.Fire,
		Stat:         game.Special,
		Accuracy:     game.P(0.7),
		Power:        120,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
	},
	Resolve: game.MakeAttack(game.AttackConfig{
		OnSuccessResult: func(g *game.Game, context game.Context, this *game.ActionContext, result game.DamageResult) {
			game.AddResultEffects(
				0.5,
				effects.Burned(),
			)(g, context, this, result)
		},
	}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors),
	DisabledCheck: func(g *game.Game, source game.Actor) bool {
		position, ok := g.GetPosition(source.PositionID)
		if !ok {
			return true
		}

		return position.Rank == 0
	},
}

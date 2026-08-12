package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func CollateralShot() game.Action {
	id := uuid.MustParse("019f8624-d636-7f7e-9a7a-31bb0f0fa529")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Collateral Shot",
			"Damages all enemy actors positioned in front of the target. This action is only usable from 2nd or 3rd position.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Ranged,
			Accuracy:     game.P(0.80),
			Power:        85,
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  1,
		},
		Resolve:          game.MakeAttack(game.AttackConfig{}),
		MapContext:       game.CtxTargetPreCollateral(),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Enemies),
		DisabledCheck: func(g *game.Game, source game.Actor) bool {
			position, ok := g.GetPosition(source.PositionID)
			if !ok {
				return true
			}

			return position.Rank == 0
		},
	}
}

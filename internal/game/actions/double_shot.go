package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func DoubleShot() game.Action {
	id := uuid.MustParse("019faf7f-db18-73c9-bc20-7dcbd86240e1")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Double Shot",
			"Hits 2 targets.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Ranged,
			Accuracy:     game.P(0.80),
			Power:        50,
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  2,
			Range:        game.P(3),
		},
		Resolve:          game.MakeAttack(game.AttackConfig{}),
		ValidateContext:  game.ContextTargetLength(2),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(3)),
	}
}

package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var DoubleShot = game.Action{
	ID:   uuid.MustParse("019faf7f-db18-73c9-bc20-7dcbd86240e1"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Double Shot",
		Description:  "Hits 2 targets.",
		Affinity:     game.Physical,
		Stat:         game.Ranged,
		Accuracy:     game.P(0.80),
		Power:        50,
		Lifesteal:    0,
		Hits:         1,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  2,
		Range:        game.P(3),
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	ValidateContext:  game.ContextTargetLength(2),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(3)),
}

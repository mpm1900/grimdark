package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var BurstFire = game.Action{
	ID:   uuid.MustParse("019f8ff1-372f-7631-9d9d-16663cb3e108"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Burst Fire",
		Description:  "Hits 8 times, but has low accuracy.",
		Affinity:     game.Physical,
		Stat:         game.Ranged,
		Accuracy:     game.P(0.40),
		Power:        20,
		Repeats:      7,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Range:        game.P(3),
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(3)),
}

package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

var FiftyFifty = game.Action{
	ID:   uuid.MustParse("019f91fa-3962-7986-838c-95308827ae2a"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:         "Fifty Fifty",
		Description:  "Hits 2 times, but has low accuracy.",
		Affinity:     game.Kinetic,
		Stat:         game.Ranged,
		Accuracy:     game.P(0.50),
		Power:        72,
		Lifesteal:    0,
		Hits:         2,
		CritStage:    0,
		CritModifier: 1.5,
		TargetCount:  1,
		Range:        game.P(3),
	},
	Resolve:          game.MakeAttack(game.AttackConfig{}),
	ValidateContext:  game.ContextTargetLength(1),
	TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(3)),
}

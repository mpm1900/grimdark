package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func FiftyFifty() game.Action {
	id := uuid.MustParse("019f91fa-3962-7986-838c-95308827ae2a")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Fifty Fifty",
			"Hits 2 times, but has low accuracy.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Ranged,
			Accuracy:     game.P(0.50),
			Power:        75,
			Repeats:      1,
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  1,
			Range:        game.P(3),
		},
		Resolve:          game.MakeAttack(game.AttackConfig{}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.OtherActors, game.ActionRange(3)),
	}
}

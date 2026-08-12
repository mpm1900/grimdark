package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func Poke() game.Action {
	id := uuid.MustParse("019fb005-52e8-7224-95ca-b6aae73e4723")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Poke",
			"",
		),
		Config: game.ActionConfig{
			Affinity:     game.Physical,
			Stat:         game.Melee,
			Accuracy:     game.P(0.90),
			Power:        80,
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  1,
			Range:        game.P(2),
		},
		Resolve:          game.MakeAttack(game.AttackConfig{}),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.NotSourceActor, game.ActionRange(2)),
	}
}

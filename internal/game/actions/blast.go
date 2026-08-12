package actions

import (
	"grimdark/internal/game"

	"github.com/google/uuid"
)

func Blast() game.Action {
	id := uuid.MustParse("019f287e-fdf6-7fc4-87b0-0a0060efc424")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Blast",
			"A blast of arcane energy.",
		),
		Config: game.ActionConfig{
			Affinity:     game.Arcane,
			Stat:         game.Special,
			Accuracy:     game.P(1.0),
			Power:        80,
			CritStage:    0,
			CritModifier: 1.5,
			TargetCount:  1,
		},
		Resolve:          game.MakeAttack(game.AttackConfig{}),
		MapContext:       game.CtxTargetPreCollateral(),
		ValidateContext:  game.ContextTargetLength(1),
		TargetsPredicate: game.CombineFilters(game.ActiveActors, game.Enemies, game.PositionRank(2)),
	}
}

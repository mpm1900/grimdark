package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func LockOn() game.Action {
	id := uuid.MustParse("019f87b5-9d77-74d1-a6dc-57b031b0f124")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Lock On",
			"Raises users's ranged and accuracy stats by 1 stage.",
		),
		Config: game.ActionConfig{
			Affinity:    game.Physical,
			TargetCount: 0,
		},
		Resolve: game.AddSourceEffects(
			game.StatusConfig{},
			1,
			effects.StatUpSource(game.Accuracy, 1),
			effects.StatUpSource(game.Ranged, 1),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
	}
}

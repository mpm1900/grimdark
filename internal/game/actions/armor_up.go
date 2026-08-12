package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func ArmorUp() game.Action {
	id := uuid.MustParse("019fceb1-2da1-739b-9394-01899dfd8e2f")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Armor Up",
			"Raises user's Martial Defense stat.",
		),
		Config: game.ActionConfig{
			Affinity:    game.Physical,
			TargetCount: 0,
		},
		Resolve: game.AddSourceEffects(
			game.StatusConfig{},
			1,
			effects.StatUpSource(game.MartialDefense, 1),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
	}
}

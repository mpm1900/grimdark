package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

func Sharpen() game.Action {
	id := uuid.MustParse("019fcea7-e563-740e-9640-8e862c1b6185")

	return game.Action{
		ID:   id,
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Entity: game.MakeEntity(
			id,
			"Sharpen",
			"Raises user's Melee stat.",
		),
		Config: game.ActionConfig{
			Affinity:    game.Physical,
			TargetCount: 0,
		},
		Resolve: game.AddSourceEffects(
			game.StatusConfig{},
			1,
			effects.StatUpSource(game.Melee, 1),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
	}
}

package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var Sharpen = game.Action{
	ID:   uuid.MustParse("019fcea7-e563-740e-9640-8e862c1b6185"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Sharpen",
		Description: "Raises user's Melee stat.",
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

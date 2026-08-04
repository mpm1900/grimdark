package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var ArmorUp = game.Action{
	ID:   uuid.MustParse("019fceb1-2da1-739b-9394-01899dfd8e2f"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Armor Up",
		Description: "Raises user's Martial Defense stat.",
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

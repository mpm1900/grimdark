package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var LockOn = game.Action{
	ID:   uuid.MustParse("019f87b5-9d77-74d1-a6dc-57b031b0f124"),
	Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
	Config: game.ActionConfig{
		Name:        "Lock On",
		Description: "Raises users's ranged and accuracy stats by 1 stage.",
		Affinity:    game.Kinetic,
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

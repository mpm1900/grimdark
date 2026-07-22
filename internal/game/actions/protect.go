package actions

import (
	"grimdark/internal/game"
	"grimdark/internal/game/effects"

	"github.com/google/uuid"
)

var Protect = protect()

func protect() game.Action {
	return game.Action{
		ID:   uuid.MustParse("019f87f2-fd37-72b4-ab5d-bf23e5830011"),
		Tags: []game.ActionTag{game.ATActor, game.ATWeapon},
		Config: game.ActionConfig{
			Name:        "Protect",
			Description: "User is protected from attacks and actions.",
			Affinity:    game.Kinetic,
			TargetCount: 0,
			Priority:    game.ActionPriorityProtect,
			Cooldown:    1,
		},
		Resolve: game.AddSourceEffects(
			game.StatusConfig{},
			1,
			effects.ProtectedSource(),
		),
		ValidateContext:  game.TrueGameFilter,
		TargetsPredicate: game.NoneActors,
	}
}
